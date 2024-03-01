package main

import (
	"fmt"
	"gorm.io/gorm"
)

func getStudentsByDepartment(db *gorm.DB, departmentID uint) ([]string, error) {
	var students []Student
	var studentNames []string

	if err := db.Where("department_id = ?", departmentID).Find(&students).Error; err != nil {
		return nil, err
	}

	for _, student := range students {
		studentNames = append(studentNames, student.Name)
	}

	return studentNames, nil
}

func getCoursesByInstructor(db *gorm.DB, instructorID uint) ([]string, error) {
	var courses []Course
	var courseNames []string

	if err := db.Where("instructor_id = ?", instructorID).Find(&courses).Error; err != nil {
		return nil, err
	}

	for _, course := range courses {
		courseNames = append(courseNames, course.Name)
	}

	return courseNames, nil
}

func getEnrollmentsByStudent(db *gorm.DB, studentID uint) ([]string, error) {
	var enrollments []Enrollment
	var enrollmentDetails []string

	if err := db.Where("student_id = ?", studentID).Preload("Course").Find(&enrollments).Error; err != nil {
		return nil, err
	}

	for _, enrollment := range enrollments {
		enrollmentDetails = append(enrollmentDetails, fmt.Sprintf("Course: %s, Grade: %s", enrollment.CourseID, enrollment.Grade))
	}

	return enrollmentDetails, nil
}

// Soft deletes
func softDeleteStudent(db *gorm.DB, id uint) error {
	if err := db.Unscoped().Delete(&Student{}, id).Error; err != nil {
		return err
	}
	return nil
}
func getAllStudents(db *gorm.DB) ([]Student, error) {
	var students []Student
	if err := db.Unscoped().Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

// Custom queries
func getStudentsEnrolledInCourse(db *gorm.DB, courseID uint) ([]Student, error) {
	var enrolledStudents []Student
	if err := db.Table("students").
		Joins("INNER JOIN enrollments ON students.id = enrollments.student_id").
		Where("enrollments.course_id = ?", courseID).
		Find(&enrolledStudents).Error; err != nil {
		return nil, err
	}
	return enrolledStudents, nil
}
func getTotalStudentsInDepartment(db *gorm.DB, departmentID uint) (int64, error) {
	var totalStudents int64
	if err := db.Model(&Student{}).Where("department_id = ?", departmentID).Count(&totalStudents).Error; err != nil {
		return 0, err
	}
	return totalStudents, nil
}
