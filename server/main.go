package main

import (
	"fmt"
	"golang-training/db"
	"golang-training/handler"
	"golang-training/helper"
	log "golang-training/log"
	"golang-training/repository/repo_impl"
	"golang-training/router"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func init() {
	fmt.Println("DEV ENVIROMENT")
	os.Setenv("APP_NAME", "github")
	log.InitLogger(false)
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	sql := &db.Sql{
		Host:     "postgres",
		Port:     5432,
		UserName: "postgres",
		PassWord: "postgres",
		DbName:   "golang-training",
	}

	sql.Connect()
	defer sql.Close()
	e := echo.New()
	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	imageHandler := handler.ImageHandler{
		ImageRepo: repo_impl.NewImageRepo(sql),
	}

	StructValidator := helper.NewStructValidator()
	StructValidator.RegisterValidate()

	e.Validator = StructValidator

	api := router.API{
		Echo:         e,
		UserHandler:  userHandler,
		ImageHandler: imageHandler,
	}

	api.SetupRouter()
	e.Logger.Fatal(e.Start(":8080"))

}
