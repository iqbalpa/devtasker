package utils

import (
	"devtasker/internal/model"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"gorm.io/gorm"
)

func readFile() []byte {
	dir, _ := os.Getwd()
	fpath := fmt.Sprintf("%s/internal/utils/data.json", dir)
	jsonFile, err := os.Open(fpath)
	if err != nil {
		fmt.Println("Failed to read data.json", err)
	}
	defer jsonFile.Close()

	byteVal, _ := io.ReadAll(jsonFile)
	return byteVal
}

func getDummyData() []*model.Task {
	var tasks []*model.Task
	byteVal := readFile()
	json.Unmarshal(byteVal, &tasks)
	return tasks
}

func SeedTasks(db *gorm.DB) {
	data := getDummyData()
	db.Save(data)
}
