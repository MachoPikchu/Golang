package main

import (
	"gorm.io/gorm"
	"time"
)

type Student struct {
	ID           uint
	Name         string
	DepartmentID uint
	Enrollments  []Enrollment
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"` // Soft delete marker

}

type Department struct {
	ID   uint
	Name string
}

type Course struct {
	ID           uint
	Name         string
	InstructorID uint
	DepartmentID uint
	Enrollments  []Enrollment
}

type Enrollment struct {
	ID        uint
	StudentID uint
	CourseID  uint
	Grade     string
}

type Instructor struct {
	ID        uint
	Name      string
	Courses   []Course
	UpdatedAt time.Time
}

func (student *Student) BeforeCreate(tx *gorm.DB) (err error) {
	student.CreatedAt = time.Now()
	student.UpdatedAt = time.Now()
	return
}

func (instructor *Instructor) BeforeUpdate(tx *gorm.DB) (err error) {
	instructor.UpdatedAt = time.Now()
	return
}
