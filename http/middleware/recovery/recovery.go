/*
Package recovery is net/http handler to handle error from recovery
*/
package recovery

import (
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	Handler    http.HandlerFunc
	LoggerFunc func(template string, args ...interface{})
}

func New(config *Config) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					// Check for a broken connection, as it is not really a
					// condition that warrants a panic stack trace.
					var brokenPipe bool
					if ne, ok := err.(*net.OpError); ok {
						if se, ok := ne.Err.(*os.SyscallError); ok {
							if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
								strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
								brokenPipe = true
							}
						}
					}

					if config == nil {
						config = &Config{
							LoggerFunc: log.Printf,
						}
					} else {
						if config.LoggerFunc == nil {
							config.LoggerFunc = log.Printf
						}
					}

					// If the connection is dead, we can't write a status to it.
					if !brokenPipe {
						config.LoggerFunc("recover from panic, reason: %s", err)
						if config.Handler != nil {
							config.Handler(w, r)
						} else {
							http.Error(w, "recover from panic!", http.StatusInternalServerError)
						}
					}
				}
			}()
			h.ServeHTTP(w, r)
		})
	}
}
