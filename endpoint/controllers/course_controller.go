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

type CourseController struct{}

func (cc CourseController) GetAll(c *gin.Context) {
	name := c.Query("name")

	repo := repositories.CourseRepository{}

	courses, err := repo.GetAll(name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func (cc CourseController) GetByID(c *gin.Context) {
	lang := helpers.GetLang(c)

	id, err := helpers.ParseUIntParam(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": i18n.T(lang, "invalid_id"),
		})
		return
	}

	repo := repositories.CourseRepository{}

	course, err := repo.GetByID(uint(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": i18n.T(lang, "course_not_found"),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, course)
}

func (cc CourseController) Create(c *gin.Context) {
	var request dto.CreateCourseRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	course := models.Course{
		Code:    request.Code,
		Name:    request.Name,
		Credits: request.Credits,
	}

	repo := repositories.CourseRepository{}

	if err := repo.Create(&course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, course)
}

func (cc CourseController) Update(c *gin.Context) {
	lang := helpers.GetLang(c)

	id, err := helpers.ParseUIntParam(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": i18n.T(lang, "invalid_id"),
		})
		return
	}

	repo := repositories.CourseRepository{}

	course, err := repo.GetByID(uint(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": i18n.T(lang, "course_not_found"),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var request dto.UpdateCourseRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if request.Code != nil {
		course.Code = *request.Code
	}
	if request.Name != nil {
		course.Name = *request.Name
	}
	if request.Credits != nil {
		course.Credits = *request.Credits
	}

	if err := repo.Update(&course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, course)
}

func (cc CourseController) Delete(c *gin.Context) {
	lang := helpers.GetLang(c)

	id, err := helpers.ParseUIntParam(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": i18n.T(lang, "invalid_id"),
		})
		return
	}

	repo := repositories.CourseRepository{}

	if err := repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": i18n.T(lang, "course_deleted"),
	})
}