package auth

type UserModel struct {
	ID           uint    `gorm:"primary_key"`
	Firstname    string  `gorm:"column:first_name"`
	Lastname     string  `gorm:"column:last_name"`
	Email        string  `gorm:"column:email;unique_index"`
	Phone        string  `gorm:"column:phone;unique_index"`
	Bio          string  `gorm:"column:bio;size:1024"`
	Image        *string `gorm:"column:image"`
	PasswordHash string  `gorm:"column:password;not null"`
}
