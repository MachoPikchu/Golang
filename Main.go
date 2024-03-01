package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	username := "root"
	password := "happyday020304"
	dbName := "order_service"
	dbHost := "localhost"
	dbPort := "3306"

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, dbHost, dbPort, dbName)
	ConnectToDB(dataSourceName)
	db.AutoMigrate(&Student{}, &Course{}, &Department{}, &Enrollment{}, &Instructor{})

	//course := Course{Name: "Introduction to Programming", InstructorID: 1, DepartmentID: 1}
	//db.Create(&course)

	//fmt.Println("Sample data inserted successfully")
	// Querys
	/*var students []Student
	db.Where("department_id = ?", 1).Find(&students)
	fmt.Println("Students in department 1:", students)

	var courses []Course
	db.Where("instructor_id = ?", 1).Find(&courses)
	fmt.Println("Courses taught by instructor 1:", courses)

	var enrollments []Enrollment
	db.Where("student_id = ?", 1).Find(&enrollments)
	fmt.Sprintf("Enrollments for student %s", enrollments)*/
}

func ConnectToDB(dsn string) error {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func CreateStudent(student *Student) error {
	return db.Create(student).Error
}

func GetCourseByID(id uint) (*Course, error) {
	var course Course
	err := db.First(&course, id).Error

	return &course, err
}
