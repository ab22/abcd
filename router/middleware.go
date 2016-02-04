package router

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/ab22/abcd/config"
	"github.com/ab22/abcd/router/httputils"
	"github.com/gorilla/sessions"
	"golang.org/x/net/context"
)

type MiddlewareFunc func(httputils.ContextHandler) httputils.ContextHandler

// Go's http.FileServer by default, lists the directories and files
// of the specified folder to serve and cannot be disabled.
// To prevent directory listing, noDirListing checks if the
// path requests ends in '/'. If it does, then the client is requesting
// to explore a folder and we return a 404 (Not found), else, we just
// call the http.Handler passed as parameter.
func NoDirListing(h httputils.ContextHandler) httputils.ContextHandler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		urlPath := r.URL.Path

		if urlPath == "" || strings.HasSuffix(urlPath, "/") {
			http.NotFound(w, r)
			return nil
		}

		return h(ctx, w, r)
	}
}

// Validates that the user cookie is set up before calling the handler
// passed as parameter.
func ValidateAuth(h httputils.ContextHandler) httputils.ContextHandler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		var (
			sessionData *SessionData
			err         error
			ok          bool
			cookieStore *sessions.CookieStore
			session     *sessions.Session
			cfg         *config.Config
		)

		cookieStore, ok = ctx.Value("cookieStore").(*sessions.CookieStore)

		if !ok {
			httputils.WriteError(w, http.StatusInternalServerError, "")
			return fmt.Errorf("validate auth: could not cast value as cookie store:", ctx.Value("cookieStore"))
		}

		session, err = cookieStore.Get(r, SessionCookieName)

		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return nil
		}

		sessionData, ok = session.Values["data"].(*SessionData)

		if !ok {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return nil
		} else if time.Now().After(sessionData.ExpiresAt) {
			session.Options.MaxAge = -1
			session.Save(r, w)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)

			return nil
		}

		// Extend the session's lifetime.
		cfg, ok = ctx.Value("config").(*config.Config)

		if !ok {
			httputils.WriteError(w, http.StatusInternalServerError, "")
			return fmt.Errorf("validate auth: error casting config object", ctx.Value("config"))
		}

		sessionData.ExpiresAt = time.Now().Add(cfg.SessionLifeTime)

		session.Save(r, w)

		authenticatedContext := context.WithValue(ctx, "sessionData", sessionData)
		return h(authenticatedContext, w, r)
	}
}

// gzipContent is a middleware function for handlers to encode content to gzip.
func GzipContent(h httputils.ContextHandler) httputils.ContextHandler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		w.Header().Add("Vary", "Accept-Encoding")

		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			return h(ctx, w, r)
		}

		w.Header().Set("Content-Encoding", "gzip")

		gzipResponseWriter := httputils.NewGzipResponseWriter(w)
		defer gzipResponseWriter.Close()

		return h(ctx, gzipResponseWriter, r)
	}
}

// Authorize validates privileges for the current user. Each route must have
// an array of privileges that point which users can make a call to it.
//
// Note:
//
// It is assumed that ValidateAuth was called before this function, or at
// least some other session check was done before this.
func Authorize(h httputils.ContextHandler) httputils.ContextHandler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		var (
			requiredRoles []string
			sessionData   *SessionData
			route         Route
			ok            bool
		)

		sessionData, ok = ctx.Value("sessionData").(*SessionData)

		if !ok {
			httputils.WriteError(w, http.StatusInternalServerError, "")
			return fmt.Errorf("authorize: could not cast value as session data:", ctx.Value("sessionData"))
		}

		route, ok = ctx.Value("route").(Route)

		if !ok {
			httputils.WriteError(w, http.StatusInternalServerError, "")
			return fmt.Errorf("authorize: could not cast value as route:", ctx.Value("route"))
		}

		requiredRoles = route.RequiredRoles()

		if len(requiredRoles) == 0 {
			return h(ctx, w, r)
		}

		for _, role := range requiredRoles {
			if role == "ADMIN" && sessionData.IsAdmin {
				return h(ctx, w, r)
			} else if role == "TEACHER" && sessionData.IsTeacher {
				return h(ctx, w, r)
			}
		}

		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return nil
	}
}

// HandleError sets the appropriate headers to the response if a http handler
// returned an error. This might be used in the future if different types of
// errors are returned.
func HandleHttpError(h httputils.ContextHandler) httputils.ContextHandler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		err := h(ctx, w, r)

		if err != nil {
			httputils.WriteError(w, http.StatusInternalServerError, "")
		}

		return err
	}
}
