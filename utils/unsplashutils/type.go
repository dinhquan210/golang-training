package unsplashutils

import "time"

type URLs struct {
	URLs_Raw     string `json:"raw"`
	URLs_full    string `json:"full"`
	RULs_regular string `json:"regular"`
}

type ResultType struct {
	ImageID     string    `json:"id"`
	URLs        URLs      `json:"urls"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Width       int       `json:"width"`
	Height      int       `json:"height"`
	Description string    `json:"descriptions"`
}
