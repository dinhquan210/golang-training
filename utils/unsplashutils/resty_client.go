package unsplashutils

import (
	"golang-training/model"

	"github.com/go-resty/resty/v2"
)

func CreateUnsplash() model.Image {
	Client := resty.New()
	reBody := ResultType{}
	Client.R().SetResult(&reBody).Get("https://api.unsplash.com/photos/random/?client_id=05qCv0koWY-_KqKyyCRmtrBqtbBISysGPznnA6wCNNg")
	image := model.Image{
		ImageID:      reBody.ImageID,
		URLs_full:    reBody.URLs.URLs_full,
		URLs_regular: reBody.URLs.RULs_regular,
		URLs_Raw:     reBody.URLs.URLs_Raw,
		Width:        reBody.Width,
		Height:       reBody.Height,
		Description:  &reBody.Description,
	}
	return image
}
