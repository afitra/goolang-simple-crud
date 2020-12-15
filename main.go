package main

import (
	"fmt"
	"log"
	"simpleCrudGolang/auth"
	"simpleCrudGolang/handler"
	"simpleCrudGolang/helper"
	"simpleCrudGolang/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//
func main() {
	databaseName := helper.GoDotEnvVariable("DATABASE_NAME")
	dsn := fmt.Sprintf("root:root@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", databaseName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	db.Migrator().CreateTable(user.User{})
	fmt.Println("koneksi DB berhasil *******")

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	fmt.Println(userService, ">>>>>>")
	router := gin.Default()
	router.Static("/profile", "./profile") // kiri routenya , kanan directory folder

	apiUser := router.Group("/api/v1/users")

	apiUser.POST("/register", userHandler.RegisterUser)
	apiUser.POST("/login", userHandler.Login)
	PORT := helper.GoDotEnvVariable("PORT")

	router.Run(fmt.Sprintf(":%s", PORT))
}
