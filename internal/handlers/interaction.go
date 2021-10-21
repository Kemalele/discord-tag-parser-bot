package handlers

import (
	"log"
	"strings"

	"github.com/Kemalele/discord-tag-parser-bot/internal/common"
	"github.com/bwmarrin/discordgo"
)

func HandleInteractions(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch m.Content {
	case "!start":
		handleStart(s, m)
	case "!find":
		handleFind(s, m)
	}

}

func handleStart(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg, err := s.ChannelMessageSend(m.ChannelID, "hi!")
	if err != nil {
		log.Println(err)
	}

	log.Println("Some message?", msg)
}

func handleFind(s *discordgo.Session, m *discordgo.MessageCreate) {
	history, err := s.ChannelMessages(common.KNOWLEDGE_BASE_ID, 100, "", "", "")
	if err != nil {
		log.Println(err)
		return
	}

	userChannel, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		log.Println(err)
		return
	}

	for _, msg := range history {
		if strings.Contains(msg.Content, "#") {
			s.ChannelMessageSend(userChannel.ID, msg.Content)
		}
	}
}
