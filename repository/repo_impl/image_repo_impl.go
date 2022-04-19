package repo_impl

import (
	"context"
	"database/sql"
	"fmt"
	"golang-training/db"
	"golang-training/log"
	"golang-training/model"
	"golang-training/model/req"
	"golang-training/repository"
	"golang-training/utils/errorutil"
	"time"
)

type ImageRepoImpl struct {
	sql *db.Sql
}

func NewImageRepo(sql *db.Sql) repository.ImageRepo {
	return &ImageRepoImpl{
		sql: sql,
	}
}

func (i *ImageRepoImpl) SaveImage(context context.Context, image model.Image) (model.Image, error) {
	statement := `
	INSERT INTO images(id, urls_full, urls_raw, urls_regular, created_at, updated_at, width,height)
	VALUES(:id, :urls_full, :urls_raw, :urls_regular, :created_at, :updated_at, :width, :height)
`
	image.CreatedAt = time.Now()
	image.UpdatedAt = time.Now()
	_, err := i.sql.Db.NamedExecContext(context, statement, image)
	if err != nil {
		fmt.Println(err)
		return image, errorutil.GetRandomFail
	}
	return image, err
}

func (i *ImageRepoImpl) UpdateImageDescription(context context.Context, image model.Image) (model.Image, error) {
	sqlStatement := `
		UPDATE images
		SET 
			description =:description,
			updated_at 	  = COALESCE (:updated_at, updated_at)
		WHERE id    = :id
	`
	image.UpdatedAt = time.Now()

	_, err := i.sql.Db.NamedExecContext(context, sqlStatement, image)
	if err != nil {
		log.Error(err.Error())
		return image, err
	}
	return image, nil
}

func (i *ImageRepoImpl) SelectImage(context context.Context, arr []model.Image) ([]model.Image, error) {
	err := i.sql.Db.SelectContext(context, &arr, "SELECT * FROM images;")
	if err != nil {
		log.Error(err.Error())
		return arr, err
	}
	return arr, nil
}

func (i *ImageRepoImpl) SelectImageById(context context.Context, imageId string) (model.Image, error) {

	return model.Image{}, nil
}

func (i *ImageRepoImpl) CheckIdImage(context context.Context, req_id req.ReqImageUpdate) (model.Image, error) {
	var image = model.Image{}
	err := i.sql.Db.GetContext(context, &image, "SELECT * FROM images WHERE id=$1", req_id.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			return image, errorutil.ImageNotFound
		}
		log.Error(err.Error())
		return image, err
	}

	return image, nil
}
