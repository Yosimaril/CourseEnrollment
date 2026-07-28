package routes

import (
	"yosimaril/CourseEnrollment/controllers"
	"yosimaril/CourseEnrollment/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.RateLimiter())

	authController := controllers.AuthController{}
	userController := controllers.UserController{}
	courseController := controllers.courseController{}
	coursePlanController := controllers.CoursePlanController{}
	coursePlanItemController := controllers.CoursePlanItemController{}

	users := r.Group("/users")
	{
		users.GET("", userController.GetAll)
		users.GET("/:id", userController.GetByID)
		users.POST("", userController.Create)
		users.PUT("/:id", userController.Update)
		users.DELETE("/:id", userController.Delete)
	}

	courses := r.Group("/courses")
	{
		courses.GET("", courseController.GetAll)
		courses.GET("/:id", courseController.GetByID)
		courses.POST("", courseController.Create)
		courses.PUT("/:id", courseController.Update)
		courses.DELETE("/:id", courseController.Delete)
	}

	coursePlans := r.Group("/coursePlans")
	{
		coursePlans.GET("", coursePlanController.GetAll)
		coursePlans.GET("/:id", coursePlanController.GetByID)
		coursePlans.POST("", coursePlanController.Create)
		coursePlans.PUT("/:id", coursePlanController.Update)
		coursePlans.DELETE("/:id", coursePlanController.Delete)
	}

	coursePlanItems := r.Group("/coursePlanItems")
	{
		coursePlanItems.GET("", coursePlanItemController.GetAll)
		coursePlanItems.GET("/:id", coursePlanItemController.GetByID)
		coursePlanItems.POST("", coursePlanItemController.Create)
		coursePlanItems.PUT("/:id", coursePlanItemController.Update)
		coursePlanItems.DELETE("/:id", coursePlanItemController.Delete)
	}

	return r
}
