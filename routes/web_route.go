package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/maonks/absen-rfid-backend/controllers"
	"gorm.io/gorm"
)

func WebRoutes(app *fiber.App, db *gorm.DB) {

	app.Get("/websocket", websocket.New(controllers.WebsocketHandler))

	app.Get("/", controllers.Dashboard(db))
	app.Get("/absensi", controllers.AbsensiPage(db))
	app.Get("/api/absen/table", controllers.SearchAbsen(db))
	app.Get("/api/kartu/:uid/edit", controllers.EditModal(db))

	app.Get("/riwayat", controllers.RiwayatPage(db))

}
