package req

type ReqUserCreatImage struct {
	URLs_full   string `json:"urls"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Description string `json:"description"`
}
