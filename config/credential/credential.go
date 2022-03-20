package credential

import (
	"os"
)

var (
	AppHost    string
	AppPort    string
	DbHost     string
	DbPort     string
	DbUsername string
	DbPassword string
	DbName     string
)

func CredentialsConfig() (err error) {
	AppHost = os.Getenv("APP_HOST")
	AppPort = os.Getenv("APP_PORT")
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbUsername = os.Getenv("DB_USERNAME")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbName = os.Getenv("DB_DATABASE")
	return err
}
