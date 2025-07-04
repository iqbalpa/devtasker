package main

import (
	"devtasker/internal/repository"
	"devtasker/internal/service"
	"fmt"
)

func main() {
	fmt.Println("Welcome!")
	var tr repository.ITaskRepository = repository.New()
	var ts service.ITaskService = service.New(&tr)
}
