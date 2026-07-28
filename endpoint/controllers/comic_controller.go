package controllers

import (
	"yosimaril/CourseEnrollment/repositories"

	"github.com/gin-gonic/gin"
)

type ComicController struct{}

func (cc ComicController) GetAll(c *gin.Context) {
	keyword := c.Query("keyword")

	repo := repositories.CategoryRepository{}

	categories, err := repo.GetAll(keyword)

	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, categories)
}
