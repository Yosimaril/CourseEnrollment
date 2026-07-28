package service

import (
	"CourseEnrollment/endpoint/dto"
	"CourseEnrollment/internal/model"
	"CourseEnrollment/internal/repository"
	"errors"
)

type StudentService interface {
	CreateStudent(req *dto.StudentCreateRequest) (*dto.StudentResponse, error)
	GetStudentByID(id uint) (*dto.StudentResponse, error)
	GetAllStudents() ([]dto.StudentResponse, error)
	UpdateStudent(id uint, req *dto.StudentUpdateRequest) (*dto.StudentResponse, error)
	DeleteStudent(id uint) error
}

type studentService struct {
	studentRepo repository.StudentRepository
}

func NewStudentService(studentRepo repository.StudentRepository) StudentService {
	return &studentService{studentRepo: studentRepo}
}

func (s *studentService) CreateStudent(req *dto.StudentCreateRequest) (*dto.StudentResponse, error) {
	student := &model.Student{
		Name:  req.Name,
		Email: req.Email,
		Major: req.Major,
	}
	if err := s.studentRepo.CreateStudent(student); err != nil {
		return nil, err
	}
	return &dto.StudentResponse{
		ID:        student.ID,
		Name:      student.Name,
		Email:     student.Email,
		Major:     student.Major,
		CreatedAt: student.CreatedAt,
		UpdatedAt: student.UpdatedAt,
	}, nil
}

func (s *studentService) GetStudentByID(id uint) (*dto.StudentResponse, error) {
	student, err := s.studentRepo.GetStudentByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.StudentResponse{
		ID:        student.ID,
		Name:      student.Name,
		Email:     student.Email,
		Major:     student.Major,
		CreatedAt: student.CreatedAt,
		UpdatedAt: student.UpdatedAt,
	}, nil
}

func (s *studentService) GetAllStudents() ([]dto.StudentResponse, error) {
	students, err := s.studentRepo.GetAllStudents()
	if err != nil {
		return nil, err
	}
	var studentResponses []dto.StudentResponse
	for _, student := range students {
		studentResponses = append(studentResponses, dto.StudentResponse{
			ID:        student.ID,
			Name:      student.Name,
			Email:     student.Email,
			Major:     student.Major,
			CreatedAt: student.CreatedAt,
			UpdatedAt: student.UpdatedAt,
		})
	}
	return studentResponses, nil
}

func (s *studentService) UpdateStudent(id uint, req *dto.StudentUpdateRequest) (*dto.StudentResponse, error) {
	student, err := s.studentRepo.GetStudentByID(id)
	if err != nil {
		return nil, err
	}

	student.Name = req.Name
	student.Email = req.Email
	student.Major = req.Major

	if err := s.studentRepo.UpdateStudent(student); err != nil {
		return nil, err
	}
	return &dto.StudentResponse{
		ID:        student.ID,
		Name:      student.Name,
		Email:     student.Email,
		Major:     student.Major,
		CreatedAt: student.CreatedAt,
		UpdatedAt: student.UpdatedAt,
	}, nil
}

func (s *studentService) DeleteStudent(id uint) error {
	_, err := s.studentRepo.GetStudentByID(id)
	if err != nil {
		return errors.New("student not found")
	}
	return s.studentRepo.DeleteStudent(id)
}
