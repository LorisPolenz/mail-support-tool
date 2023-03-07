package helper

import (
	"log"
	"os"
)

func ValidateEnVariables() {

	if os.Getenv("OMA_TOOL_YEAR") == "" {
		log.Fatal("OMA_TOOL_YEAR is not set")
	}

	if os.Getenv("OMA_TOOL_FOLDER_PATH") == "" {
		log.Fatal("OMA_TOOL_FOLDER_PATH is not set")
	}

	// if os.Getenv("OMA_TOOL_RECEIVERS") == "" {
	// 	log.Fatal("OMA_TOOL_RECEIVERS is not set")
	// }

	// if os.Getenv("OMA_TOOL_EMAIL_USERNAME") == "" {
	// 	log.Fatal("OMA_TOOL_EMAIL_USERNAME is not set")
	// }

	// if os.Getenv("OMA_TOOL_EMAIL_PASSWORD") == "" {
	// 	log.Fatal("OMA_TOOL_EMAIL_PASSWORD is not set")
	// }
}
