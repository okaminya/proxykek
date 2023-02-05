package model

import (
	"github.com/google/uuid"
	"net/http"
)

type Req struct {
	Method  string `json:"method"`
	Url     string `json:"url"`
	Headers map[string]string
	ReqId   uuid.UUID
}

type Res struct {
	Id      uuid.UUID   `json:"id"`
	Status  int         `json:"status"`
	Headers http.Header `json:"headers"`
	Length  int64       `json:"length"`
}
