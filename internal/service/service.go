package service

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"sample/internal/model"
)

type Cred struct {
	log    *logrus.Logger
	client *http.Client
}

func NewCred(logger *logrus.Logger, client *http.Client) *Cred {
	return &Cred{
		log:    logger,
		client: client,
	}
}

func (s *Cred) Proxy(ctx context.Context, req model.Req) (rs model.Res, err error) {
	s.log.Infof("req: %+v", req)

	newReq, reqErr := http.NewRequestWithContext(ctx, req.Method, req.Url, nil)
	if reqErr != nil {
		msg := fmt.Errorf("newRequest error -> %w", reqErr)
		s.log.Error(msg)
		return rs, msg
	}
	newReq.Header.Set("Content-Type", "application/json")
	if len(req.Headers) != 0 {
		for key, val := range req.Headers {
			newReq.Header.Set(key, val)
		}
	}
	fmt.Println(newReq)

	res, resErr := s.client.Do(newReq)
	if resErr != nil {
		msg := fmt.Errorf("doing request for %s error -> %w", req.Url, resErr)
		s.log.Error(msg)
		return rs, msg
	}

	s.log.Infof("%+v", res)

	out := model.Res{
		Id:      req.ReqId,
		Status:  res.StatusCode,
		Headers: res.Header,
		Length:  res.ContentLength,
	}

	return out, nil
}
