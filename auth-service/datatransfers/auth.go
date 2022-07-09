package datatransfers

type AuthRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required,min=8,max=16" json:"password"`
}
