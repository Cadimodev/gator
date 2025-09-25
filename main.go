package main

import (
	"fmt"

	"github.com/cadimodev/gator/internal/config"
)

func main() {
	fmt.Println("Welcome to Gator!")
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = cfg.SetUser("Carlos")
	if err != nil {
		fmt.Println(err.Error())
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Config File Content:")
	fmt.Println("CurrentUserName: ", cfg.CurrentUserName)
	fmt.Println("DbURL: ", cfg.DbURL)
}
