package handler

import (
	"fmt"
)

type APIResponse struct {
	Msg string `json:"msg,omitempty"`
}

type APIRequest struct {
	Token     string `json:"token" binding:"required"`
	Sign      string `json:"sign" binding:"required"`
	Timestamp int64  `json:"ts" binding:"required"`
}

type ErrorCode uint

const (
	BADREQUEST_ERROR ErrorCode = 400
	INTERNAL_ERROR   ErrorCode = 500
)

type APIError struct {
	Code ErrorCode `json:"code,omitempty"`
	Msg  string    `json:"msg,omitempty"`
}

func (this APIError) Error() string {
	return fmt.Sprintf("CODE:%d, MSG:%s", this.Code, this.Msg)
}
