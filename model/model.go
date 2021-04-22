package model

import "gin-gorm-postgres/database"

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	ClassName string `gorm:"type:varchar(90)"`
	Teachers []Teacher `gorm:"many2many:class_teachers;"`
	Student []Student
}

type Student struct {
	gorm.Model
	StudentName string `gorm:"type:varchar(90)"`
	ClassID     uint
	IDCard      IDCard
	Teachers    []Teacher `gorm:"many2many:student_teachers;"`
}

type IDCard struct {
	gorm.Model
	StudentID uint
	Num       int
}

type Teacher struct {
	gorm.Model
	TeacherName string `gorm:"type:varchar(90)"`
	TeachName string `gorm:"type:varchar(90);default:''"`
	Students    []Student  `gorm:"many2many:student_teachers;"`
	Class []Class `gorm:"many2many:class_teachers;"`
}

func Migrate() {
	database.DB.AutoMigrate(&Class{}, &Student{}, &IDCard{}, &Teacher{})

}
