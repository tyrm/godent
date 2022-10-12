package http

import (
	"github.com/spf13/viper"
	"github.com/tyrm/godent/internal/config"
	"net/http"
	"time"
)

// middleware sends http request metrics.
func (*Server) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		started := time.Now()
		l := logger.WithField("func", "middlewareMetrics")

		wx := NewResponseWriter(w)

		// set headers
		wx.Header().Set("Access-Control-Allow-Origin", "*")
		wx.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		wx.Header().Set("Access-Control-Allow-Headers", "*")
		wx.Header().Set("Content-Type", "application/json")
		wx.Header().Set("Permissions-Policy", "interest-cohort=()")
		wx.Header().Set("X-Godent-Version", viper.GetString(config.Keys.SoftwareVersion))

		// Do Request
		next.ServeHTTP(wx, r)

		go func() {
			l.Debugf("rendering %s took %d ms", r.URL.Path, time.Since(started))
		}()
	})
}

// WrapInMiddlewares wraps an http.Handler in the server's middleware.
func (s *Server) WrapInMiddlewares(h http.Handler) http.Handler {
	return s.middleware(
		h,
	)
}
