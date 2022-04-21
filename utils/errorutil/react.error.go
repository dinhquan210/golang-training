package errorutil

import "errors"

var (
	DeleteReactFail = errors.New("xoá bản ghi react thất bại")
	ReactFail       = errors.New("Loi khong react duoc")
	ReactNotFound   = errors.New("khong ton tai id_react")
)
