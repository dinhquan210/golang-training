package repo_impl

import (
	"context"
	"fmt"
	"golang-training/banana"
	"golang-training/db"
	"golang-training/log"
	"golang-training/model"
	"golang-training/repository"
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
		return image, banana.GetRandomFail
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
