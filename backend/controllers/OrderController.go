package controllers

import (
	"fmt"
	"log"
	"strings"

	"github.com/Zyprush18/resto-go/backend/model/entity"
	"github.com/Zyprush18/resto-go/backend/model/request"
	"github.com/Zyprush18/resto-go/backend/model/response"
	"github.com/Zyprush18/resto-go/backend/repositories/databases"
	"github.com/Zyprush18/resto-go/backend/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func OrderControllerIndex(c *fiber.Ctx) error {
	var order []entity.Order
	// ambil user id dan role dari request context
	role := c.Locals(service.RoleKey).(string)
	user_id := c.Locals(service.UserIdKey).(int)
	fmt.Println(user_id)
	
	query := databases.DB.Preload("User").Preload("OrderItem")
	if strings.ToLower(role)  != "admin" {
		log.Println("jalan lagii")
		query = query.Where("user_id = ?", user_id)
	}

	fmt.Println(query)
	if err := query.Find(&order).Error; err != nil {
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

	user_id := c.Locals(service.UserIdKey).(int)

	if err := c.BodyParser(&inputOrder); err != nil {
		return err
	}

	validate := validator.New()

	if err := validate.Struct(inputOrder); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "failed validation",
			"error":   err.Error(),
		})
	}


	order := &response.Order{
		TotalPrice: inputOrder.TotalPrice,
		Status:     inputOrder.Status,
		UserId:     inputOrder.UserId | uint(user_id),
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

	role := c.Locals(service.RoleKey).(string)
	user_id := c.Locals(service.UserIdKey).(int)

	query := databases.DB.Preload("User").Preload("OrderItem")
	if strings.ToLower(role) != "admin" {
		query = query.Where("user_id =?", user_id)
	}

	if err := query.First(&Order, "id = ?", id).Error; err != nil {
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
	user_id := c.Locals(service.UserIdKey).(int)


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
		order.UserId = orderInput.UserId | uint(user_id)
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

	query := databases.DB 

	role := c.Locals(service.RoleKey).(string)
	user_id := c.Locals(service.UserIdKey).(int)

	if strings.ToLower(role) != "admin" {
		query = query.Where("user_id =?", user_id)
	}

	if err := query.First(&Order, "id = ? ", id).Delete(&Order).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "failed delete",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success delete",
	})
}
