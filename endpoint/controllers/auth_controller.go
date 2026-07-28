package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"yosimaril/CourseEnrollment/dto"
	"yosimaril/CourseEnrollment/i18n"
	"yosimaril/CourseEnrollment/models"
	"yosimaril/CourseEnrollment/repositories"
	"yosimaril/CourseEnrollment/utils/helpers"
	"yosimaril/CourseEnrollment/utils/token"
)

type AuthController struct{}

func (ac AuthController) Login(c *gin.Context) {
	lang := helpers.GetLang(c)
	var request dto.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userRepo := repositories.UserRepository{}
	user, err := userRepo.GetByEmail(request.Email)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": i18n.T(lang, "invalid_credentials"),
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": i18n.T(lang, "invalid_credentials"),
		})
		return
	}

	tokenString, err := token.GenerateToken(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": i18n.T(lang, "failed_to_generate_token"),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": i18n.T(lang, "login_successful"),
		"token":   tokenString,
	})
}

func (ac AuthController) Register(c *gin.Context) {
	lang := helpers.GetLang(c)
	var request dto.RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userRepo := repositories.UserRepository{}

	_, err := userRepo.GetByEmail(request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": i18n.T(lang, "email_already_registered"),
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": i18n.T(lang, "failed_to_hash_password"),
		})
		return
	}

	user := models.User{
		Username:   request.Username,
		Email:      request.Email,
		Password:   string(hashedPassword),
		Role:       request.Role,
		Nrp:        request.Nrp,
		MaxCredits: request.MaxCredits,
	}

	if err := userRepo.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": i18n.T(lang, "registration_successful"),
		"user":    user,
	})
}
