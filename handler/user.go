package handler

import (
	"fmt"
	"net/http"
	"simpleCrudGolang/auth"
	"simpleCrudGolang/helper"
	"simpleCrudGolang/user"
	"time"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Register Account failed", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.ApiResponse("Register Account failed", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)

	if err != nil {
		response := helper.ApiResponse("Register Account failed", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatOneUser(newUser, token)

	response := helper.ApiResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	// 1. user memasukan email dan password
	// 2. input ditangkap handler
	// 3. mapping dari input user ke input struct
	// 4. struct input di parsing ke bentuk service
	// 5. di service , akan mencari dengan bantuan repository user dengan email
	// 6. cek validasi password benar atau salah

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {

		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Login Account failed", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Login Account failed", http.StatusUnprocessableEntity, "error", errorMessage)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	token, err := h.authService.GenerateToken(loggedInUser.ID)
	if err != nil {
		response := helper.ApiResponse("Login  failed", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatOneUser(loggedInUser, token)

	response := helper.ApiResponse("Login successfull", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}

func (h *userHandler) GetUserByID(c *gin.Context) {
	fmt.Println("masokkkkkkk")
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	userDetail, err := h.userService.GetUserByID(userID)

	if userDetail.ID == 0 {

		errorMessage := gin.H{"errors": "Failed to get user id"}

		response := helper.ApiResponse("Failed to get user id", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err != nil {
		response := helper.ApiResponse("Error to get detail user", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return

	}

	response := helper.ApiResponse("user detail", http.StatusOK, "success", user.FormatOneUser(userDetail, ""))
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) GetAllUser(c *gin.Context) {

	// currentUser := c.MustGet("currentUser").(user.User)
	// userID := currentUser.ID

	allUser, err := h.userService.GetAllUser()

	if err != nil {

		response := helper.ApiResponse("failed to get user's transactions", http.StatusBadRequest, "error", err)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("User's transactions", http.StatusOK, "success", user.FormatUsers(allUser))

	c.JSON(http.StatusBadRequest, response)

}

func (h *userHandler) UploadProfile(c *gin.Context) {
	fmt.Println("masokkkkkkkkk  000000")
	file, err := c.FormFile("foto")

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("failed to upload Foto image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//  next pakat jwt bukan
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID
	currentTime := time.Now()

	// path := "images/" + + currentTime.Format("2006#01#02") + "#" + file.Filename
	fmt.Println("masokkkkkkkkk  111")
	path := fmt.Sprintf("profile/%d-%s-%s", userID, currentTime.Format("2006-01-02-3:4:5"), file.Filename)
	_, err = h.userService.SaveProfile(userID, path)
	err = c.SaveUploadedFile(file, path)
	fmt.Println("masokkkkkkkkk  222")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("failed to upload Foto image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveProfile(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("failed to upload Foto image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.ApiResponse("Foto successfully uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}
