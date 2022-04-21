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

	"github.com/lib/pq"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repository.UserRepo {
	return &UserRepoImpl{
		sql: sql,
	}
}

func (u UserRepoImpl) SaveUser(context context.Context, user model.User) (model.User, error) {
	statement := `
		INSERT INTO users(user_id, email, password, role, full_name, created_at, update_at)
		VALUES(:user_id, :email, :password, :role, :full_name, :created_at, :update_at )
	`
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := u.sql.Db.NamedExecContext(context, statement, user)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, errorutil.UesrConflict
			}

		}
		return user, errorutil.SignUpFail
	}
	return user, nil
}

func (u UserRepoImpl) CheckLogin(context context.Context, loginReq req.ReqSignIn) (model.User, error) {
	var user = model.User{}
	err := u.sql.Db.GetContext(context, &user, "SELECT * FROM users WHERE email=$1", loginReq.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, errorutil.UserNotFound
		}
		log.Error(err.Error())
		return user, err
	}

	return user, nil
}

func (u UserRepoImpl) SelectUserById(context context.Context, userId string) (model.User, error) {
	var user model.User

	err := u.sql.Db.GetContext(context, &user,
		"SELECT * FROM users WHERE user_id = $1", userId)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, errorutil.UserNotFound
		}
		log.Error(err.Error())
		return user, err
	}

	return user, nil
}

func (u UserRepoImpl) UpdateUser(context context.Context, user model.User) (model.User, error) {
	sqlStatement := `
		UPDATE users
		SET 
			full_name  = (CASE WHEN LENGTH(:full_name) = 0 THEN full_name ELSE :full_name END),
			email = (CASE WHEN LENGTH(:email) = 0 THEN email ELSE :email END),
			update_at 	  = COALESCE (:update_at, update_at)
		WHERE user_id    = :user_id
	`
	user.UpdatedAt = time.Now()

	result, err := u.sql.Db.NamedExecContext(context, sqlStatement, user)
	if err != nil {
		log.Error(err.Error())
		return user, err
	}

	count, err := result.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return user, errorutil.UserNotUpdated
	}
	if count == 0 {
		return user, errorutil.UserNotUpdated
	}

	return user, nil
}

func (u UserRepoImpl) SaveImageCreatByUser(context context.Context, image model.Image) (model.Image, error) {
	statement := `
		INSERT INTO images(id, urls_full,updated_at,created_at,width,height,description,user_creat )
		VALUES(:id, :urls_full, :created_at, :updated_at, :width, :height, :description, :user_creat )
	`
	image.CreatedAt = time.Now()
	image.UpdatedAt = time.Now()
	_, err := u.sql.Db.NamedExecContext(context, statement, image)
	if err != nil {
		return image, errorutil.UserCreatImageFail
	}
	return image, nil
}

func (u UserRepoImpl) SaveReactImage(context context.Context, react model.ReactImage) (model.ReactImage, error) {
	statement := `INSERT INTO reacts(id_react, id_image, react, id_user) VALUES(:id_react, :id_image, :react, :id_user)`
	count, _ := u.CountReactByUserAndImage(react.ImageId, react.UserId)
	if count != 1 {
		_, err := u.sql.Db.NamedExecContext(context, statement, react)
		if err != nil {
			return react, errorutil.ReactFail
		}
	}
	_, err := u.UpdateReact(context, react)
	if err != nil {
		return react, nil
	}
	return react, nil
}

func (u UserRepoImpl) SelectReactsByUserId(context context.Context, id string) ([]model.ReactImage, error) {
	arr := make([]model.ReactImage, 0)
	err := u.sql.Db.SelectContext(context, &arr, "SELECT * FROM reacts WHERE id_user = $1", id)
	if err != nil {
		log.Error(err.Error())
		return arr, err
	}
	return arr, nil
}

func (u *UserRepoImpl) CountReactByUserAndImage(id_image string, id_user string) (int, error) {
	var count int
	err := u.sql.Db.Get(&count, "SELECT COUNT(*) FROM reacts WHERE id_user = $1 AND id_image=$2", id_user, id_image)
	if err != nil {
		log.Error(err.Error())
		return count, err
	}
	return count, nil
}

func (u *UserRepoImpl) UpdateReact(context context.Context, react model.ReactImage) (model.ReactImage, error) {
	sqlStatement := `
	UPDATE reacts 
	SET 
		react = :react
	WHERE id_user    = :id_user AND id_image = :id_image
`
	_, err := u.sql.Db.NamedExecContext(context, sqlStatement, react)
	if err != nil {
		log.Error(err.Error())
		return react, err
	}
	return react, nil
}
