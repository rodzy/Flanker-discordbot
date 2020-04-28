package flankbot

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rodzy/flanker-discordbot/config"
	reader "github.com/rodzy/flanker-discordbot/pdf"
)

//FlankerID is the id for FlankBot
var FlankerID string
var flankSession *discordgo.Session

//FlankStart func to init the connection
func FlankStart() {
	flankSession, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//Defintion of the bot discord user
	us, err := flankSession.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}
	//Setting the user id to the bot
	FlankerID = us.ID

	flankSession.AddHandler(MessageHandler)

	err = flankSession.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Flanker is running!!!")
}

//MessageHandler to handle all the commands and text from the pdf
func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, config.BotCom) {
		if m.Author.ID == FlankerID {
			return
		}
	}

	if m.Content == "$help" {
		//Message introduction
		embed := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{
				Name:    "Flankerbot",
				URL:     "https://github.com/rodzy",
				IconURL: "https://cdn.discordapp.com/avatars/703454326722396161/768085d0b4991979ffda4218a977364e.webp?size=128",
			},
			Color:       0x66ccff,
			Description: "Git is the open source distributed version control system that facilitates GitHub activities on your laptop or desktop.\n To start just write: ``$<Command>``",
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "Fundamentals",
					Value:  "```css\nCommand: $Basics:``````The most common comands for daily bases```",
					Inline: false,
				},
				&discordgo.MessageEmbedField{
					Name: "Working with branches",
					Value: "```css\n Command: $Branches:``````Branches are an important part of working with Git.```",
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "Creating/cloning from GitHub",
					Value:  "```css\n Command: $Create:``````When starting out with a new repository, you only need to do it once.```",
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name: "Syncronization",
					Value: "```css\n Command: $Sync:``````Synchronize your local repository with the remote repository on GitHub.com```",
					Inline: false,
				},
				&discordgo.MessageEmbedField{
					Name:   "Managing changes & redo commits",
					Value:  "```css\n Command: $Changes:``````Browse and inspect the evolution of project files, erase mistakes and craft replacement history```",
					Inline: true,
				},
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://gitforwindows.org/img/gwindows_logo.png",
			},
			Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
			Title:     "Flankerbot - Git Commands",
		}
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}
	fmt.Println(reader.MainTitle)

}
