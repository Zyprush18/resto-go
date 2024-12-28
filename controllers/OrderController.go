package controllers

import (
	"fmt"

	"github.com/Zyprush18/resto-go/model/entity"
	"github.com/Zyprush18/resto-go/model/request"
	"github.com/Zyprush18/resto-go/model/response"
	"github.com/Zyprush18/resto-go/repositories/databases"
	"github.com/gofiber/fiber/v2"
)

func OrderControllerIndex(c *fiber.Ctx) error {
	var order []entity.Order

	if err := databases.DB.Preload("User").Find(&order).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error",
			"error":   err.Error(),
		})
	}

	if len(order) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "data belum ada",
		})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "success",
			"data":    order,
		})
	}

}

func OrderControllerCreate(c *fiber.Ctx) error {
	inputOrder := new(request.Order)

	if err := c.BodyParser(&inputOrder); err != nil {
		return err
	}

	order := &response.Order{
		TotalPrice: inputOrder.TotalPrice,
		Status:     inputOrder.Status,
		UserId:     inputOrder.UserId,
	}

	if err := databases.DB.Create(&order).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed create order",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success create order",
		"data":    order,
	})
}

func OrderControllerShow(c *fiber.Ctx) error {
	var Order entity.Order
	id := c.Params("id")

	if err := databases.DB.Preload("User").Find(&Order, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "not found data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    Order,
	})
}

func OrderControllerUpdated(c *fiber.Ctx) error {
	orderInput := new(request.Order)
	id := c.Params("id")

	if err := c.BodyParser(orderInput); err != nil {
		return err
	}

	var order response.Order

	if err := databases.DB.First(&order, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "not found data",
		})
	}

	fmt.Println(order.Status)

	if orderInput.TotalPrice != 0 {
		order.TotalPrice = orderInput.TotalPrice
	}

	if orderInput.Status != "" {
		order.Status = orderInput.Status
	}
	if orderInput.UserId != 0 {
		order.UserId = orderInput.UserId
	}

	if err := databases.DB.Save(&order).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed update",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success update data",
		"data":    order,
	})
}

func OrderControllerDelete(c *fiber.Ctx) error {
	var Order entity.Order
	id := c.Params("id")

	if err := databases.DB.First(&Order, "id = ? ", id).Delete(&Order).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "failed delete",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success delete",
	})
}
