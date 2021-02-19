package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/skyapps-id/K8s-Style-Configmap-and-Secret/api/controllers"
)

var server = controllers.Server{}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func Run() {
	fmt.Println("Database Host : ", os.Getenv("DB_HOST"))
	fmt.Println("Database Name : ", os.Getenv("DB_NAME"))
	fmt.Println("Database User : ", os.Getenv("DB_USER"))
	fmt.Println("Database Password : ", os.Getenv("DB_PASSWORD"))

	server.Run(":8080")
}
