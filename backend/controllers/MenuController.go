package controllers

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Zyprush18/resto-go/backend/model/entity"
	"github.com/Zyprush18/resto-go/backend/model/request"
	"github.com/Zyprush18/resto-go/backend/repositories/databases"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func MenuControllerIndex(c *fiber.Ctx) error {
	var menu []entity.Menu

	if err := databases.DB.Preload("OrderItem").Find(&menu).Error; err != nil {
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
		"data":    menu,
	})
}

func MenuControllerCreate(c *fiber.Ctx) error {
	menu := new(request.Menu)

	if err := c.BodyParser(menu); err != nil {
		return err
	}

	validate := validator.New()

	if err := validate.Struct(menu); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "failed validation",
			"error":   err.Error(),
		})
	}

	var filename string
	var randomfilename string

	file, err := c.FormFile("image")

	if err != nil {
		log.Println(err)
	}

	if file != nil {
		filename = file.Filename

		// fmt.Println(path.Ext(file.Filename))
		extension := path.Ext(filename)
		rand.Seed(time.Now().UnixNano())
		length := 25
		const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		result := make([]byte, length)
		for i := range result {
			result[i] = charset[rand.Intn(len(charset))]
		}
		randomfilename = string(result)

		err := c.SaveFile(file, fmt.Sprintf("./public/storage/img/%s%s", randomfilename, extension))
		if err != nil {
			return err
		}
	} else {
		log.Println("failed upload file")
	}

	info := c.FormValue("is_available")
	boolInfo, err := strconv.ParseBool(info)
	if err != nil {
		return err
	}

	newMenu := &entity.Menu{
		Name:        menu.Name,
		Price:       menu.Price,
		IsAvailable: &boolInfo,
		Image:       randomfilename,
	}

	if err := databases.DB.Create(&newMenu).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed create menu",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success create",
		"data":    newMenu,
	})

}

func MenuControllerShow(c *fiber.Ctx) error {
	menu := new(entity.Menu)
	id := c.Params("id")

	if err := databases.DB.Preload("OrderItem").First(&menu, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    menu,
	})
}

func MenuControllerUpdate(c *fiber.Ctx) error {
	inputMenu := new(request.Menu)
	id := c.Params("id")
	if err := c.BodyParser(inputMenu); err != nil {
		return err
	}

	var Menu entity.Menu

	if err := databases.DB.First(&Menu, "id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Not found",
		})
	}

	if inputMenu.Name != "" {
		Menu.Name = inputMenu.Name
	}
	if inputMenu.Price != 0 {
		Menu.Price = inputMenu.Price
	}

	info := c.FormValue("is_available")
	if info != "" {
		bools, err := strconv.ParseBool(info)
		if err != nil {
			return err
		}
		Menu.IsAvailable = &bools
	}

	file, _ := c.FormFile("image")
	// if err != nil {
	// 	return err
	// }

	if file != nil {

		if Menu.Image != "" {
			namefie := Menu.Image
			pathdir := "./public/storage/img"

			files, err := filepath.Glob(filepath.Join(pathdir, namefie+".*"))

			if err != nil {
				return err
			}

			var coba string
			for _, f := range files {
				coba = f
			}

			if err := os.Remove(coba); err != nil {
				return err
			}
		}

		filename := file.Filename
		extension := path.Ext(filename)

		rand.Seed(time.Now().UnixNano())
		length := 25
		const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		result := make([]byte, length)
		for i := range result {
			result[i] = charset[rand.Intn(len(charset))]
		}
		randomfilename := string(result)

		err := c.SaveFile(file, fmt.Sprintf("./public/storage/img/%s%s", randomfilename, extension))
		if err != nil {
			return err
		}

		Menu.Image = randomfilename
	}

	if err := databases.DB.Save(&Menu).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed updated menu",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success update menu",
		"data":    Menu,
	})
}

func MenuControllerDelete(c *fiber.Ctx) error {
	menu := new(entity.Menu)
	id := c.Params("id")

	if err := databases.DB.First(&menu, "id = ?", id).Delete(&menu).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed delete menu",
		})
	}

	namefie := menu.Image
	pathdir := "./public/storage/img"

	files, err := filepath.Glob(filepath.Join(pathdir, namefie+".*"))

	if err != nil {
		return err
	}

	var coba string
	for _, f := range files {
		coba = f
	}

	if err := os.Remove(coba); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success delete menu",
	})
}
