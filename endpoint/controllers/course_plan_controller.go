package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"yosimaril/CourseEnrollment/endpoint/dto"
	"yosimaril/CourseEnrollment/i18n"
	"yosimaril/CourseEnrollment/models"
	"yosimaril/CourseEnrollment/repositories"
	"yosimaril/CourseEnrollment/utils/helpers"
)

type CoursePlanController struct{}

func (cpc CoursePlanController) GetAll(c *gin.Context) {
	lang := helpers.GetLang(c)
	studentIDStr := c.Query("student_id")
	var studentID uint

	if studentIDStr != "" {
		parsedID, err := helpers.ParseUInt(studentIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": i18n.T(lang, "invalid_student_id"),
			})
			return
		}
		studentID = parsedID
	}

	repo := repositories.CoursePlanRepository{}

	coursePlans, err := repo.GetAll(studentID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, coursePlans)
}

func (cpc CoursePlanController) GetByID(c *gin.Context) {
	lang := helpers.GetLang(c)

	id, err := helpers.ParseUIntParam(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": i18n.T(lang, "invalid_id"),
		})
		return
	}

	repo := repositories.CoursePlanRepository{}

	coursePlan, err := repo.GetByID(uint(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": i18n.T(lang, "course_plan_not_found"),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, coursePlan)
}

func (cpc CoursePlanController) Create(c *gin.Context) {
	var request dto.CreateCoursePlanRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	coursePlan := models.CoursePlan{
		StudentID: request.StudentID,
		Status:    request.Status,
	}

	repo := repositories.CoursePlanRepository{}

	if err := repo.Create(&coursePlan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, coursePlan)
}

func (cpc CoursePlanController) Update(c *gin.Context) {
	lang := helpers.GetLang(c)

	id, err := helpers.ParseUIntParam(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": i18n.T(lang, "invalid_id"),
		})
		return
	}

	repo := repositories.CoursePlanRepository{}

	coursePlan, err := repo.GetByID(uint(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": i18n.T(lang, "course_plan_not_found"),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var request dto.UpdateCoursePlanRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if request.StudentID != nil {
		coursePlan.StudentID = *request.StudentID
	}
	if request.Status != nil {
		coursePlan.Status = *request.Status
	}

	if err := repo.Update(&coursePlan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, coursePlan)
}

func (cpc CoursePlanController) Delete(c *gin.Context) {
	lang := helpers.GetLang(c)

	id, err := helpers.ParseUIntParam(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": i18n.T(lang, "invalid_id"),
		})
		return
	}

	repo := repositories.CoursePlanRepository{}

	if err := repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": i18n.T(lang, "course_plan_deleted"),
	})
}