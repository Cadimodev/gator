package main

import (
	"context"
	"fmt"
)

const feedURL = "https://www.wagslane.dev/index.xml"

func handlerAgg(s *state, cmd command) error {

	feed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}

	fmt.Println("Title: ", feed.Channel.Title)
	fmt.Println("Description: ", feed.Channel.Description)
	for _, item := range feed.Channel.Item {
		fmt.Println("	Title: ", item.Title)
		fmt.Println("	Link: ", item.Link)
		fmt.Println("	Description: ", item.Description)
		fmt.Println("	PubDate: ", item.PubDate)
		fmt.Println("***********")
	}

	return nil
}
