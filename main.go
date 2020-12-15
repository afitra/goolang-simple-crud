package main

import (
	"fmt"
	"log"
	"net/http"
	"simpleCrudGolang/auth"
	"simpleCrudGolang/handler"
	"simpleCrudGolang/helper"
	"simpleCrudGolang/user"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
	apiUser.GET("/detail", authMiddlewere(authService, userService), userHandler.GetUserByID)
	apiUser.POST("/profile", authMiddlewere(authService, userService), userHandler.UploadProfile)
	apiUser.PUT("/:id", authMiddlewere(authService, userService), userHandler.UpdateUser)
	PORT := helper.GoDotEnvVariable("PORT")

	router.Run(fmt.Sprintf(":%s", PORT))
}

func authMiddlewere(authService auth.Service, userService user.Service) gin.HandlerFunc {
	// Midleware
	// 1. ambil nilai header authorization ->> bearer token
	// 2. dari header authorization , kita ambil token saja
	// 3. validasi tokennya
	// 4. ambil nilai user_id
	// 5. ambil user di db berdasar user_id
	// 6. set context isinya user
	return func(c *gin.Context) { // -->> gin handler adl fungsi yg punya param gin.Context

		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Auth") {

			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")

		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]

		}
		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)

		if err != nil {

			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
