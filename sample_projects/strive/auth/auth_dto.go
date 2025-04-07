package auth

type RegisterUserDto struct {
	Phone    string
	Email    string
	Password string
}

type UpdateUserDto struct {
	FirstName string
	LastName  string
	Status    uint
}
