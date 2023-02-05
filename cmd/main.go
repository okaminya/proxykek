package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"runtime"
	"sample/transport"
	"strconv"
	"syscall"
)

var log = logrus.New()

func main() {
	log.Fatalf(fmt.Sprintf("Service shut down: %s", run()))
}

func run() error {
	log.Formatter = &logrus.TextFormatter{DisableColors: false,
		FullTimestamp: true,
		ForceColors:   true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := " [FUNC]:" + frame.Function + ":" + strconv.Itoa(frame.Line) + " [MSG]:"
			return "", fileName
		},
	}

	app := echo.New()

	errCh := make(chan error, 1)

	go transport.StartServer(errCh, app, log)

	go gracefulShutdown(errCh)

	return <-errCh
}

func gracefulShutdown(errCh chan<- error) {
	signCh := make(chan os.Signal, 1)
	signal.Notify(signCh, syscall.SIGTERM, syscall.SIGINT)
	errCh <- fmt.Errorf("%s", <-signCh)
}
