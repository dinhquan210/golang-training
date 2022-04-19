package req

type ReqUserCreatImage struct {
	URLs_full   string `JSON:"urls"`
	Width       int    `JSON:"width"`
	Height      int    `JSON:"height"`
	Description string `JSON:"description"`
}
