package model

import "time"

type Image struct {
	ImageID      string    `JSON:"-" db:"id, omitempty"`
	URLs_Raw     string    `JSON:"raw,omitempty" db:"urls_raw,omitempty"`
	URLs_full    string    `JSON:"full,omitempty" db:"urls_full,omitempty"`
	URLs_regular string    `JSON:"regular,omitempty" db:"urls_regular,omitempty"`
	CreatedAt    time.Time `JSON:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt    time.Time `JSON:"updated_at,omitempty" db:"updated_at,omitempty"`
	Width        int       `JSON:"width,omitempty" db:"width,omitempty"`
	Height       int       `JSON:"height,omitempty" db:"height,omitempty"`
	Description  string    `JSON:"description,omitempty" db:"description,omitempty"`
}
