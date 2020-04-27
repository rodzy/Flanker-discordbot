package flankbot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/rodzy/flanker-discordbot/config"
)
//FlankerID is the id for FlankBot
var FlankerID string
var flankSession *discordgo.Session

//FlankStart func to init the connection
func FlankStart()  {
	flankSession,err:=discordgo.New("Bot "+ config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//Defintion of the bot discord user
	us,err:=flankSession.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}
	//Setting the user id to the bot
	FlankerID=us.ID

	// flankSession.AddHandler()

	err=flankSession.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Flanker is running!!!")
}