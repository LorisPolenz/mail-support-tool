package helper

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func BuildFolderPath(path string, currentYear int) string {
	if runtime.GOOS == "windows" {
		return fmt.Sprintf("%v\\%v\\", path, currentYear)
	}

	return fmt.Sprintf("%v/%v/", path, currentYear)
}

func CheckFilePath(path string) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Fatal("Folder does not exist.")
	}
}
