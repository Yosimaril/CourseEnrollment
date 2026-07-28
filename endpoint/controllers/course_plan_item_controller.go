package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"yosimaril/CourseEnrollment/dto"
	"yosimaril/CourseEnrollment/i18n"
	"yosimaril/CourseEnrollment/models"
	"yosimaril/CourseEnrollment/repositories"
	"yosimaril/CourseEnrollment/utils/helpers"
)

type CoursePlanItemController struct{}

func (cpc CoursePlanItemController) GetAll(c *gin.Context) {
	lang := helpers.GetLang(c)
	var coursePlanID, courseID uint
	var err error

	coursePlanIDStr := c.Query("course_plan_id")
	if coursePlanIDStr != "" {
		coursePlanID, err = helpers.ParseUintQuery(c, "course_plan_id")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": i18n.T(lang, "invalid_course_plan_id"),
			})
			return
		}
	}

	courseIDStr := c.Query("course_id")
	if courseIDStr != "" {
		courseID, err = helpers.ParseUintQuery(c, "course_id")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": i18n.T(lang, "invalid_course_id"),
			})
			return
		}
	}

	repo := repositories.CoursePlanItemRepository{}

	coursePlanItems, err := repo.GetAll(coursePlanID, courseID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, coursePlanItems)
}

func (cpc CoursePlanItemController) GetByID(c *gin.Context) {
	lang := helpers.GetLang(c)

	coursePlanID, err := helpers.ParseUintParam(c, "course_plan_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": i18n.T(lang, "invalid_course_plan_id"),
		})
		return
	}

	courseID, err := helpers.ParseUintParam(c, "course_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": i18n.T(lang, "invalid_course_id"),
		})
		return
	}

	repo := repositories.CoursePlanItemRepository{}

	coursePlanItem, err := repo.GetByID(coursePlanID, courseID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": i18n.T(lang, "course_plan_item_not_found"),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, coursePlanItem)
}

func (cpc CoursePlanItemController) Create(c *gin.Context) {
	var request dto.CreateCoursePlanItemRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	coursePlanItem := models.CoursePlanItem{
		CoursePlanID: request.CoursePlanID,
		CourseID:     request.CourseID,
		Status:       request.Status,
		Remarks:      request.Remarks,
	}

	repo := repositories.CoursePlanItemRepository{}

	if err := repo.Create(&coursePlanItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, coursePlanItem)
}

func (cpc CoursePlanItemController) Update(c *gin.Context) {
	lang := helpers.GetLang(c)

	coursePlanID, err := helpers.ParseUintParam(c, "course_plan_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": i18n.T(lang, "invalid_course_plan_id"),
		})
		return
	}

	courseID, err := helpers.ParseUintParam(c, "course_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": i18n.T(lang, "invalid_course_id"),
		})
		return
	}

	repo := repositories.CoursePlanItemRepository{}

	coursePlanItem, err := repo.GetByID(coursePlanID, courseID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": i18n.T(lang, "course_plan_item_not_found"),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var request dto.UpdateCoursePlanItemRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if request.Status != nil {
		coursePlanItem.Status = *request.Status
	}
	if request.Remarks != nil {
		coursePlanItem.Remarks = request.Remarks
	}

	if err := repo.Update(&coursePlanItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, coursePlanItem)
}

func (cpc CoursePlanItemController) Delete(c *gin.Context) {
	lang := helpers.GetLang(c)

	coursePlanID, err := helpers.ParseUintParam(c, "course_plan_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": i18n.T(lang, "invalid_course_plan_id"),
		})
		return
	}

	courseID, err := helpers.ParseUintParam(c, "course_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": i18n.T(lang, "invalid_course_id"),
		})
		return
	}

	repo := repositories.CoursePlanItemRepository{}

	if err := repo.Delete(coursePlanID, courseID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": i18n.T(lang, "course_plan_item_deleted"),
	})
}