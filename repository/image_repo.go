package repository

import (
	"context"
	"golang-training/model"
	"golang-training/model/req"
)

type ImageRepo interface {
	SaveImage(context context.Context, image model.Image) (model.Image, error)
	UpdateImageDescription(context context.Context, image model.Image) (model.Image, error)
	SelectImage(context context.Context, arr []model.Image) ([]model.Image, error)
	SelectImageById(context context.Context, imageId string) (model.Image, error)
	CheckIdImage(context context.Context, id string) (model.Image, error)
	DelImageById(context context.Context, req_id req.ReqImageDelete) error
}
