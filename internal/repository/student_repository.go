package repository

import (
	"CourseEnrollment/internal/model"
	"gorm.io/gorm"
)

type StudentRepository interface {
	CreateStudent(student *model.Student) error
	GetStudentByID(id uint) (*model.Student, error)
	GetAllStudents() ([]model.Student, error)
	UpdateStudent(student *model.Student) error
	DeleteStudent(id uint) error
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db: db}
}

func (r *studentRepository) CreateStudent(student *model.Student) error {
	return r.db.Create(student).Error
}

func (r *studentRepository) GetStudentByID(id uint) (*model.Student, error) {
	var student model.Student
	if err := r.db.First(&student, id).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *studentRepository) GetAllStudents() ([]model.Student, error) {
	var students []model.Student
	if err := r.db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (r *studentRepository) UpdateStudent(student *model.Student) error {
	return r.db.Save(student).Error
}

func (r *studentRepository) DeleteStudent(id uint) error {
	return r.db.Delete(&model.Student{}, id).Error
}
