package req

type ReqSignUp struct {
	FullName string `json:"fullname,omitempty" validate:"required"` // tags
	Email    string `json:"email,omitempty" validate:"required"`
	PassWord string `json:"password,omitempty" validate:"pwd"`
}
