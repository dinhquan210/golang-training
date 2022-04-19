package handler

import (
	"fmt"
	"golang-training/log"
	"golang-training/model"
	req "golang-training/model/req"
	"golang-training/repository"
	"golang-training/security"
	"golang-training/utils/errorutil"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserRepo repository.UserRepo
}

func (u *UserHandler) HandlerSignUp(c echo.Context) error {
	req := req.ReqSignUp{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		model.ResponseHelper(c, http.StatusBadRequest, err.Error(), nil)
	}
	if err := c.Validate(req); err != nil {
		log.Error(err.Error())
		model.ResponseHelper(c, http.StatusBadRequest, err.Error(), nil)
	}

	hash := security.HashAndSalt([]byte(req.PassWord))
	role := model.MEMBER.String()

	userId, err := uuid.NewUUID()
	if err != nil {
		log.Error(err.Error())
		model.ResponseHelper(c, http.StatusForbidden, err.Error(), nil)
	}

	user := model.User{
		UserId:   userId.String(),
		FullName: req.FullName,
		Email:    req.Email,
		PassWord: hash,
		Role:     role,
		Token:    "",
	}

	user, err = u.UserRepo.SaveUser(c.Request().Context(), user)
	if err != nil {
		log.Error(err.Error())
		model.ResponseHelper(c, http.StatusConflict, err.Error(), nil)
	}

	//GEN token
	token, err := security.GenToken(user)
	if err != nil {
		model.ResponseHelper(c, http.StatusInternalServerError, err.Error(), nil)
	}
	user.Token = token
	return model.ResponseHelper(c, http.StatusOK, "Xử lý thành công", user)
}

func (u *UserHandler) HandlerSignIn(c echo.Context) error {
	req := req.ReqSignIn{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		model.ResponseHelper(c, http.StatusBadRequest, err.Error(), nil)
	}

	if err := c.Validate(req); err != nil {
		log.Error(err.Error())
		model.ResponseHelper(c, http.StatusBadRequest, err.Error(), nil)
	}

	user, err := u.UserRepo.CheckLogin(c.Request().Context(), req)
	if err != nil {
		return model.ResponseHelper(c, http.StatusUnauthorized, err.Error(), nil)
	}

	// neu ko error sql -> check pass
	isTheSame := security.ComparePasswords(user.PassWord, []byte(req.Password))
	if !isTheSame {
		return model.ResponseHelper(c, http.StatusUnauthorized, "Đăng nhập thất bại", nil)
	}

	token, err := security.GenToken(user)
	if err != nil {
		model.ResponseHelper(c, http.StatusInternalServerError, err.Error(), nil)
	}
	user.Token = token
	return model.ResponseHelper(c, http.StatusOK, "Xử lý thành công", user)
}

func (u *UserHandler) Profile(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)
	fmt.Println(claims.FullName)

	user, err := u.UserRepo.SelectUserById(c.Request().Context(), claims.UserId)
	if err != nil {
		if err == errorutil.UserNotFound {
			model.ResponseHelper(c, http.StatusNotFound, err.Error(), nil)
		}
		model.ResponseHelper(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return model.ResponseHelper(c, http.StatusOK, "Xử lý thành công", user)
}

func (u *UserHandler) UpdateProfile(c echo.Context) error {
	req := req.ReqUpdateUser{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	// validate thông tin gửi lên
	err := c.Validate(req)
	if err != nil {
		model.ResponseHelper(c, http.StatusBadRequest, err.Error(), nil)
	}

	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.JwtCustomClaims)
	user := model.User{
		UserId:   claims.UserId,
		FullName: req.FullName,
		Email:    req.Email,
	}

	user, err = u.UserRepo.UpdateUser(c.Request().Context(), user)
	if err != nil {
		model.ResponseHelper(c, http.StatusUnprocessableEntity, err.Error(), nil)
	}
	return model.ResponseHelper(c, http.StatusCreated, "Xử lý thành công", user)
}

func (u *UserHandler) CreateImage(c echo.Context) error {
	req := req.ReqUserCreatImage{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	ImageID, err := uuid.NewUUID()
	if err != nil {
		log.Error(err.Error())
		model.ResponseHelper(c, http.StatusForbidden, err.Error(), nil)
	}
	image := model.Image{
		ImageID:     ImageID.String(),
		URLs_full:   req.URLs_full,
		Width:       req.Width,
		Height:      req.Height,
		Description: &req.Description,
	}
	return model.ResponseHelper(c, http.StatusOK, "Xử lý thành công", image)
}
