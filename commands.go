package main

import "errors"

type command struct {
	Name string
	Args []string
}

type commandHandler func(*state, command) error

type commands struct {
	registeredCommands map[string]commandHandler
}

func (c *commands) run(s *state, cmd command) error {

	cmdHandler, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}

	return cmdHandler(s, cmd)
}

func (c *commands) register(name string, cmdHandler commandHandler) {
	c.registeredCommands[name] = cmdHandler
}

func (c *commands) initCommands() {
	c.register("login", handlerLogin)
	c.register("register", handlerRegister)
	c.register("reset", handlerReset)
	c.register("users", handlerUsers)
	c.register("agg", handlerAgg)
	c.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	c.register("feeds", handlerFeeds)
	c.register("follow", middlewareLoggedIn(handlerFollow))
	c.register("following", middlewareLoggedIn(handlerListFeedFollows))
	c.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	c.register("browse", middlewareLoggedIn(handlerBrowse))
}
