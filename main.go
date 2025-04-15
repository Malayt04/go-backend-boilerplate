package main

import (
	"fmt"
	"log"

	"github.com/go-backend/bootstrap"
)

func main() {
	env, err := bootstrap.NewEnv()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println(env)
}
