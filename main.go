package main

import (
	"backend-crud-user/controller"
	"backend-crud-user/models"
	"backend-crud-user/repository"
	"backend-crud-user/service"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var dsn = "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Profile{})
	db.Migrator().CreateConstraint(models.Profile{}, "User")

	repositoryUser := repository.NewRepositoryUser(db)
	serviceUser := service.NewServiceUser(repositoryUser)
	userController := controller.NewUserController(serviceUser)

	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api")

	api.GET("/", userController.GetData)
	api.GET("/detail/:id", userController.GetUserByID)

	router.Run(":8080")
}
