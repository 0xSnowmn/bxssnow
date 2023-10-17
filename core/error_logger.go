package core

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

var errorMsg string

func LogErrorDiscord(errTxt string) {
	dg, err := discordgo.New("Bot " + os.Getenv("Bot_Token"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	errorMsg = errTxt
	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandlerOnce(logError)

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

func logError(session *discordgo.Session, ready *discordgo.Ready) {
	session.ChannelMessageSend(os.Getenv("Error_Channel"), errorMsg)
}
