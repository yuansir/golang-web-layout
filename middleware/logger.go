package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"strconv"
	"time"
)

type (
	// LoggerConfig Logger configs
	LoggerConfig struct {
		Skipper   middleware.Skipper
		Formatter log.Formatter
		Output    io.Writer
	}
)

var (
	// DefaultLoggerConfig is the default Logger middleware config.
	DefaultLoggerConfig = LoggerConfig{
		Skipper:   middleware.DefaultSkipper,
		Formatter: &log.JSONFormatter{},
		Output:    os.Stdout,
	}
)

// LoggerWithConfig Logger with config
func LoggerWithConfig(config LoggerConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = DefaultLoggerConfig.Skipper
	}
	if config.Formatter == nil {
		config.Formatter = DefaultLoggerConfig.Formatter
	}
	if config.Output == nil {
		config.Output = DefaultLoggerConfig.Output
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()
			res := c.Response()
			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}

			stop := time.Now()
			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}

			path := req.URL.Path
			if path == "" {
				path = "/"
			}
			status := res.Status
			latency := stop.Sub(start).String()

			log.WithFields(log.Fields{
				"at":        strconv.FormatInt(time.Now().UnixNano(), 10),
				"status":    status,
				"ip":        c.RealIP(),
				"id":        id,
				"host":      req.Host,
				"uri":       req.RequestURI,
				"method":    req.Method,
				"path":      path,
				"proto":     req.Proto,
				"refer":     req.Referer(),
				"userAgent": req.UserAgent(),
				"size":      strconv.FormatInt(res.Size, 10),
				"latency":   latency,
				"header":    req.Header,
			}).Info("Http request info")
			return nil
		}
	}
}

// Logger middleware logs request
func Logger() echo.MiddlewareFunc {
	return LoggerWithConfig(DefaultLoggerConfig)
}
