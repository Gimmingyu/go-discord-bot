package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var (
	BotToken       = flag.String("token", "", "Bot token")
	GuildID        = flag.String("guild", "", "Test guild ID")
	VoiceChannelID = flag.String("voice", "", "Test voice channel ID")
)

func init() {
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + *BotToken)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(pingpong)
}
