package httpserv

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// GetClientIP ...
func GetClientIP(r *http.Request) string {
	addr := r.Header.Get("X-Real-Ip")
	if addr == "" {
		addr = r.Header.Get("X-Forwarded-For")
	}
	if addr == "" {
		addr = r.RemoteAddr
	}
	return addr
}

// Middleware ...
type Middleware func(http.Handler) http.Handler

// MiddlewareChain ...
func MiddlewareChain(handler http.Handler, midllewares ...Middleware) http.Handler {
	if len(midllewares) == 0 {
		return handler
	}
	wrapped := handler
	for i := 0; i < len(midllewares); i++ {
		wrapped = midllewares[i](wrapped)
	}
	return wrapped // mw1(mw2(mw3(handler)))
}

// AccessLog ...
func AccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := GetClientIP(r)
		log.WithFields(log.Fields{
			"module":   "httpserv",
			"clientIP": clientIP,
			"method":   r.Method,
		}).Debug(r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
