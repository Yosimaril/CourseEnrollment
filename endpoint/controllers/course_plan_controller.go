package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"yosimaril/CourseEnrollment/config"
	"yosimaril/CourseEnrollment/constants"
	"yosimaril/CourseEnrollment/dto"
	"yosimaril/CourseEnrollment/i18n"
	"yosimaril/CourseEnrollment/models"
	"yosimaril/CourseEnrollment/repositories"
	"yosimaril/CourseEnrollment/utils/helpers"
	"yosimaril/CourseEnrollment/utils/token"
)

type CoursePlanController struct{}

func (cpc CoursePlanController) GetAll(c *gin.Context) {
	lang := helpers.GetLang(c)
	var studentID uint
	status := c.Query("status")

	studentIDStr := c.Query("student_id")
	if studentIDStr != "" {
		parsedID, err := helpers.ParseUintQuery(c, "student_id")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": i18n.T(lang, "invalid_student_id"),
			})
			return
		}
		studentID = parsedID
	}

	repo := repositories.CoursePlanRepository{}

	coursePlans, err := repo.GetAll(studentID, status)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, coursePlans)
}

func (cpc CoursePlanController) GetCurrentCoursePlan(c *gin.Context) {
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

	repo := repositories.CoursePlanRepository{}
	coursePlan, err := repo.GetDraftByStudent(claims.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Draft course plan not found"})
		return
	}

	c.JSON(http.StatusOK, coursePlan)
}

func (cpc CoursePlanController) GetCoursePlans(c *gin.Context) {
	claimsValue, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	claims, ok := claimsValue.(*token.Claims)
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
		return
	}

	repo := repositories.CoursePlanRepository{}
	coursePlans, err := repo.GetAll(claims.UserID, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, coursePlans)
}

func (cpc CoursePlanController) SubmitCurrentCoursePlan(c *gin.Context) {
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

	repo := repositories.CoursePlanRepository{}
	coursePlan, err := repo.GetDraftByStudent(claims.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Draft course plan not found"})
		return
	}

	if len(coursePlan.Items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Course plan is empty"})
		return
	}

	coursePlan.Status = constants.CoursePlanSubmitted
	if err := repo.Update(&coursePlan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	repositories.InvalidateCachePrefixes("course_plan_item")

	c.JSON(http.StatusOK, coursePlan)
}

func (cpc CoursePlanController) CancelCoursePlan(c *gin.Context) {
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

	id, err := helpers.ParseUintParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	repo := repositories.CoursePlanRepository{}
	coursePlan, err := repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Course plan not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if coursePlan.StudentID != claims.UserID {
		c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
		return
	}

	if coursePlan.Status == constants.CoursePlanApproved {
		c.JSON(http.StatusForbidden, gin.H{"message": "Approved course plan cannot be cancelled"})
		return
	}

	if err := repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	repositories.InvalidateCachePrefixes("course_plan_item")

	c.JSON(http.StatusOK, gin.H{"message": "Course plan cancelled"})
}

func (cpc CoursePlanController) ReviewCoursePlan(c *gin.Context) {
	var request dto.ReviewCoursePlanRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, err := helpers.ParseUintParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	repo := repositories.CoursePlanRepository{}
	coursePlan, err := repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "Course plan not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if len(request.CourseIDs) == 0 {
		for i := range coursePlan.Items {
			coursePlan.Items[i].Status = request.ItemStatus
		}
	} else {
		selected := map[uint]struct{}{}
		for _, courseID := range request.CourseIDs {
			selected[courseID] = struct{}{}
		}
		for i := range coursePlan.Items {
			if _, ok := selected[coursePlan.Items[i].CourseID]; ok {
				coursePlan.Items[i].Status = request.ItemStatus
			}
		}
	}

	approvedCount := 0
	rejectedCount := 0
	pendingCount := 0

	for _, item := range coursePlan.Items {
		switch item.Status {
		case constants.CoursePlanItemApproved:
			approvedCount++
		case constants.CoursePlanItemRejected:
			rejectedCount++
		default:
			pendingCount++
		}
	}

	switch {
	case len(coursePlan.Items) == 0:
		coursePlan.Status = constants.CoursePlanDraft
	case approvedCount == len(coursePlan.Items):
		coursePlan.Status = constants.CoursePlanApproved
	case rejectedCount == len(coursePlan.Items):
		coursePlan.Status = constants.CoursePlanRejected
	case approvedCount == 0 && rejectedCount == 0:
		coursePlan.Status = constants.CoursePlanSubmitted
	default:
		coursePlan.Status = constants.CoursePlanPartiallyApproved
	}

	if err := repo.Update(&coursePlan); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	for i := range coursePlan.Items {
		item := coursePlan.Items[i]
		if err := config.DB.Model(&models.CoursePlanItem{}).Where("course_plan_id = ? AND course_id = ?", coursePlan.ID, item.CourseID).Updates(map[string]any{"status": item.Status}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	repositories.InvalidateCachePrefixes("course_plan_item")

	updatedPlan, err := repo.GetByID(coursePlan.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedPlan)
}

func (cpc CoursePlanController) GetByID(c *gin.Context) {
	lang := helpers.GetLang(c)

	id, err := helpers.ParseUintParam(c, "id")

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

	repositories.InvalidateCachePrefixes("course_plan_item")

	c.JSON(http.StatusCreated, coursePlan)
}

func (cpc CoursePlanController) Update(c *gin.Context) {
	lang := helpers.GetLang(c)

	id, err := helpers.ParseUintParam(c, "id")

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

	repositories.InvalidateCachePrefixes("course_plan_item")

	c.JSON(http.StatusOK, coursePlan)
}

func (cpc CoursePlanController) Delete(c *gin.Context) {
	lang := helpers.GetLang(c)

	id, err := helpers.ParseUintParam(c, "id")

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

	repositories.InvalidateCachePrefixes("course_plan_item")

	c.JSON(http.StatusOK, gin.H{
		"message": i18n.T(lang, "course_plan_deleted"),
	})
}
