package middleware

import (
	"log"
	"strconv"
	"strings"

	"github.com/Zyprush18/resto-go/backend/service"
	"github.com/gofiber/fiber/v2"
)

func MiddlewareGlobal(c *fiber.Ctx) error {
	// ambil token authorization di header
	token := strings.Split(c.Get("Authorization")," ")
	if token == nil || token[0] != "Bearer" || token[1] == "" {
		log.Println("Token is Missing")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message":"Unauthorized",
		})
	}

	// parsing token jwt
	parsToken,err := service.ParsedToken(token[1])
	if err != nil {
		log.Println("Failed Parsing Token: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	id_user, err := strconv.Atoi(parsToken.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something Went Wrong",
		})
	}

	// simpan role ke dalam context
	c.Locals(service.RoleKey, parsToken.Role)
	c.Locals(service.UserIdKey, id_user)


	return c.Next()
}

func MiddlewareAccess(roles ...string) fiber.Handler  {
	return func (c *fiber.Ctx) error  {
		role := c.Locals(service.RoleKey)
		for _, v := range roles {
			if role == v {
				return  c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message":"Forbidden Access",
		})
	}
}