package handlers

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

// Go's http.FileServer by default, lists the directories and files
// of the specified folder to serve and cannot be disabled.
// To prevent directory listing, noDirListing checks if the
// path requests ends in '/'. If it does, then the client is requesting
// to explore a folder and we return a 404 (Not found), else, we just
// call the http.Handler passed as parameter.
func NoDirListing(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path

		if strings.HasSuffix(urlPath, "/") {
			http.NotFound(w, r)
			return
		}

		h.ServeHTTP(w, r)
	})
}

// Validates that the user cookie is set up before calling the handler
// passed as parameter.
func ValidateAuth(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := cookieStore.Get(r, sessionCookieName)

		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		if _, ok := session.Values["data"].(*SessionData); !ok {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	})
}

// JsonHandler is a middleware function for the ApiHandler type. The
// JsonHandler sets the content type of the response as application/json. It
// also checks if the ApiHandler returned an error. When the ApiHandler returns
// an error, then a http.Error is filled with the error data return from the
// ApiHandler. If there's no error, the payload that returned from the
// ApiHandler is parsed to json (if any) and written to the
// http.ResponseWriter.
func JsonHandler(h ApiHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		payload, apiError := h(w, r)
		if apiError != nil {
			if apiError.Error != nil {
				log.Println(apiError.Error)
			}

			// If no message is set, get the default error message from the
			// http module.
			if apiError.Message == "" {
				apiError.Message = http.StatusText(apiError.HttpCode)
			}

			http.Error(w, apiError.Message, apiError.HttpCode)
			return
		}

		// If nothing is to be converted to json then return
		if payload == nil {
			return
		}

		// Convert payload to json and return HTTP OK (200)
		b, err := json.Marshal(payload)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Write(b)
	})
}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

// Write bytes to the io.Writer in the gzipResponseWriter. If no Content-Type
// was set to the response, we must detect it's content type before the
// default ResponseWriter does. If we don't, then the response will be of
// Content-Type 'application/x-gzip'.
func (w gzipResponseWriter) Write(b []byte) (int, error) {
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", http.DetectContentType(b))
	}

	return w.Writer.Write(b)
}

// gzipContent is a middleware function for handlers to encode content to gzip.
func GzipContent(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Accept-Encoding")

		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			h.ServeHTTP(w, r)
			return
		}

		w.Header().Set("Content-Encoding", "gzip")
		gw := gzip.NewWriter(w)
		defer gw.Close()

		grw := &gzipResponseWriter{
			Writer:         gw,
			ResponseWriter: w,
		}

		h.ServeHTTP(grw, r)
	})
}

// Authorize validates privileges for the current user. Each route must have
// an array of privileges that point which users can make a call to it.
//
// Note:
//
// It is assumed that ValidateAuth was called before this function, or at
// least some other session check was done before this.
func Authorize(requiredRoles []string, h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ok          bool
			sessionData *SessionData
		)

		session, err := cookieStore.Get(r, sessionCookieName)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		if sessionData, ok = session.Values["data"].(*SessionData); !ok {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		if len(requiredRoles) == 0 {
			h.ServeHTTP(w, r)
			return
		}

		for _, role := range requiredRoles {
			if role == "ADMIN" && sessionData.IsAdmin {
				h.ServeHTTP(w, r)
				return
			} else if role == "TEACHER" && sessionData.IsTeacher {
				h.ServeHTTP(w, r)
				return
			}
		}

		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	})
}
