package helper

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/user"
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
	fmt.Println(getFilePath())
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
		username, _ := user.Current()
		dir = fmt.Sprintf("/home/%s/.linex/lib", username.Username)
	}

	fileName := "commands.json"
	file := filepath.Join(dir, filepath.Base(fileName))

	_, err := os.Open(file)
	if err != nil {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Fatal("folder creation error >>> ", err)
		}
		_, err = os.Create(file)
		if err != nil {
			log.Fatal("file creation error >>> ", err)
		}
	}

	return file
}
