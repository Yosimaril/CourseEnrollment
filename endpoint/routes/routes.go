package routes

import (
	"yosimaril/CourseEnrollment/controllers"
	"yosimaril/CourseEnrollment/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.RateLimiter())

	categoryController := controllers.CategoryController{}
	//userController := controllers.UserController{}
	//comicController := controllers.ComicController{}

	categories := r.Group("/categories")
	{
		categories.GET("", categoryController.GetAll)
		categories.GET("/:id", categoryController.GetByID)
		categories.POST("", categoryController.Create)
		categories.PUT("/:id", categoryController.Update)
		categories.DELETE("/:id", categoryController.Delete)
	}

	return r
}
