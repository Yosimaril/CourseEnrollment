package handler

import (
	"CourseEnrollment/endpoint/dto"
	"CourseEnrollment/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	studentService service.StudentService
}

func NewStudentHandler(studentService service.StudentService) *StudentHandler {
	return &StudentHandler{studentService: studentService}
}

// CreateStudent godoc
// @Summary Create a new student
// @Description Create a new student with the input payload
// @Tags students
// @Accept  json
// @Produce  json
// @Param student body dto.StudentCreateRequest true "Create student"
// @Success 201 {object} dto.StudentResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /students [post]
func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var req dto.StudentCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student, err := h.studentService.CreateStudent(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, student)
}

// GetStudentByID godoc
// @Summary Get a student by ID
// @Description Get a student by ID
// @Tags students
// @Produce  json
// @Param id path int true "Student ID"
// @Success 200 {object} dto.StudentResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /students/{id} [get]
func (h *StudentHandler) GetStudentByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	student, err := h.studentService.GetStudentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}

// GetAllStudents godoc
// @Summary Get all students
// @Description Get all students
// @Tags students
// @Produce  json
// @Success 200 {array} dto.StudentResponse
// @Failure 500 {object} map[string]string
// @Router /students [get]
func (h *StudentHandler) GetAllStudents(c *gin.Context) {
	students, err := h.studentService.GetAllStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}

// UpdateStudent godoc
// @Summary Update a student
// @Description Update a student with the input payload
// @Tags students
// @Accept  json
// @Produce  json
// @Param id path int true "Student ID"
// @Param student body dto.StudentUpdateRequest true "Update student"
// @Success 200 {object} dto.StudentResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /students/{id} [put]
func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	var req dto.StudentUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student, err := h.studentService.UpdateStudent(uint(id), &req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}

// DeleteStudent godoc
// @Summary Delete a student
// @Description Delete a student by ID
// @Tags students
// @Produce  json
// @Param id path int true "Student ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /students/{id} [delete]
func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	if err := h.studentService.DeleteStudent(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
