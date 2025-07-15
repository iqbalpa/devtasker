package main

import (
	_ "devtasker/docs"
	"devtasker/internal"
	"devtasker/internal/utils"
	"fmt"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	db := utils.ConnectDb()
	utils.MigrateDb(db)

	if os.Getenv("APP_ENV") != "production" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			utils.SeedTasks(db)
		}()
	}

	app := internal.App(db)
	app.Listen(":3000")

	fmt.Println("App listen to port 3000!")

	wg.Wait()
}
