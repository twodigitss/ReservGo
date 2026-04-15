package configs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

var	(
	DB_URL string
	DB_APIK string
	URL string
	ENV string
)

func LoadEnv(){
	_, filename, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(filename), "..")
	envPath := filepath.Join(root, ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Printf("Could not load .env from %s: %v\n", envPath, err)
	}

	DB_URL = os.Getenv("DB_POOL_URL")
	DB_APIK = os.Getenv("DB_APIK")
	PORT := os.Getenv("PORT")
	ENV = os.Getenv("ENV")
	DEV := os.Getenv("DEV_URL")
	PROD := os.Getenv("PROD_URL")

	switch ENV {
		case "development":{ URL = fmt.Sprintf("%s:%s", DEV, PORT) }
		case "production": { URL = PROD }
		default: { URL = fmt.Sprintf("%s:%s", DEV, PORT) }
	}

}
