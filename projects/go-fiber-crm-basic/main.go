package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/notfreecap/go-fiber-crm-basic/database"
	"github.com/notfreecap/go-fiber-crm-basic/lead"
)

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(8080)
	defer database.DBConn.Close()
}

func setupRoutes(app *fiber.App) {
	basePath := "/api/v1/lead"
	app.Get(basePath, lead.GetLeads)
	app.Get(basePath+"/:id", lead.GetLead)
	app.Post(basePath, lead.NewLead)
	app.Delete(basePath+"/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("failed to connect database")
	}
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Connection opened to database")
	fmt.Println("Database Migrated")
}
