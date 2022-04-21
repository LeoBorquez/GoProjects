package main

import (
	"fmt"

	"github.com/LeoBorquez/GoProjects/fiber-crm/database"
	"github.com/LeoBorquez/GoProjects/fiber-crm/lead"
	"github.com/gofiber/fiber"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)          // get all leads
	app.Get("/api/v1/lead/:id", lead.GetLead)       // get one lead by id
	app.Post("/api/v1/lead", lead.NewLead)          // create a new lead
	app.Delete("/api/v1/lead/:id", lead.DeleteLead) // delete lead by id
}

func initDatabase() {
	var err error
	dsn := "host=localhost user=gborquez password= dbname=lead port=5432 sslmode=disable"
	database.DBconn, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connection opened to database")
	database.DBconn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)

}
