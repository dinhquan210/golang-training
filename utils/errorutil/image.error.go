package errorutil

import "errors"

var (
	GetRandomFail   = errors.New("Lưu ảnh thất bại")
	ImageNotFound   = errors.New("không tồn tại id ảnh này")
	DeleteImageFail = errors.New("xoá ảnh thất bại")
)
