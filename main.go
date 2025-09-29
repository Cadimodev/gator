package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/cadimodev/gator/internal/config"
	"github.com/cadimodev/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

const dbDriverName = "postgres"

func main() {
	fmt.Println("Welcome to Gator!")
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open(dbDriverName, cfg.DbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	programState := &state{
		db:  dbQueries,
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
