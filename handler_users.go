package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	for _, user := range users {

		if s.cfg.CurrentUserName == user.Name {
			fmt.Printf("%s (current)\n", user.Name)
		} else {
			fmt.Printf("%s\n", user.Name)
		}
	}

	fmt.Println("User switched successfully!")

	return nil
}
