package main

import (
	"fmt"
	"yosimaril/CourseEnrollment/config"
	"yosimaril/CourseEnrollment/i18n"
	"yosimaril/CourseEnrollment/middleware"
	"yosimaril/CourseEnrollment/routes"
)

func main() {
	config.ConnectDatabase()
	config.ConnectRedis()
	config.DB = config.DB.Debug()

	fmt.Println("Connected!")

	if err := i18n.LoadLanguage("en", "locales/en.json"); err != nil {
		panic(err)
	}

	if err := i18n.LoadLanguage("id", "locales/id.json"); err != nil {
		panic(err)
	}

	//config.DB.AutoMigrate(
	//	&models.User{},
	//	&models.Category{},
	//	&models.Comic{},
	//	&models.Chapter{},
	//	&models.ChapterPage{},
	//)

	go middleware.CleanupVisitors()

	r := routes.SetupRouter()
	r.Run(":8000")
}
