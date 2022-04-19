package errorutil

import "errors"

var (
	GetRandomFail = errors.New("luu anh that bai")
	ImageNotFound = errors.New("khong co id anh nay")
)
