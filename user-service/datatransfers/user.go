package datatransfers

type UpdateUserRequest struct {
	Email    string `json:"email"`
	FullName string `json:"fullName"`
}
