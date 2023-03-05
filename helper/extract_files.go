package helper

import (
	"log"
	"os"
)

func Extract_file_paths(dir_path string) []string {
	file_names := []string{}

	files, err := os.ReadDir(dir_path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		file_names = append(file_names, file.Name())
	}

	return file_names

}
