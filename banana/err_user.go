package banana

import "errors"

var (
	UesrConflict   = errors.New("nguoi dung da ton tai")
	SignUpFail     = errors.New("dang ky that bai")
	UserNotFound   = errors.New("khong tim thay nguoi dung nay")
	UserNotUpdated = errors.New("update thong tin that bai")
)
