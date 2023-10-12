package core

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

var Msg string

var Template = `
	
`

func S() {
	dg, err := discordgo.New("Bot " + os.Getenv("Token"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	fmt.Println("Bot " + os.Getenv("Token"))

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
	session.ChannelMessageSend("1161157889197625445", Msg)
}
