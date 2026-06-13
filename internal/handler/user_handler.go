package handlers

import (
	"context"
	"strconv"
	"time"

	"user-api/database"
	"user-api/db/sqlc"
	"user-api/internal/logger"
	"user-api/internal/models"
	"user-api/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func CreateUser(c *fiber.Ctx) error {

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := utils.Validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	dob, err := time.Parse("2006-01-02", user.DOB)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid date format. Use YYYY-MM-DD",
		})
	}

	createdUser, err := database.Queries.CreateUser(
		context.Background(),
		sqlc.CreateUserParams{
			Name: user.Name,
			Dob:  dob,
		},
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	logger.Log.Info(
		"User created",
		zap.Int32("id", createdUser.ID),
		zap.String("name", createdUser.Name),
	)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User Created Successfully",
		"id":      createdUser.ID,
	})

}

func GetUserByID(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	dbUser, err := database.Queries.GetUserByID(
		context.Background(),
		int32(id),
	)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User Not Found",
		})
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
	return c.JSON(user)

}

func GetAllUsers(c *fiber.Ctx) error {

	page := c.Query("page", "1")
	limit := c.Query("limit", "5")

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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}

	start := (pageNum - 1) * limitNum

	if start >= len(allUsers) {
		return c.JSON([]models.User{})
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

	return c.JSON(users)

}

func UpdateUser(c *fiber.Ctx) error {

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := utils.Validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	dob, err := time.Parse("2006-01-02", user.DOB)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid date format. Use YYYY-MM-DD",
		})
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}

	logger.Log.Info(
		"User updated",
		zap.Int("id", id),
		zap.String("name", user.Name),
	)

	return c.JSON(fiber.Map{
		"message": "User Updated Successfully",
	})

}

func DeleteUser(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	err = database.Queries.DeleteUser(
		context.Background(),
		int32(id),
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete user",
		})
	}

	logger.Log.Info(
		"User deleted",
		zap.Int("id", id),
	)

	return c.SendStatus(fiber.StatusNoContent)

}
