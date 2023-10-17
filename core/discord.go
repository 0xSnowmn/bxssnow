package core

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

var Msg string
var FileN string
var Error string

func HitDiscord() {
	dg, err := discordgo.New("Bot " + os.Getenv("Bot_Token"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandlerOnce(notify)

	dg.Identify.Intents = discordgo.IntentsGuildMessages
	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildMembers | discordgo.IntentsGuildPresences

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Cleanly close down the Discord session.
	defer dg.Close()
}

func notify(session *discordgo.Session, r *discordgo.Ready) {
	ff, errors := os.Open(FileN)
	if errors != nil {
		LogErrorDiscord(errors.Error())
	}
	_, err := session.ChannelMessageSendComplex(os.Getenv("Fire_Channel"), &discordgo.MessageSend{
		File: &discordgo.File{
			Name:   FileN,
			Reader: bufio.NewReader(ff),
		},
		Content: Msg,
	})

	if err != nil {
		LogErrorDiscord(errors.Error())
	}
}
