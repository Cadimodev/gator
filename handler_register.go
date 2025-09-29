package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cadimodev/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully!")
	fmt.Println("ID: ", user.ID)
	fmt.Println("Created At: ", user.CreatedAt)
	fmt.Println("Updated At: ", user.UpdatedAt)
	fmt.Println("Name: ", user.Name)

	return nil
}
