package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func createAmazingEvent(s *discordgo.Session) *discordgo.GuildScheduledEvent {
	startingTime := time.Now().Add(1 * time.Hour)
	endingTime := startingTime.Add(30 * time.Minute)
	scheduledEvent, err := s.GuildScheduledEventCreate(*GuildID, &discordgo.GuildScheduledEventParams{
		Name:               "Amazing Event",
		Description:        "This event will start in 1 hour and last 30 minutes",
		ScheduledStartTime: &startingTime,
		ScheduledEndTime:   &endingTime,
		EntityType:         discordgo.GuildScheduledEventEntityTypeVoice,
		ChannelID:          *VoiceChannelID,
		PrivacyLevel:       discordgo.GuildScheduledEventPrivacyLevelGuildOnly,
	})
	if err != nil {
		log.Printf("Error creating scheduled event: %v", err)
		return nil
	}

	fmt.Println("Created scheduled event:", scheduledEvent.Name)
	return scheduledEvent
}

func transformEventToExternalEvent(s *discordgo.Session, event *discordgo.GuildScheduledEvent) {
	scheduledEvent, err := s.GuildScheduledEventEdit(*GuildID, event.ID, &discordgo.GuildScheduledEventParams{
		Name:       "Amazing Event @ Discord Website",
		EntityType: discordgo.GuildScheduledEventEntityTypeExternal,
		EntityMetadata: &discordgo.GuildScheduledEventEntityMetadata{
			Location: "https://discord.com",
		},
	})
	if err != nil {
		log.Printf("Error during transformation of scheduled voice event into external event: %v", err)
		return
	}

	fmt.Println("Created scheduled event:", scheduledEvent.Name)
}
