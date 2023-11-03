package middlewares

import (
	"log"
	"net/http"
)

type loggingResWriter struct {
	http.ResponseWriter
	code int
}

func NewLoggingResWriter(w http.ResponseWriter) *loggingResWriter {
	return &loggingResWriter{
		ResponseWriter: w,
		code:           http.StatusOK,
	}
}

func (lw *loggingResWriter) WriteHeader(code int) {
	lw.code = code
	lw.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		traceID := NewTraceID()

		log.Printf("[%d]%s %s\n", traceID, req.RequestURI, req.Method)

		ctx := SetTraceID(req.Context(), traceID)
		req = req.WithContext(ctx)

		lw := NewLoggingResWriter(w)

		next.ServeHTTP(lw, req)

		log.Printf("[%d] res: %d\n", traceID, lw.code)
	})
}
