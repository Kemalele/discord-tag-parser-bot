package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Kemalele/discord-tag-parser-bot/internal/config"
	"github.com/Kemalele/discord-tag-parser-bot/internal/handlers"

	"github.com/bwmarrin/discordgo"
)

const TOKEN_PREFIX = "Bot "

func main() {
	c := new(config.Config)
	err := c.Load(config.ServiceName)
	if err != nil {
		log.Fatal(err)
	}

	err = listen(c.Token)
	if err != nil {
		log.Fatal("Error listening bot session", err)
	}

}

func listen(token string) error {
	session, err := discordgo.New(TOKEN_PREFIX + token)
	if err != nil {
		return err
	}

	session.AddHandler(handlers.HandleInteractions)
	// In this example, we only care about receiving message events.
	session.Identify.Intents = discordgo.IntentsGuildMessages

	err = session.Open()
	if err != nil {
		return err
	}

	defer session.Close()
	defer fmt.Println("Closing discord session...")

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	return nil
}
