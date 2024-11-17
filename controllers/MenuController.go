package controllers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Zyprush18/resto/model/entity"
	"github.com/Zyprush18/resto/model/request"
	"github.com/Zyprush18/resto/repositories/databases"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func MenuControllerIndex(c *fiber.Ctx) error {
	var menu []entity.Menu

	if err := databases.DB.Find(&menu).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed Show User",
		})
	}
	if len(menu) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "data belum ada",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data": menu,
	})
}

func MenuControllerCreate(c *fiber.Ctx) error {
	menu := new(request.Menu)

	if err := c.BodyParser(menu); err != nil {
		return err
	}

	validate := validator.New()

	if err:= validate.Struct(menu);err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "failed validation",
			"error": err.Error(),
		})
	}


	var filename string

	file, err := c.FormFile("image")

	if err != nil {
		log.Println(err)
	}

	if file != nil {
		filename = file.Filename

		err := c.SaveFile(file, fmt.Sprintf("./public/img/%s", filename))
		if err != nil {
			return err
		}
	}else{
		log.Println("failed upload file")
	}


	info := c.FormValue("is_available")
	boolInfo,err := strconv.ParseBool(info)
	if err != nil {
		return err
	}
	
	newMenu := &entity.Menu{
		Name: menu.Name,
		Price: menu.Price,
		IsAvailable: &boolInfo,
		Image: filename,
	}


	if err := databases.DB.Create(&newMenu).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed create menu",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success create",
		"data": newMenu,
	})

}

func MenuControllerShow(c *fiber.Ctx) error {
	menu := new(entity.Menu)
	id := c.Params("id")

	if err := databases.DB.First(&menu, "id = ?",id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Not found",
		})
	}


	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"success",
		"data": menu,
	})
}