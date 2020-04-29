package main

import (
	"fmt"

	"github.com/rodzy/flanker-discordbot/config"
	"github.com/rodzy/flanker-discordbot/flankbot"
	reader "github.com/rodzy/flanker-discordbot/pdf"
)

func main() {
	err := reader.ReadPdf()
	if err != nil {
		fmt.Println("Can't read sorry")
	}

	er := config.ReadConfig()
	if er != nil {
		fmt.Print("Can't connect sorry")
	}
	flankbot.FlankStart()

}
