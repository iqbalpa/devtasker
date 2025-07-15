package utils

import (
	"devtasker/internal/model"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"gorm.io/gorm"
)

const (
	UserDataPath string = "user.json"
	TaskDataPath string = "task.json"
)

func readFile(fname string) []byte {
	dir, _ := os.Getwd()
	fpath := fmt.Sprintf("%s/internal/data/%s", dir, fname)
	jsonFile, err := os.Open(fpath)
	if err != nil {
		fmt.Println("Failed to read data.json", err)
	}
	defer jsonFile.Close()

	byteVal, _ := io.ReadAll(jsonFile)
	return byteVal
}

func getDummyData[T any](fname string) []T {
	var data []T
	byteVal := readFile(fname)
	json.Unmarshal(byteVal, &data)
	return data
}

func SeedTasks(db *gorm.DB) {
	tasks := getDummyData[model.Task](TaskDataPath)
	db.Save(tasks)
	users := getDummyData[model.User](UserDataPath)
	db.Save(users)
}
