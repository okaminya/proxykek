package transport

import (
	"crypto/tls"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"sample/internal/service"
	"sample/transport/handler"
	"time"
)

func StartServer(errCh chan error, app *echo.Echo, l *logrus.Logger) {
	l.SetReportCaller(true)

	app.GET("/ping", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:    100,
		IdleConnTimeout: 90 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Transport: transport,
	}
	defer client.CloseIdleConnections()

	s := service.NewCred(l, client)
	h := handler.NewHandler(s, l)

	group := app.Group("/api/v1")
	group.POST("/proxy", h.Proxy)

	errCh <- app.Start(":9090")
}
