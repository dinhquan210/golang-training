package repository

import (
	"context"
	"golang-training/model"
	"golang-training/model/req"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
	CheckLogin(context context.Context, loginReq req.ReqSignIn) (model.User, error)
	SelectUserById(context context.Context, userId string) (model.User, error)
	UpdateUser(context context.Context, user model.User) (model.User, error)
	SaveImageCreatByUser(context context.Context, image model.Image) (model.Image, error)
	SaveReactImage(context context.Context, react model.ReactImage) (model.ReactImage, error)
	SelectReactsByUserId(context context.Context, id string) ([]model.ReactImage, error)
}
