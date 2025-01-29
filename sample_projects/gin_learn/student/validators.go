package student

import (
	"gin_learn/common"

	"github.com/gin-gonic/gin"
)

type StudentModelValidator struct {
	Student struct {
		FirstName  string `form:"firstname" json:"firstname" binding:"required,alphanum,min=4,max=255"`
		LastName   string `form:"lastname" json:"lastname" binding:"required,alphanum,min=4,max=255"`
		RoleNumber string `form:"rolenumber" json:"rolenumber" binding:"required,alphanum,min=3,max=255"`
	} `json:"student"`
	studentModel StudentModel `json:"-"`
}

// You can put the default value of a Validator here
func NewStudentModelValidator() StudentModelValidator {
	studentModelValidator := StudentModelValidator{}
	//studentModelValidator.User.Email ="w@g.cn"
	return studentModelValidator
}

func (self *StudentModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}
	self.studentModel.FirstName = self.Student.FirstName
	self.studentModel.LastName = self.Student.LastName
	self.studentModel.RoleNumber = self.Student.RoleNumber

	return nil
}
