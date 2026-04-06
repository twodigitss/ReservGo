package configs

import (
	"os"
	"path/filepath"
	"log"
	"github.com/joho/godotenv"
	"runtime"
)

var	(
	DB_URL string
	DB_APIK string
)

func LoadEnv(){
	_, filename, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(filename), "..")
	envPath := filepath.Join(root, ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Printf("Could not load .env from %s: %v\n", envPath, err)
	}

	DB_URL = os.Getenv("DB_URL")
	DB_APIK = os.Getenv("DB_APIK")

}

