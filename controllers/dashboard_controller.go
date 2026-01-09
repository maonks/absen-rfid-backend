package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Dashboard(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("OPEN DASHBOARD")

		return c.Render("fragments/home_page", fiber.Map{
			"Title": "Dashboard",
		}, "layouts/main")

	}
}
