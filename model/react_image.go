package model

type ReactImage struct {
	ReactId string `json:"-" db:"id_react"`
	ImageId string `json:"image_id" db:"id_image"`
	React   string `json:"react" db:"react"`
	UserId  string `json:"user_id" db:"id_user"`
}
