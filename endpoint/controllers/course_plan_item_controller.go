package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"yosimaril/CourseEnrollment/constants"
	"yosimaril/CourseEnrollment/dto"
	"yosimaril/CourseEnrollment/i18n"
	"yosimaril/CourseEnrollment/models"
	"yosimaril/CourseEnrollment/repositories"
	"yosimaril/CourseEnrollment/utils/helpers"
	"yosimaril/CourseEnrollment/utils/token"
)

type CoursePlanItemController struct{}

func totalCoursePlanCredits(items []models.CoursePlanItem) int {
	total := 0

	for _, item := range items {
		total += item.Course.Credits
	}

	return total
}

func exceedsCoursePlanCreditLimit(maxCredits *int, items []models.CoursePlanItem, additionalCredits int) bool {
	if maxCredits == nil {
		return false
	}

	return totalCoursePlanCredits(items)+additionalCredits > *maxCredits
}

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
	lang := helpers.GetLang(c)
	var request dto.CreateCoursePlanItemRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	courseRepo := repositories.CourseRepository{}
	course, err := courseRepo.GetByID(request.CourseID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Course not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	coursePlanRepo := repositories.CoursePlanRepository{}
	coursePlan, err := coursePlanRepo.GetByID(request.CoursePlanID)
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

	if exceedsCoursePlanCreditLimit(coursePlan.Student.MaxCredits, coursePlan.Items, course.Credits) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": i18n.T(lang, "course_plan_credit_limit_exceeded"),
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

	repositories.InvalidateCachePrefixes("course_plan_item")

	c.JSON(http.StatusCreated, coursePlanItem)
}

func (cpc CoursePlanItemController) AddToPickedCourses(c *gin.Context) {
	lang := helpers.GetLang(c)
	claimsValue, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	claims, ok := claimsValue.(*token.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	if claims.UserRole != constants.RoleStudent {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Forbidden",
		})
		return
	}

	var request struct {
		CourseID uint `json:"course_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	courseRepo := repositories.CourseRepository{}
	course, err := courseRepo.GetByID(request.CourseID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Course not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	userRepo := repositories.UserRepository{}
	student, err := userRepo.GetByID(claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	coursePlanRepo := repositories.CoursePlanRepository{}
	coursePlan, err := coursePlanRepo.GetDraftByStudent(claims.UserID)
	planExists := err == nil
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	if planExists {
		repo := repositories.CoursePlanItemRepository{}
		_, err = repo.GetByID(coursePlan.ID, request.CourseID)
		if err == nil {
			c.JSON(http.StatusConflict, gin.H{
				"message": "Course already added",
			})
			return
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	if exceedsCoursePlanCreditLimit(student.MaxCredits, coursePlan.Items, course.Credits) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": i18n.T(lang, "course_plan_credit_limit_exceeded"),
		})
		return
	}

	if !planExists {
		newCoursePlan := models.CoursePlan{
			StudentID: claims.UserID,
			Status:    constants.CoursePlanDraft,
		}

		if err := coursePlanRepo.Create(&newCoursePlan); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		coursePlan = newCoursePlan
	}

	repo := repositories.CoursePlanItemRepository{}

	coursePlanItem := models.CoursePlanItem{
		CoursePlanID: coursePlan.ID,
		CourseID:     request.CourseID,
		Status:       constants.CoursePlanItemPending,
	}

	if err := repo.Create(&coursePlanItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	repositories.InvalidateCachePrefixes("course_plan_item")

	c.JSON(http.StatusCreated, coursePlanItem)
}

func (cpc CoursePlanItemController) DeleteFromPickedCourses(c *gin.Context) {
	claimsValue, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	claims, ok := claimsValue.(*token.Claims)
	if !ok || claims.UserRole != constants.RoleStudent {
		c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
		return
	}

	courseID, err := helpers.ParseUintParam(c, "course_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid course_id"})
		return
	}

	coursePlanRepo := repositories.CoursePlanRepository{}
	coursePlan, err := coursePlanRepo.GetDraftByStudent(claims.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Draft course plan not found"})
		return
	}

	repo := repositories.CoursePlanItemRepository{}
	if err := repo.Delete(coursePlan.ID, courseID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	repositories.InvalidateCachePrefixes("course_plan_item")

	c.JSON(http.StatusOK, gin.H{"message": "Course removed"})
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

	repositories.InvalidateCachePrefixes("course_plan_item")

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

	repositories.InvalidateCachePrefixes("course_plan_item")

	c.JSON(http.StatusOK, gin.H{
		"message": i18n.T(lang, "course_plan_item_deleted"),
	})
}
