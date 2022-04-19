package model

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	StatusCode int         `json:"code,omitempty"`
	Message    string      `json:"mesage,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func ResponseHelper(c echo.Context, StatusCode int, Message string, Data interface{}) error {
	return c.JSON(StatusCode, Response{
		StatusCode: StatusCode,
		Message:    Message,
		Data:       Data,
	})
}
