package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Employee struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"not null"`
	KTP    string `gorm:"not null"`
	Status string `gorm:"not null"`
}

type Exam struct {
	ID         uint
	EmployeeID uint
	Score      int
}

var (
	db  *gorm.DB
	err error
)

func main() {
	// Membuat koneksi ke database
	dsn := "root:@tcp(localhost:3306)/assignment_cbn?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Membuat tabel jika belum ada
	db.AutoMigrate(&Employee{}, &Exam{})

	// Membuat instance Gin
	r := gin.Default()

	// Endpoint untuk mendapatkan daftar karyawan
	r.GET("/employees", getEmployees)

	// Endpoint untuk menambahkan karyawan
	r.POST("/employees", createEmployee)

	// Endpoint untuk mengupdate data karyawan
	r.PUT("/employees/:id", updateEmployee)

	// Endpoint untuk menghapus karyawan
	r.DELETE("/employees/:id", deleteEmployee)

	// Endpoint untuk mendaftarkan ujian karyawan
	r.POST("/exams", createExam)

	// Endpoint untuk menginput nilai ujian
	r.PUT("/exams/:id", updateExam)

	// Menjalankan server
	r.Run(":8080")
}

// Handler untuk mendapatkan daftar karyawan
func getEmployees(c *gin.Context) {
	var employees []Employee
	db.Find(&employees)
	c.JSON(http.StatusOK, gin.H{"data": employees})
}

// Handler untuk menambahkan karyawan
func createEmployee(c *gin.Context) {
	var employee Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&employee)
	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// Handler untuk mengupdate data karyawan
func updateEmployee(c *gin.Context) {
	var employee Employee
	id := c.Param("id")
	if err := db.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Employee not found"})
		return
	}
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&employee)
	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// Handler untuk menghapus karyawan
func deleteEmployee(c *gin.Context) {
	var employee Employee
	id := c.Param("id")
	if err := db.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Employee not found"})
		return
	}
	db.Delete(&employee)
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}

// Middleware untuk memeriksa status karyawan sebelum mendaftar ujian
func checkEmployeeStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var employee Employee
		id := c.PostForm("employee_id")
		if err := db.First(&employee, id).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Employee not found"})
			c.Abort()
			return
		}
		if employee.Status != "active" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Employee is not active"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// Handler untuk mendaftarkan ujian karyawan
func createExam(c *gin.Context) {
	var exam Exam
	if err := c.ShouldBindJSON(&exam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&exam)
	c.JSON(http.StatusOK, gin.H{"data": exam})
}

// Handler untuk menginput nilai ujian
func updateExam(c *gin.Context) {
	var exam Exam
	id := c.Param("id")
	if err := db.First(&exam, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Exam not found"})
		return
	}
	if err := c.ShouldBindJSON(&exam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&exam)
	c.JSON(http.StatusOK, gin.H{"data": exam})
}
