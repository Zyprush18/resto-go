package controllers

import (
	"github.com/Zyprush18/resto-go/model/entity"
	"github.com/Zyprush18/resto-go/model/request"
	"github.com/Zyprush18/resto-go/model/response"
	"github.com/Zyprush18/resto-go/repositories/databases"
	"github.com/Zyprush18/resto-go/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserControllerIndex(c *fiber.Ctx) error {
	// User := new(entity.User)
	var User []entity.User

	if err := databases.DB.Preload("Order").Find(&User).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed Show User",
		})
	}

	if len(User) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "data belum ada",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    User,
	})
}

func UserControllerCreate(c *fiber.Ctx) error {
	User := new(request.User)

	if err := c.BodyParser(User); err != nil {
		return err
	}

	validate := validator.New()

	if err := validate.Struct(User); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "failed validation",
			"error":   err.Error(),
		})
	}

	hash, err := service.HashingPas(User.Password)

	if err != nil {
		return err
	}

	CreateUser := &response.User{
		Name:     User.Name,
		Email:    User.Email,
		Phone:    User.Phone,
		Password: hash,
	}

	if err := databases.DB.Create(&CreateUser).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed Create User",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success create",
		"data":    CreateUser,
	})

}

func UserControllerShow(c *fiber.Ctx) error {
	User := new(entity.User)

	id := c.Params("id")

	if err := databases.DB.Preload("Order").First(&User, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Not Found ",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    User,
	})
}

func UserControllerUpdate(c *fiber.Ctx) error {
	User := new(request.User)
	id := c.Params("id")

	if err := c.BodyParser(&User); err != nil {
		return err
	}

	update := &entity.User{}

	if err := databases.DB.First(&update, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Not Found ",
		})
	}

	if User.Name != "" {
		update.Name = User.Name
	}

	if User.Email != "" {
		update.Email = User.Email
	}
	if User.Phone != "" {
		update.Phone = User.Phone
	}

	if User.Password != "" {
		hash, err := service.HashingPas(User.Password)

		if err != nil {
			return err
		}

		update.Password = hash
	}

	if err := databases.DB.Save(&update).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed Update User",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    update,
	})
}

func UserControllerDelete(c *fiber.Ctx) error {
	user := new(entity.User)

	id := c.Params("id")

	if err := databases.DB.First(&user, "id = ?", id).Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed delete user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success delete user",
	})
}
