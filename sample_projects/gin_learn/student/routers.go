package student

import (
	"fmt"
	"gin_learn/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

var list []StudentModel

func init() {
	fmt.Println("Nice! Wow")
}

func StudentsRegister(router *gin.RouterGroup) {
	router.POST("/", StudentRegistration)
	router.GET("/", GetStudents)
}

// @BasePath /api

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /api/students [post]
func StudentRegistration(c *gin.Context) {
	studentModelValidator := NewStudentModelValidator()
	if err := studentModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	fmt.Print(studentModelValidator)
	lastIndex := cap(list)
	var nextIndex uint = 0
	if lastIndex > 0 {
		lastItem := list[lastIndex-1]
		nextIndex = lastItem.ID + 1
	}

	newItem := StudentModel{
		ID:         nextIndex,
		FirstName:  studentModelValidator.studentModel.FirstName,
		LastName:   studentModelValidator.studentModel.LastName,
		RoleNumber: studentModelValidator.studentModel.RoleNumber,
	}
	list = append(list, newItem)
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Student data received",
	})
}

// @BasePath /api

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /api/students [get]
func GetStudents(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"data": list,
	})
}
