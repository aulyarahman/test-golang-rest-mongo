package logging

import (
	"fmt"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func makeLogEntry(c echo.Context) *log.Entry {
	if c == nil {
		return log.WithFields(log.Fields{
			"at": time.Now(),
		})
	}

	return log.WithFields(log.Fields{
		"at":     time.Now(),
		"method": c.Request().Method,
		"uri":    c.Request().URL.String(),
		"ip":     c.Request().RemoteAddr,
	})
}

func MiddlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		makeLogEntry(c).Info("incomig request")
		return next(c)
	}
}

func ErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if ok {
		report.Message = fmt.Sprintf("Http error &d - %v", report.Code, report.Message)
	} else {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	makeLogEntry(c).Error(report.Message)
	c.HTML(report.Code, report.Message.(string))
}

func Info(args ...interface{}) {
	logEntry := makeLogEntry(nil)
	logEntry.Info(args...)
}
