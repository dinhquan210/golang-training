package middleware

import (
	"golang-training/model"
	"golang-training/model/req"
	"net/http"

	"github.com/labstack/echo/v4"
)

func IsAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			//handler logic
			req := req.ReqSignIn{}
			if err := c.Bind(&req); err != nil {
				model.ResponseHelper(c, http.StatusBadRequest, err.Error(), nil)
			}
			if req.Email != "admin@gmail.com" {
				model.ResponseHelper(c, http.StatusBadRequest, "ban khong co quyen goi api nay", nil)
			}
			return next(c)
		}
	}
}
