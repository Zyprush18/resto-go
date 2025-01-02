package controllers

import (
	"fmt"

	"github.com/Zyprush18/resto-go/model/entity"
	"github.com/Zyprush18/resto-go/model/request"
	"github.com/Zyprush18/resto-go/model/response"
	"github.com/Zyprush18/resto-go/repositories/databases"
	"github.com/Zyprush18/resto-go/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ReservationControllerIndex(c *fiber.Ctx) error {
	var Reservation []entity.Reservation

	if err := databases.DB.Preload("User").Find(&Reservation).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed get data",
			"error": err.Error(),
		})
	}

	if len(Reservation) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "data belum ada",
		})
	}else{
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "success",
			"data": Reservation,
		})
	}

}

func ReservationControllerCreate(c *fiber.Ctx) error {
	inputReservation := new(request.Reservation)

	if err := c.BodyParser(inputReservation); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(inputReservation); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "failed validation",
			"error": err.Error(),
		})
	}

	parsedTime, err := service.ParsedTime(inputReservation.Time)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse time",
			"error": err.Error(),
		})
	}
	format := parsedTime.Format("15:04:05")

	parsedDate, err := service.ParsedDate(inputReservation.Date)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse date",
			"error": err.Error(),
		})
	}

	formatDate := parsedDate.Format("2006-01-02")

	fmt.Println("formatDate", formatDate)
	fmt.Println("formatTime", format)

	Reservation := &response.Reservation{
		Date: formatDate,
		Time: format,
		GuestCount: inputReservation.GuestCount,
		Status: inputReservation.Status,
		UserId: inputReservation.UserId,
	}

	if err := databases.DB.Create(&Reservation).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed Create Data",
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data": Reservation,
	})
	
}

func ReservationControllerShow(c *fiber.Ctx) error  {
	Reservation := new(entity.Reservation)
	id := c.Params("id")

	if err := databases.DB.Preload("User").First(&Reservation, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error",
			"error": err.Error(),
		})
	} 

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data": Reservation,
	})
}

func ReservationControllerUpdate(c *fiber.Ctx) error  {
	inputReservation := new(request.Reservation)
	id := c.Params("id")

	if err := c.BodyParser(inputReservation); err != nil {
		return err
	}

	var Reservation response.Reservation

	if err := databases.DB.First(&Reservation, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Not Found",
		})
	}

	if inputReservation.Date != "" {
		Reservation.Date = inputReservation.Date
	}

	if inputReservation.Time != "" {
		Reservation.Time = inputReservation.Time
	}

	if inputReservation.GuestCount != 0 {
		Reservation.GuestCount = inputReservation.GuestCount
	}

	if inputReservation.Status != "" {
		Reservation.Status = inputReservation.Status
	}

	if inputReservation.UserId != 0 {
		Reservation.UserId = inputReservation.UserId
	}

	if err := databases.DB.Save(&Reservation).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed Update",
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Update",
		"data": Reservation,
	})

}

func ReservationControllerDelete(c *fiber.Ctx) error {
	reservation := new(entity.Reservation)
	id := c.Params("id")

	if err := databases.DB.First(&reservation, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Not Found",
		})
	}

	if err := databases.DB.Delete(&reservation, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed Delete",
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Delete",
	})
}