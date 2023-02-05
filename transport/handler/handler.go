package handler

import (
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"sample/internal/model"
	"sample/internal/service"
	"time"
)

type Handlers struct {
	srv *service.Cred
	log *logrus.Logger
}

func NewHandler(service *service.Cred, log *logrus.Logger) *Handlers {
	return &Handlers{srv: service, log: log}
}

func (h *Handlers) Proxy(c echo.Context) error {
	ctx, cancel := context.WithTimeout(c.Request().Context(), 10*time.Second)
	defer cancel()

	reqId := uuid.New()

	req := model.Req{}
	if errBind := c.Bind(&req); errBind != nil {
		return errBind
	}
	req.ReqId = reqId

	res, resErr := h.srv.Proxy(ctx, req)
	if resErr != nil {
		return c.JSON(http.StatusInternalServerError, resErr.Error())
	}

	return c.JSON(http.StatusOK, res)
}
