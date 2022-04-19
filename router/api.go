package router

import (
	"golang-training/handler"
	"golang-training/middleware"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo         *echo.Echo
	UserHandler  handler.UserHandler
	ImageHandler handler.ImageHandler
}

func (api *API) SetupRouter() {
	//welcome
	api.Echo.GET("/", handler.Welcome)
	//user
	api.Echo.POST("/user/sign-up", api.UserHandler.HandlerSignUp)
	api.Echo.POST("/user/sign-in", api.UserHandler.HandlerSignIn)
	//profile
	user := api.Echo.Group("/user", middleware.JWTMiddleware())
	user.GET("/profile", api.UserHandler.Profile)
	user.PUT("/profile/update", api.UserHandler.UpdateProfile)
	user.PUT("/creat-image", api.UserHandler.CreateImage)

	//image
	api.Echo.GET("/random-image", api.ImageHandler.RandomImage)
	api.Echo.PUT("/update-image", api.ImageHandler.UpdateImage) // by Id
	api.Echo.PUT("/show-images", api.ImageHandler.ShowImages)
	api.Echo.PUT("/delete-image", api.ImageHandler.DeleteImage)
}
