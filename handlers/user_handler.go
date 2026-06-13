package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"user-api/database"
	"user-api/db/sqlc"
	"user-api/logger"
	"user-api/models"
	"user-api/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := utils.Validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	dob, err := time.Parse("2006-01-02", user.DOB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid date format. Use YYYY-MM-DD",
		})
		return
	}

	createdUser, err := database.Queries.CreateUser(
		context.Background(),
		sqlc.CreateUserParams{
			Name: user.Name,
			Dob:  dob,
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	logger.Log.Info(
		"User created",
		zap.Int32("id", createdUser.ID),
		zap.String("name", createdUser.Name),
	)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User Created Successfully",
		"id":      createdUser.ID,
	})

}

func GetUserByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	dbUser, err := database.Queries.GetUserByID(
		context.Background(),
		int32(id),
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User Not Found",
		})
		return
	}

	user := models.User{
		ID:   int(dbUser.ID),
		Name: dbUser.Name,
		DOB:  dbUser.Dob.Format("2006-01-02"),
		Age:  utils.CalculateAge(dbUser.Dob.Format("2006-01-02")),
	}

	if dbUser.CreatedAt.Valid {
		user.CreatedAt = dbUser.CreatedAt.Time
	}

	c.JSON(http.StatusOK, user)

}

func GetAllUsers(c *gin.Context) {

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "5")

	pageNum, _ := strconv.Atoi(page)
	limitNum, _ := strconv.Atoi(limit)

	if pageNum < 1 {
		pageNum = 1
	}

	if limitNum < 1 {
		limitNum = 5
	}

	allUsers, err := database.Queries.GetAllUsers(
		context.Background(),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch users",
		})
		return
	}

	start := (pageNum - 1) * limitNum

	if start >= len(allUsers) {
		c.JSON(http.StatusOK, []models.User{})
		return
	}

	end := start + limitNum
	if end > len(allUsers) {
		end = len(allUsers)
	}

	users := []models.User{}

	for _, dbUser := range allUsers[start:end] {

		user := models.User{
			ID:   int(dbUser.ID),
			Name: dbUser.Name,
			DOB:  dbUser.Dob.Format("2006-01-02"),
			Age:  utils.CalculateAge(dbUser.Dob.Format("2006-01-02")),
		}

		if dbUser.CreatedAt.Valid {
			user.CreatedAt = dbUser.CreatedAt.Time
		}

		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)

}

func UpdateUser(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := utils.Validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	dob, err := time.Parse("2006-01-02", user.DOB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid date format. Use YYYY-MM-DD",
		})
		return
	}

	err = database.Queries.UpdateUser(
		context.Background(),
		sqlc.UpdateUserParams{
			Name: user.Name,
			Dob:  dob,
			ID:   int32(id),
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user",
		})
		return
	}

	logger.Log.Info(
		"User updated",
		zap.Int("id", id),
		zap.String("name", user.Name),
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "User Updated Successfully",
	})

}

func DeleteUser(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	err = database.Queries.DeleteUser(
		context.Background(),
		int32(id),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete user",
		})
		return
	}

	logger.Log.Info(
		"User deleted",
		zap.Int("id", id),
	)

	c.Status(http.StatusNoContent)

}
