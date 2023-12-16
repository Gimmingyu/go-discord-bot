package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func dm(sess *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == sess.State.User.ID {
		return
	}

	if !strings.EqualFold(msg.Content, "ping me") && !strings.EqualFold(msg.Content, "pong me") {
		log.Println(msg.Content)
		return
	}

	// We create the private channel with the user who sent the message.
	channel, err := sess.UserChannelCreate(msg.Author.ID)
	if err != nil {
		log.Println("error creating channel:", err)
		_, _ = sess.ChannelMessageSend(
			msg.ChannelID,
			"Something went wrong while sending the DM!",
		)
		return
	}

	switch {
	case strings.EqualFold(msg.Content, "ping me"):
		_, err := sess.ChannelMessageSend(channel.ID, "pong")
		if err != nil {
			log.Println("error sending DM message:", err)
			_, _ = sess.ChannelMessageSend(
				msg.ChannelID,
				"Failed to send you a DM. "+
					"Did you disable DM in your privacy settings?",
			)
		}
	case strings.EqualFold(msg.Content, "pong me"):
		_, err := sess.ChannelMessageSend(channel.ID, "PONG")
		if err != nil {
			log.Println("error sending DM message:", err)
			_, _ = sess.ChannelMessageSend(
				msg.ChannelID,
				"Failed to send you a DM. "+
					"Did you disable DM in your privacy settings?",
			)
		}
	}
}
