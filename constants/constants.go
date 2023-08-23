package constants

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kunxl-gg/no-issues/helpers"
)

var (
	OPENAI_API_KEY  string
	GH_API_KEY		string
	GITHUB_API_BASE_URL = "https://api.github.com"
	OPENAI_BASE_URL = "https://api.openai.com"
)

// Initializing .env variables
func InitialiseEnvironmentVariables() {

	// loading .env file
	err := godotenv.Load()
	helpers.CheckError(err)

	// read the variables
	OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")
	GH_API_KEY = os.Getenv("GH_API_KEY")

}