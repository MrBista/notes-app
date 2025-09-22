package req

type UserRegsiter struct {
	FullName string `json:"fullName" validate:"required"`
	Username string `json:"username" validate:"username"`
	Email    string `json:"email" validate:"required, email, max=30, min=1"`
	Password string `json:"password" validate:"required"`
}
