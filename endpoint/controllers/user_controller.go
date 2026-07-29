package controllers

import (
	"errors"
	"net/http"

	"yosimaril/CourseEnrollment/dto"
	"yosimaril/CourseEnrollment/i18n"
	"yosimaril/CourseEnrollment/models"
	"yosimaril/CourseEnrollment/repositories"
	"yosimaril/CourseEnrollment/utils/helpers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct{}

func (uc UserController) GetAll(c *gin.Context) {
	username := c.Query("username")
	role := c.Query("role")

	repo := repositories.UserRepository{}

	users, err := repo.GetAll(username, role)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (uc UserController) GetByID(c *gin.Context) {
	lang := helpers.GetLang(c)

	id, err := helpers.ParseUintParam(c, "id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": i18n.T(lang, "invalid_id"),
		})
		return
	}

	repo := repositories.UserRepository{}

	user, err := repo.GetByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc UserController) Create(c *gin.Context) {
	var request dto.CreateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user := models.User{
		Username:   request.Username,
		Email:      request.Email,
		Password:   request.Password,
		Role:       request.Role,
		Nrp:        request.Nrp,
		MaxCredits: request.MaxCredits,
	}

	repo := repositories.UserRepository{}

	if err := repo.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	repositories.InvalidateCachePrefixes("user")

	c.JSON(http.StatusCreated, user)
}

func (uc UserController) Update(c *gin.Context) {
	lang := helpers.GetLang(c)

	id, err := helpers.ParseUintParam(c, "id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": i18n.T(lang, "invalid_id"),
		})
		return
	}

	repo := repositories.UserRepository{}

	user, err := repo.GetByID(uint(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": i18n.T(lang, "user_not_found"),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var request dto.UpdateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if request.Username != "" {
		user.Username = request.Username
	}
	if request.Email != "" {
		user.Email = request.Email
	}
	if request.Role != "" {
		user.Role = request.Role
	}
	if request.Nrp != nil {
		user.Nrp = request.Nrp
	}
	if request.MaxCredits != nil {
		user.MaxCredits = request.MaxCredits
	}

	if err := repo.Update(&user); err != nil { 
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	repositories.InvalidateCachePrefixes("user")

	c.JSON(http.StatusOK, user)
}

func (uc UserController) Delete(c *gin.Context) {
	lang := helpers.GetLang(c)

	id, err := helpers.ParseUintParam(c, "id")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": i18n.T(lang, "invalid_id"),
		})
		return
	}

	repo := repositories.UserRepository{}

	if err := repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	repositories.InvalidateCachePrefixes("user")

	c.JSON(http.StatusOK, gin.H{
		"message": i18n.T(lang, "user_deleted"),
	})
}
