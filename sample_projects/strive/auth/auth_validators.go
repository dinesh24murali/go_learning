package auth

type RegisterValidator struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"password"`
}

func NewRegisterValidator() RegisterValidator {
	productModelValidator := RegisterValidator{}
	return productModelValidator
}
