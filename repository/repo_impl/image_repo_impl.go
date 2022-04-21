package repo_impl

import (
	"context"
	"database/sql"
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

func (i *ImageRepoImpl) SelectImage(context context.Context) ([]model.Image, error) {
	arr := make([]model.Image, 0)
	err := i.sql.Db.SelectContext(context, &arr, "SELECT * FROM images;")
	if err != nil {
		log.Error(err.Error())
		return arr, err
	}

	for j := 0; j < len(arr); j++ {
		arr[j].Like, _ = i.CountLikeImage(arr[j].ImageID)
		arr[j].DisLike, _ = i.CountDislikeImage(arr[j].ImageID)
	}
	return arr, nil
}

func (i *ImageRepoImpl) CheckIdImage(context context.Context, id string) (model.Image, error) {
	var image = model.Image{}
	err := i.sql.Db.GetContext(context, &image, "SELECT * FROM images WHERE id=$1", id)

	if err != nil {
		if err == sql.ErrNoRows {
			return image, errorutil.ImageNotFound
		}
		log.Error(err.Error())
		return image, err
	}

	return image, nil
}

func (i *ImageRepoImpl) DelImageById(context context.Context, req_id req.ReqImageDelete) error {

	_, err := i.sql.Db.Exec("DELETE FROM images WHERE id=$1", req_id.Id)
	if err != nil {
		return errorutil.DeleteImageFail
	}
	return nil
}

func (i *ImageRepoImpl) SelectImageByUser(context context.Context, user string) ([]model.Image, error) {
	arr := make([]model.Image, 0)
	err := i.sql.Db.SelectContext(context, &arr, "SELECT * FROM images WHERE user_creat = $1", user)
	if err != nil {
		log.Error(err.Error())
		return arr, err
	}
	return arr, nil
}

func (i *ImageRepoImpl) CountLikeImage(id string) (int, error) {
	var count int
	err := i.sql.Db.Get(&count, "SELECT COUNT(*) FROM reacts WHERE react = 'like' AND id_image=$1", id)
	if err != nil {
		log.Error(err.Error())
		return count, err
	}
	return count, nil
}

func (i *ImageRepoImpl) CountDislikeImage(id string) (int, error) {
	var count int
	err := i.sql.Db.Get(&count, "SELECT COUNT(*) FROM reacts WHERE react = 'dislike' AND id_image=$1", id)
	if err != nil {
		log.Error(err.Error())
		return count, err
	}
	return count, nil
}

func (i *ImageRepoImpl) DeleteReactById(id string) error {
	_, err := i.sql.Db.Exec("DELETE FROM reacts WHERE id=$1", id)
	if err != nil {
		return errorutil.DeleteReactFail
	}
	return nil
}

func (i *ImageRepoImpl) SelectReactById(id string) (model.ReactImage, error) {
	var react = model.ReactImage{}
	err := i.sql.Db.Select(react, "SELECT * FROM images WHERE user_creat = $1", id)
	if err != nil {
		log.Error(err.Error())
		return react, errorutil.ReactNotFound
	}
	return react, nil
}
