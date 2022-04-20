package model

import "time"

type Image struct {
	ImageID      string    `json:"-" db:"id, omitempty"`
	URLs_Raw     *string   `json:"urls_raw,omitempty" db:"urls_raw,omitempty"`
	URLs_full    string    `json:"urls_full,omitempty" db:"urls_full,omitempty"`
	URLs_regular *string   `json:"urls_egular,omitempty" db:"urls_regular,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" db:"updated_at,omitempty"`
	Width        int       `json:"width,omitempty" db:"width,omitempty"`
	Height       int       `json:"height,omitempty" db:"height,omitempty"`
	Description  *string   `json:"description,omitempty" db:"description, omitempty"`
	User_Creat   *string   `json:"user,omitempty" db:"user_creat, omitempty"`
	Like         int
	DisLike      int
}
