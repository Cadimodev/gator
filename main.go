package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cadimodev/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	fmt.Println("Welcome to Gator!")
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err.Error())
	}

	programState := &state{
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]commandHandler),
	}

	cmds.initCommands()

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	cmd := command{Name: cmdName, Args: cmdArgs}

	err = cmds.run(programState, cmd)
	if err != nil {
		log.Fatal(err)
	}

}
