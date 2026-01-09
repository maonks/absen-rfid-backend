package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/maonks/absen-rfid-backend/controllers"
	"gorm.io/gorm"
)

func WebRoutes(app *fiber.App, db *gorm.DB) {

	app.Get("/", controllers.Dashboard(db))
	app.Get("/websocket", websocket.New(controllers.WebsocketHandler))

	app.Get("/perhitungan", controllers.PerhitunganPage(db))

}
