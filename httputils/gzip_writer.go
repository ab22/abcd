package httputils

import (
	"compress/gzip"
	"io"
	"net/http"
)

// A GzipResponseWriter is a handy implementation of io.WriteCloser
// used on top of http.ResponseWriter to GZip all of the content that is
// written to the http.ResponseWriter.
type GzipResponseWriter struct {
	io.WriteCloser
	http.ResponseWriter
}

// NewGzipResponseWriter creates a new instance of GzipResponseWriter with a
// underlying gzip writer and a specified ResponseWriter.
func NewGzipResponseWriter(w http.ResponseWriter) *GzipResponseWriter {
	return &GzipResponseWriter{
		WriteCloser:    gzip.NewWriter(w),
		ResponseWriter: w,
	}
}

// Write bytes to the io.Writer in the GzipResponseWriter. If no Content-Type
// was set to the response, we must detect it's content type before the
// default ResponseWriter does. If we don't, then the response will be of
// Content-Type 'application/x-gzip'.
func (w *GzipResponseWriter) Write(b []byte) (int, error) {
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", http.DetectContentType(b))
	}

	return w.WriteCloser.Write(b)
}

// Close closes the underlying gzip writer.
func (w *GzipResponseWriter) Close() {
	w.WriteCloser.Close()
}
