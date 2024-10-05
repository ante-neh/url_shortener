package routes

import (
	"github.com/ante-neh/url_shortener/types"
	"github.com/ante-neh/url_shortener/util"
	"github.com/gofiber/fiber/v2"
)

func ShortenUrl(c *fiber.Ctx) error {
	body := new(types.Request)

	err := c.BodyParser(&body) 

	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Unable to parse request"})
	}

	if !govalidator.IsUrl(body.URL){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid url"})
	}

	if !util.RemoveDomainError(body.URL){
		return c.Status(fiber.StatusServiceUnavailable).JSON((fiber.Map{"error":""}))
	}

	body.URL = util.EnforeHTTP(body.URL)

	return nil 
}