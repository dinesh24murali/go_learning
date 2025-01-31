# SOLID Principles

Here’s an example of implementing **SOLID principles** in a **Go (Golang) CRUD application** for managing students, using **GORM** for database interactions.

---

### **1. Single Responsibility Principle (SRP)**
Each struct or function should have a **single responsibility**. We'll separate concerns into models, repositories, and services.

---

### **2. Open/Closed Principle (OCP)**
The code should be **open for extension but closed for modification**. We'll use interfaces to allow extending functionality without modifying existing code.

---

### **3. Liskov Substitution Principle (LSP)**
Derived types should be substitutable for their base types. We'll use an interface for our repository layer.

---

### **4. Interface Segregation Principle (ISP)**
Clients should not be forced to depend on **unused interfaces**. We'll create small, specific interfaces instead of a large, general one.

---

### **5. Dependency Inversion Principle (DIP)**
High-level modules should **not depend on low-level modules**, but on **abstractions**. We'll inject dependencies using interfaces.

---

### **Golang Implementation Using SOLID Principles**

#### **1. Define the Student Model**
```go
package models

import "gorm.io/gorm"

type Student struct {
    gorm.Model
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email" gorm:"unique"`
}
```
✅ **Follows SRP**: This struct only defines the student model.

---

#### **2. Create the Repository Interface**
```go
package repository

import "github.com/example/project/models"

type StudentRepository interface {
    Create(student *models.Student) error
    GetByID(id uint) (*models.Student, error)
    GetAll() ([]models.Student, error)
    Update(student *models.Student) error
    Delete(id uint) error
}
```
✅ **Follows ISP**: Only contains required student operations.

---

#### **3. Implement Repository with GORM**
```go
package repository

import (
    "gorm.io/gorm"
    "github.com/example/project/models"
)

type studentRepo struct {
    db *gorm.DB
}

// Constructor function
func NewStudentRepository(db *gorm.DB) StudentRepository {
    return &studentRepo{db}
}

// Implement methods
func (r *studentRepo) Create(student *models.Student) error {
    return r.db.Create(student).Error
}

func (r *studentRepo) GetByID(id uint) (*models.Student, error) {
    var student models.Student
    err := r.db.First(&student, id).Error
    return &student, err
}

func (r *studentRepo) GetAll() ([]models.Student, error) {
    var students []models.Student
    err := r.db.Find(&students).Error
    return students, err
}

func (r *studentRepo) Update(student *models.Student) error {
    return r.db.Save(student).Error
}

func (r *studentRepo) Delete(id uint) error {
    return r.db.Delete(&models.Student{}, id).Error
}
```
✅ **Follows LSP**: The repository implementation can be swapped with any other database storage mechanism.

✅ **Follows DIP**: High-level service depends on the `StudentRepository` interface, not `gorm.DB`.

---

#### **4. Implement the Service Layer**
```go
package service

import (
    "github.com/example/project/models"
    "github.com/example/project/repository"
)

type StudentService interface {
    CreateStudent(student *models.Student) error
    GetStudentByID(id uint) (*models.Student, error)
    GetAllStudents() ([]models.Student, error)
    UpdateStudent(student *models.Student) error
    DeleteStudent(id uint) error
}

type studentService struct {
    repo repository.StudentRepository
}

// Constructor function
func NewStudentService(repo repository.StudentRepository) StudentService {
    return &studentService{repo}
}

// Implement methods
func (s *studentService) CreateStudent(student *models.Student) error {
    return s.repo.Create(student)
}

func (s *studentService) GetStudentByID(id uint) (*models.Student, error) {
    return s.repo.GetByID(id)
}

func (s *studentService) GetAllStudents() ([]models.Student, error) {
    return s.repo.GetAll()
}

func (s *studentService) UpdateStudent(student *models.Student) error {
    return s.repo.Update(student)
}

func (s *studentService) DeleteStudent(id uint) error {
    return s.repo.Delete(id)
}
```
✅ **Follows OCP**: The service can be extended without modifying existing code.

✅ **Follows DIP**: The service depends on `StudentRepository` interface, not a concrete implementation.

---

#### **5. Create the Controller Layer**
```go
package controller

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "github.com/example/project/models"
    "github.com/example/project/service"
)

type StudentController struct {
    service service.StudentService
}

// Constructor function
func NewStudentController(service service.StudentService) *StudentController {
    return &StudentController{service}
}

func (c *StudentController) CreateStudent(ctx *gin.Context) {
    var student models.Student
    if err := ctx.ShouldBindJSON(&student); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.service.CreateStudent(&student); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, student)
}

func (c *StudentController) GetStudent(ctx *gin.Context) {
    id, _ := strconv.Atoi(ctx.Param("id"))
    student, err := c.service.GetStudentByID(uint(id))
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
        return
    }
    ctx.JSON(http.StatusOK, student)
}

func (c *StudentController) GetAllStudents(ctx *gin.Context) {
    students, err := c.service.GetAllStudents()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, students)
}

func (c *StudentController) UpdateStudent(ctx *gin.Context) {
    var student models.Student
    if err := ctx.ShouldBindJSON(&student); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.service.UpdateStudent(&student); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, student)
}

func (c *StudentController) DeleteStudent(ctx *gin.Context) {
    id, _ := strconv.Atoi(ctx.Param("id"))
    if err := c.service.DeleteStudent(uint(id)); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}
```
✅ **Follows SRP**: The controller only handles HTTP requests and responses.

---

#### **6. Initialize the Application**
```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/example/project/controller"
    "github.com/example/project/repository"
    "github.com/example/project/service"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func main() {
    db, _ := gorm.Open(sqlite.Open("students.db"), &gorm.Config{})
    db.AutoMigrate(&models.Student{})

    repo := repository.NewStudentRepository(db)
    studentService := service.NewStudentService(repo)
    studentController := controller.NewStudentController(studentService)

    r := gin.Default()

    r.POST("/students", studentController.CreateStudent)
    r.GET("/students/:id", studentController.GetStudent)
    r.GET("/students", studentController.GetAllStudents)
    r.PUT("/students", studentController.UpdateStudent)
    r.DELETE("/students/:id", studentController.DeleteStudent)

    r.Run(":8080")
}
```
✅ **Follows DIP**: High-level `main.go` depends on interfaces, not concrete implementations.

---

### **Conclusion**
This implementation follows **SOLID principles** and ensures:
✔ **Scalability**  
✔ **Loose Coupling**  
✔ **Extensibility**  
✔ **Maintainability**  

Would you like additional features, such as authentication or middleware integration? 🚀