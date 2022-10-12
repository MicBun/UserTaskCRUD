package controllers

import (
	"UserSimpleCRUD/models"
	"UserSimpleCRUD/utils/token"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Task interface {
	GetTask(ctx *gin.Context)
	PostTask(ctx *gin.Context)
	PutTask(ctx *gin.Context)
	PatchTask(ctx *gin.Context)
	DeleteTask(ctx *gin.Context)
}

// GetTask godoc
// @Summary Get GetTask.
// @Description Get all Task by login.
// @Tags Task
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Task
// @Router /task/GetTask [get]
func GetTask(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	decoded, _ := token.ExtractTokenID(ctx)
	var tasks []models.Task
	if err := db.Where("user_id = ?", decoded).Find(&tasks).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Tasks not found!"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

type PostTaskInput struct {
	TaskName    string `json:"taskName"`
	Description string `json:"description"`
}

// PostTask godoc
// @Summary Create New Task.
// @Description Creating a new Task.
// @Tags Task
// @Param Body body PostTaskInput true "the body to create a new Task"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Task
// @Router /task/PostTask [post]
func PostTask(ctx *gin.Context) {
	var input PostTaskInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	decoded, _ := token.ExtractTokenID(ctx)

	task := models.Task{
		UserID:      decoded,
		TaskName:    input.TaskName,
		Description: input.Description,
		Status:      false,
	}

	ctx.MustGet("db").(*gorm.DB).Create(&task)
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": task})
}

type PutTaskInput struct {
	TaskID      uint   `json:"taskID"`
	TaskName    string `json:"taskName"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

// PutTask godoc
// @Summary Mark Task as Done.
// @Description Change Task By ID.
// @Tags Task
// @Produce json
// @Param Body body PutTaskInput true "the body to update task"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200
// @Router /task/PutTask [put]
func PutTask(ctx *gin.Context) {
	var input PutTaskInput
	decoded, _ := token.ExtractTokenID(ctx)
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := ctx.MustGet("db").(*gorm.DB)
	var task models.Task
	if err := db.Where("id = ? AND user_id = ?", input.TaskID, decoded).First(&task).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Model(&task).Updates(models.Task{
		TaskName:    input.TaskName,
		Description: input.Description,
		Status:      input.Status,
	})

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "updated data": task})
}

type PatchTaskInput struct {
	TaskID uint `json:"taskID"`
}

// PatchTask godoc
// @Summary Set Task as Done.
// @Description Set Task as Done by ID.
// @Tags Task
// @Produce json
// @Param Body body PatchTaskInput true "the body to patch task"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200
// @Router /task/PatchTask [patch]
func PatchTask(ctx *gin.Context) {
	var input PatchTaskInput
	decoded, _ := token.ExtractTokenID(ctx)
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := ctx.MustGet("db").(*gorm.DB)
	var task models.Task
	if err := db.Where("id = ? AND user_id = ?", input.TaskID, decoded).First(&task).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Model(&task).Updates(models.Task{Status: true})

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "updated data": task})
}

type DeleteTaskInput struct {
	TaskID uint `json:"taskID"`
}

// DeleteTask godoc
// @Summary Delete one Item.
// @Description Delete a task by id.
// @Tags Task
// @Produce json
// @Param Body body DeleteTaskInput true "the body to delete task"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /task/DeleteTask [delete]
func DeleteTask(ctx *gin.Context) {
	var input DeleteTaskInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := ctx.MustGet("db").(*gorm.DB)
	var datum models.Task
	decoded, _ := token.ExtractTokenID(ctx)

	if err := db.Where("id = ? AND user_id = ?", input.TaskID, decoded).First(&datum).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&datum)

	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
