package controllers

import (
	"github.com/Zyprush18/resto-go/model/entity"
	"github.com/Zyprush18/resto-go/model/request"
	"github.com/Zyprush18/resto-go/model/response"
	"github.com/Zyprush18/resto-go/repositories/databases"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func OrderItemControllerIndex(c *fiber.Ctx) error {
	var orderitem []entity.OrderItem

	if err := databases.DB.Preload("Order").Preload("Menu").Find(&orderitem).Error; err != nil {
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

	validate := validator.New()

	if err := validate.Struct(orderItemInput); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Failed Validation",
			"error": err.Error(),
		})
	}


	orderItemResponse := &response.OrderItem{
		Quantity: orderItemInput.Quantity,
		Price: orderItemInput.Price,
		OrderId: orderItemInput.OrderId,
		MenuId: orderItemInput.MenuId,
	}

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

func OrderItemControllerShow(c *fiber.Ctx) error  {
	orderItem := new(entity.OrderItem)
	id := c.Params("id")

	if err := databases.DB.Preload("Order").Preload("Menu").First(&orderItem, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error",
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data": orderItem,
	})
}

func OrderItemControllerUpdate(c *fiber.Ctx) error {
	inputOrderItem := new(request.OrderItem)
	id := c.Params("id")

	if err := c.BodyParser(inputOrderItem); err != nil {
		return err
	}

	var orderItem response.OrderItem

	if err := databases.DB.First(&orderItem, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Not Found",
		})
	}


	if inputOrderItem.Quantity != 0 {
		orderItem.Quantity = inputOrderItem.Quantity
	}
	if inputOrderItem.Price != 0 {
		orderItem.Price = inputOrderItem.Price
	}
	if inputOrderItem.OrderId != 0 {
		orderItem.OrderId = inputOrderItem.OrderId
	}
	if inputOrderItem.MenuId != 0 {
		orderItem.MenuId = inputOrderItem.MenuId
	}

	if err := databases.DB.Save(&orderItem).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed Update",
			"error": err.Error(),
		})
	}

	return c. Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Update",
		"data": orderItem,
	})

}

func OrderItemControllerDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	var orderItem response.OrderItem

	if err := databases.DB.First(&orderItem, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Not Found",
		})
	}

	if err := databases.DB.Delete(&orderItem, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed Delete",
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Delete",
	})
}