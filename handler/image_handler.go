package handler

import (
	"golang-training/log"
	"golang-training/model"
	"golang-training/model/req"
	"golang-training/repository"
	"golang-training/utils/unsplashutils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ImageHandler struct {
	ImageRepo repository.ImageRepo
}

func (i *ImageHandler) RandomImage(c echo.Context) error {
	// Create a Resty Client
	image := unsplashutils.CreateUnsplash()
	image, err := i.ImageRepo.SaveImage(c.Request().Context(), image)
	if err != nil {
		log.Error(err.Error())
		return model.ResponseHelper(c, http.StatusConflict, err.Error(), nil)
	}
	return model.ResponseHelper(c, http.StatusOK, "lấy ảnh thành công", image)
}

func (i *ImageHandler) UpdateImage(c echo.Context) error {
	req := req.ReqImageUpdate{}
	if err := c.Bind(&req); err != nil {
		return err
	}
	image, err := i.ImageRepo.CheckIdImage(c.Request().Context(), req)
	if err != nil {
		return model.ResponseHelper(c, http.StatusUnauthorized, err.Error(), nil)
	}

	image, err = i.ImageRepo.UpdateImageDescription(c.Request().Context(), image)
	if err != nil {
		model.ResponseHelper(c, http.StatusUnprocessableEntity, err.Error(), nil)
	}
	return model.ResponseHelper(c, http.StatusCreated, "Update thong tin anh thanh cong", image)
}

func (i *ImageHandler) ShowImages(c echo.Context) error {

	arr, err := i.ImageRepo.SelectImage(c.Request().Context(), []model.Image{})
	if err != nil {
		log.Error()
	}
	return model.ResponseHelper(c, http.StatusOK, "Xử lý thành công", arr)
}
