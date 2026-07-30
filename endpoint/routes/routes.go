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

	// Authentication
	auth := r.Group("/auth")
	{
		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Register)
	}

	api := r.Group("/")
	api.Use(middleware.JWT())

	{
		// Users
		users := api.Group("/users")
		{
			users.GET("", userController.GetAll)
			users.GET("/:id", userController.GetByID)
			users.POST("", userController.Create)
			users.PUT("/:id", userController.Update)
			users.DELETE("/:id", userController.Delete)
		}

		// Courses (Read-only for authenticated users)
		courses := api.Group("/courses")
		{
			courses.GET("", courseController.GetAll)
			courses.GET("/:id", courseController.GetByID)
		}

		// Course Plans
		coursePlans := api.Group("/course-plans")
		{
			coursePlans.GET("", coursePlanController.GetAll)
			coursePlans.GET("/:id", coursePlanController.GetByID)
			coursePlans.POST("", coursePlanController.Create)
			coursePlans.PUT("/:id", coursePlanController.Update)
			coursePlans.DELETE("/:id", coursePlanController.Delete)
		}

		// Course Plan Items
		coursePlanItems := api.Group("/course-plan-items")
		{
			coursePlanItems.GET("", coursePlanItemController.GetAll)
			coursePlanItems.GET("/:id", coursePlanItemController.GetByID)
			coursePlanItems.POST("", coursePlanItemController.Create)
			coursePlanItems.PUT("/:id", coursePlanItemController.Update)
			coursePlanItems.DELETE("/:id", coursePlanItemController.Delete)
		}

		// Student
		student := api.Group("/student")
		student.Use(middleware.RequireRole(constants.RoleStudent))
		{
			student.GET("/course-plan", coursePlanController.GetCurrentCoursePlan)
			student.GET("/course-plans", coursePlanController.GetCoursePlans)

			student.POST("/course-plan/submit", coursePlanController.SubmitCurrentCoursePlan)
			student.DELETE("/course-plans/:id", coursePlanController.CancelCoursePlan)

			student.POST("/picked-courses", coursePlanItemController.AddToPickedCourses)
			student.DELETE("/picked-courses/:courseId", coursePlanItemController.DeleteFromPickedCourses)
		}

		// Admin
		admin := api.Group("/admin")
		admin.Use(middleware.RequireRole(constants.RoleAdmin))
		{
			// Students
			adminStudents := admin.Group("/students")
			{
				adminStudents.GET("", userController.GetAll)
				adminStudents.GET("/:id", userController.GetByID)
				adminStudents.POST("", userController.Create)
				adminStudents.PUT("/:id", userController.Update)
				adminStudents.DELETE("/:id", userController.Delete)
			}

			// Courses
			adminCourses := admin.Group("/courses")
			{
				adminCourses.GET("", courseController.GetAll)
				adminCourses.GET("/:id", courseController.GetByID)
				adminCourses.POST("", courseController.Create)
				adminCourses.PUT("/:id", courseController.Update)
				adminCourses.DELETE("/:id", courseController.Delete)
			}

			// Course Plans
			adminCoursePlans := admin.Group("/course-plans")
			{
				adminCoursePlans.GET("", coursePlanController.GetAll)
				adminCoursePlans.PUT("/:id/review", coursePlanController.ReviewCoursePlan)
			}
		}
	}

	return r
}
