package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	fmt.Println("Database Host : ", os.Getenv("DB_HOST"))
	fmt.Println("Database Name : ", os.Getenv("DB_NAME"))
	fmt.Println("Database User : ", os.Getenv("DB_USER"))
	fmt.Println("Database Password : ", os.Getenv("DB_PASSWORD"))
}
