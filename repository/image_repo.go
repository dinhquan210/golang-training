package repository

import (
	"context"
	"golang-training/model"
	"golang-training/model/req"
)

type ImageRepo interface {
	SaveImage(context context.Context, image model.Image) (model.Image, error)
	UpdateImageDescription(context context.Context, image model.Image) (model.Image, error)
	SelectImage(context context.Context) ([]model.Image, error)
	CheckIdImage(context context.Context, id string) (model.Image, error)
	DelImageById(context context.Context, req_id req.ReqImageDelete) error
	SelectImageByUser(context context.Context, user string) ([]model.Image, error)
	CountLikeImage(id string) (int, error)
	CountDislikeImage(id string) (int, error)
}
