package datatransfers

type RegisterRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required,min=8,max=16" json:"password"`
	Email    string `validate:"required,email" json:"email"`
}

type LoginRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}
