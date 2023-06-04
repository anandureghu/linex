package helper

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func ReadJSONToken() map[string][]string {
	byteValue, err := os.ReadFile(getFilePath())
	if err != nil {
		fmt.Println(err)
	}
	var result map[string][]string
	json.Unmarshal([]byte(byteValue), &result)

	return result
}

func WriteJSON(data []byte) (bool, error) {
	err := os.WriteFile(getFilePath(), data, fs.ModeAppend)
	if err != nil {
		return false, err
	}
	return true, nil
}

func getFilePath() string {
	dir := ""

	switch runtime.GOOS {
	case "linux":
		dir = fmt.Sprintf("/home/%s/.linex/lib", os.Getenv("USER"))
	}

	fileName := "commands.json"
	file := filepath.Join(dir, filepath.Base(fileName))

	_, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Creating folder")
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Fatal("folder creation error >>> ", err)
		}
	}
	_, err = os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal("file creation error >>> ", err)
	}

	return file
}

func SplitString(input string) []string {
	var args []string
	var currentArg string
	var inQuotes bool

	for _, char := range input {
		switch {
		case char == ' ' && !inQuotes:
			if currentArg != "" {
				args = append(args, currentArg)
				currentArg = ""
			}
		case char == '"' && !inQuotes:
			inQuotes = true
		case char == '"' && inQuotes:
			inQuotes = false
		default:
			currentArg += string(char)
		}
	}

	if currentArg != "" {
		args = append(args, currentArg)
	}

	return args
}
