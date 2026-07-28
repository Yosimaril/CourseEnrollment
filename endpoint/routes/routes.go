package routes

import (
	"time"

	"yosimaril/CourseEnrollment/constants"
	"yosimaril/CourseEnrollment/controllers"
	"yosimaril/CourseEnrollment/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Global middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(middleware.RateLimiter())

	authController := controllers.AuthController{}
	userController := controllers.UserController{}
	courseController := controllers.CourseController{}
	coursePlanController := controllers.CoursePlanController{}
	coursePlanItemController := controllers.CoursePlanItemController{}

	auths := r.Group("/auth")
	{
		auths.POST("/login", authController.Login)
		auths.POST("/register", authController.Register)
	}

	api := r.Group("/")
	api.Use(middleware.JWT())

	{
		users := api.Group("/users")
		{
			users.GET("", userController.GetAll)
			users.GET("/:id", userController.GetByID)
			users.POST("", userController.Create)
			users.PUT("/:id", userController.Update)
			users.DELETE("/:id", userController.Delete)
		}

		courses := api.Group("/courses")
		{
			courses.GET("", courseController.GetAll)
			courses.GET("/:id", courseController.GetByID)
		}

		admin := api.Group("/admin")
		admin.Use(middleware.RequireRole(constants.RoleAdmin))
		{
			adminCourses := admin.Group("/courses")
			{
				adminCourses.GET("", courseController.GetAll)
				adminCourses.GET("/:id", courseController.GetByID)
				adminCourses.POST("", courseController.Create)
				adminCourses.PUT("/:id", courseController.Update)
				adminCourses.DELETE("/:id", courseController.Delete)
			}
		}

		coursePlans := api.Group("/coursePlans")
		{
			coursePlans.GET("", coursePlanController.GetAll)
			coursePlans.GET("/:id", coursePlanController.GetByID)
			coursePlans.POST("", coursePlanController.Create)
			coursePlans.PUT("/:id", coursePlanController.Update)
			coursePlans.DELETE("/:id", coursePlanController.Delete)
		}

		coursePlanItems := api.Group("/coursePlanItems")
		{
			coursePlanItems.GET("", coursePlanItemController.GetAll)
			coursePlanItems.GET("/:id", coursePlanItemController.GetByID)
			coursePlanItems.POST("", coursePlanItemController.Create)
			coursePlanItems.PUT("/:id", coursePlanItemController.Update)
			coursePlanItems.DELETE("/:id", coursePlanItemController.Delete)
		}

		student := api.Group("/student")
		student.Use(middleware.RequireRole(constants.RoleStudent))
		{
			student.POST("/picked-courses", coursePlanItemController.AddToPickedCourses)
		}
	}

	return r
}
