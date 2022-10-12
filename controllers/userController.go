package controllers

import (
	"UserSimpleCRUD/models"
	"UserSimpleCRUD/utils/token"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type UserInterface interface {
	LoginUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	GetAllUser(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	UpdateUserByID(ctx *gin.Context)
	DeleteUserByID(ctx *gin.Context)
	GetUserByLogin(ctx *gin.Context)
}

type LoginUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginUser godoc
// @Summary Login with credential.
// @Description Logging in to get jwt token to access admin or user api by roles.
// @Tags Auth
// @Param Body body LoginUserInput true "the body to login a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func LoginUser(ctx *gin.Context) {
	var input LoginUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginCheck, err := models.UserLogin(models.User{
		Username: input.Username,
		Password: input.Password,
	}).LoginCheck(ctx.MustGet("db").(*gorm.DB))

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "login success", "token": loginCheck})
}

type CreateUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Salary   int    `json:"salary"`
	Address  string `json:"address"`
}

// CreateUser godoc
// @Summary Create New CreateUser.
// @Description Creating a new User.
// @Tags Admin
// @Param Body body CreateUserInput true "the body to create a new User"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.User
// @Router /admin/createUser [post]
func CreateUser(ctx *gin.Context) {
	isAdmin, _ := token.ExtractTokenIsAdmin(ctx)
	if !isAdmin {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Admin Only!"})
		return
	}

	var input CreateUserInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: input.Username,
		Password: input.Password,
		Name:     input.Name,
		Salary:   input.Salary,
		Address:  input.Address,
	}
	ctx.MustGet("db").(*gorm.DB).Create(&user)
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}

// GetAllUser godoc
// @Summary Get all Users.
// @Description Get a list of Users.
// @Tags Admin
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} []models.User
// @Router /admin/getAllUser [get]
func GetAllUser(ctx *gin.Context) {
	isAdmin, _ := token.ExtractTokenIsAdmin(ctx)
	if !isAdmin {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Admin Only!"})
		return
	}
	db := ctx.MustGet("db").(*gorm.DB)
	var users []models.User
	db.Find(&users)

	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

type IDInput struct {
	ID uint `json:"id"`
}

// GetUserByID godoc
// @Summary Get GetUserByID.
// @Description Get a User by id.
// @Tags Admin
// @Param Body body IDInput true "the body to get a User"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.User
// @Router /admin/getUserByID [post]
func GetUserByID(ctx *gin.Context) {
	isAdmin, _ := token.ExtractTokenIsAdmin(ctx)
	if !isAdmin {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Admin Only!"})
		return
	}
	db := ctx.MustGet("db").(*gorm.DB)
	var input IDInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := db.Where("id = ?", input.ID).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

type UpdateUserByIDInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Salary   int    `json:"salary"`
	Address  string `json:"address"`
}

// UpdateUserByID godoc
// @Summary Update User.
// @Description Update User By ID.
// @Tags Admin
// @Produce json
// @Param id path string true "user id"
// @Param Body body UpdateUserByIDInput true "the body to update User"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200
// @Router /admin/updateUserByID/{id} [patch]
func UpdateUserByID(ctx *gin.Context) {
	isAdmin, _ := token.ExtractTokenIsAdmin(ctx)
	if !isAdmin {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Admin Only!"})
		return
	}
	var input UpdateUserByIDInput
	//decoded, _ := token.ExtractTokenID(ctx)
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := ctx.MustGet("db").(*gorm.DB)
	var user models.User
	if err := db.Where("id = ?", ctx.Param("id")).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Model(&user).Updates(models.User{
		Username: input.Username,
		Password: input.Password,
		Name:     input.Name,
		Salary:   input.Salary,
		Address:  input.Address,
	})
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}

// DeleteUserByID godoc
// @Summary Delete one User.
// @Description Delete a User by id.
// @Tags Admin
// @Produce json
// @Param Body body IDInput true "the body to delete a User"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /admin/deleteUserByID [delete]
func DeleteUserByID(ctx *gin.Context) {
	isAdmin, _ := token.ExtractTokenIsAdmin(ctx)
	if !isAdmin {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Admin Only!"})
		return
	}
	var input IDInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := ctx.MustGet("db").(*gorm.DB)
	var user models.User
	if err := db.Where("id = ?", input.ID).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	db.Delete(&user)

	ctx.JSON(http.StatusOK, gin.H{"message": "Delete Success"})
}

// GetUserByLogin godoc
// @Summary Get User Profile.
// @Description Get a Profile by login auth.
// @Tags User
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.User
// @Router /user/profile [get]
func GetUserByLogin(ctx *gin.Context) {
	isAdmin, _ := token.ExtractTokenIsAdmin(ctx)
	if isAdmin {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "User Only!"})
		return
	}
	db := ctx.MustGet("db").(*gorm.DB)
	var user models.User
	decoded, _ := token.ExtractTokenID(ctx)
	if err := db.Where("id = ?", decoded).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}
