package bot

import (
	"fmt"

	"github.com/akhil/discord-ping/config"
	"github.com/bwmarrin/discordgo"
)

var botID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	botID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is Running! good job dumbass")

}

// msg handler shit idk
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == botID {
		return
	}

	if m.Content == "yo" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "wsp w yew gang")
	}
}
