package httputils

import (
	"compress/gzip"
	"io"
	"net/http"
)

type GzipResponseWriter struct {
	io.WriteCloser
	http.ResponseWriter
}

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

func (w *GzipResponseWriter) Close() {
	w.WriteCloser.Close()
}
