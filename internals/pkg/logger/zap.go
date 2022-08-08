package logger

import (
	"net/http"

	"go.uber.org/zap"
)

var sugarLogger *zap.SugaredLogger

func NewLogger() {
	logger, _ := zap.NewProduction()
	sugarLogger = logger.Sugar()
}

func Infof(r *http.Request, url string) {
	if sugarLogger == nil {
		NewLogger()
	}
	sugarLogger.Infof("Success! statusCode = %s for URL %s", http.StatusOK, url)
}

func Errorer(url string, err error) {
	if err != nil {
		l := sugarLogger.With("url", url)
		l.Errorf(err.Error())
	}
}
