package controllers

import (
	"github.com/Zyprush18/resto-go/model/entity"
	"github.com/Zyprush18/resto-go/model/request"
	"github.com/Zyprush18/resto-go/model/response"
	"github.com/Zyprush18/resto-go/repositories/databases"
	"github.com/gofiber/fiber/v2"
)

func OrderItemControllerIndex(c *fiber.Ctx) error {
	var orderitem []entity.OrderItem

	if err := databases.DB.Preload("Order", "Menu").Find(&orderitem).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error",
			"error":   err.Error(),
		})
	}

	if len(orderitem) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "data belum ada",
		})
	}else{
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "success",
			"data": orderitem,
		})
	}
}

func OrderItemControllerCreate(c *fiber.Ctx) error  {
	orderItemInput := new(request.OrderItem)
	
	if err := c.BodyParser(&orderItemInput); err != nil {
		return err
	}

	orderItemResponse := &response.OrderItem{}

	if err := databases.DB.Create(orderItemResponse).Error; err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed Create data",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success Created",
		"data": orderItemResponse,
	})

}