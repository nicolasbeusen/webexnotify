package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	webexteams "github.com/jbogarin/go-cisco-webex-teams/sdk"
)

var Client *webexteams.Client

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var webexMessage string

	var ToPersonEmail string
	flag.StringVar(&ToPersonEmail, "to", "", "Webex User")

	var webexToken string
	flag.StringVar(&webexToken, "token", "", "Webex Token")

	var messageTemplate string
	flag.StringVar(&messageTemplate, "template", "", "Message Template")

	flag.Parse()

	Client = webexteams.NewClient()
	Client.SetAuthToken(webexToken)

	// read the host and port from pipeline
	for scanner.Scan() {
		webexMessage = scanner.Text()

		message := &webexteams.MessageCreateRequest{
			Text:          strings.Replace(messageTemplate, "{{ data }}", webexMessage, -1),
			ToPersonEmail: ToPersonEmail,
		}
		newTextMessage, _, err := Client.Messages.CreateMessage(message)
		if err != nil {
			fmt.Println("Error POST:", newTextMessage.ID, newTextMessage.Text, newTextMessage.Created)
			continue
		}

	}
}
