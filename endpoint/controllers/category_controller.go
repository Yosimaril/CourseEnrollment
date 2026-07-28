package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"yosimaril/CourseEnrollment/dto"
	"yosimaril/CourseEnrollment/i18n"
	"yosimaril/CourseEnrollment/models"
	"yosimaril/CourseEnrollment/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryController struct{}

func (cc CategoryController) GetAll(c *gin.Context) {
	keyword := c.Query("keyword")

	repo := repositories.CategoryRepository{}

	categories, err := repo.GetAll(keyword)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (cc CategoryController) GetByID(c *gin.Context) {
	lang := getLang(c)

	id, err := parseID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": i18n.T(lang, "invalid_id"),
		})
		return
	}

	repo := repositories.CategoryRepository{}

	category, err := repo.GetByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (cc CategoryController) Create(c *gin.Context) {
	var request dto.CreateCategoryRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	category := models.Category{
		Name:        request.Name,
		Description: request.Description,
	}

	repo := repositories.CategoryRepository{}

	if err := repo.Create(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, category)
}

func (cc CategoryController) Update(c *gin.Context) {
	lang := getLang(c)

	id, err := parseID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	repo := repositories.CategoryRepository{}

	category, err := repo.GetByID(uint(id))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": i18n.T(lang, "category_not_found"),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var request dto.UpdateCategoryRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	category.Name = request.Name
	category.Description = request.Description

	if err := repo.Update(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (cc CategoryController) Delete(c *gin.Context) {
	lang := getLang(c)

	id, err := parseID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
		})
		return
	}

	repo := repositories.CategoryRepository{}

	if err := repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": i18n.T(lang, "category_deleted"),
	})
}

func parseID(c *gin.Context) (uint, error) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	return uint(id), err
}

func getLang(c *gin.Context) string {
	lang := c.GetHeader("Accept-Language")

	if lang == "" {
		lang = "en"
	}

	return lang
}
