package main

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func pingpong(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.EqualFold(m.Content, "ping") {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if strings.EqualFold(m.Content, "pong") {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
