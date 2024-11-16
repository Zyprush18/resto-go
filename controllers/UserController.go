package controllers

import (
	"github.com/Zyprush18/resto/model/entity"
	"github.com/Zyprush18/resto/model/request"
	"github.com/Zyprush18/resto/repositories/databases"
	"github.com/Zyprush18/resto/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserControllerIndex(c *fiber.Ctx) error  {
	// User := new(entity.User)
	var User []entity.User

	if err := databases.DB.Find(&User).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed Show User",
		})
	}

	if len(User) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":"data belum ada",
		})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"success",
		"data": User,
	})
}


func UserControllerCreate(c *fiber.Ctx) error {
	User := new(request.User)

	if err := c.BodyParser(User); err != nil {
		return err
	}


	validate := validator.New()

	if err:= validate.Struct(User);err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "failed validation",
			"error": err.Error(),
		})
	}

	hash, err := service.HashingPas(User.Password)

	if err != nil {
		return err	
	}

	CreateUser := &entity.User{
		Name: User.Name,
		Email: User.Email,
		Phone: User.Phone,
		Password: hash,
	}

	if err:= databases.DB.Create(&CreateUser).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed Create User",
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success create",
		"data": CreateUser,
	})


}