package core

import (
	"github.com/labstack/gommon/log"
	"net/http"
)

type statusCapturingResponseWriter struct {
	http.ResponseWriter
	status int
}

func (s *statusCapturingResponseWriter) WriteHeader(statusCode int) {
	s.status = statusCode
	s.ResponseWriter.WriteHeader(statusCode)
}

// LoggingMiddleware: Register Middleware and Logging
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusCapturingWriter := &statusCapturingResponseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		next.ServeHTTP(statusCapturingWriter, r)

		log.Printf("Main [%s] %s - %d\n", r.Method, r.RequestURI,
			statusCapturingWriter.status)
	})
}