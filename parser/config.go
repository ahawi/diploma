
package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	godotenv.Load()
}

func GetDBConnStr() string {
	return os.Getenv("DB_CONN")
}
