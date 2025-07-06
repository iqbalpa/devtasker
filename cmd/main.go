package main

import (
	"devtasker/internal"
	"devtasker/internal/utils"
	"fmt"
	"os"
)

func main() {
	db := utils.ConnectDb()
	utils.MigrateDb(db)

	if os.Getenv("APP_ENV") != "production" {
		utils.SeedTasks(db)
	}

	app := internal.App(db)
	app.Listen(":3000")

	fmt.Println("App listen to port 3000!")
}
