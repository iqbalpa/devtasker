package main

import (
	"devtasker/internal"
	"devtasker/internal/utils"
	"fmt"
)

func main() {
	db := utils.ConnectDb()
	utils.MigrateDb(db)

	app := internal.App(db)
	app.Listen(":3000")

	fmt.Println("App listen to port 3000!")
}
