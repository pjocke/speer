package main

import (
        "fmt"
	"strings"
	"log"
	"os"
	"io/ioutil"
        "github.com/nlopes/slack"
)

func main() {
	logger := log.New(os.Stdout, "SPEER:", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)

	log.Println("Initializing.")

	cwd, _ := os.Getwd()
	key, err := ioutil.ReadFile(cwd + "/api.key")
	if err != nil {
		log.Fatalf("Missing API key: %s\n", cwd + "/api.key")
	}

	api := slack.New(strings.TrimSpace(string(key)))
	api.SetDebug(false)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

    for msg := range rtm.IncomingEvents {
    	switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				log.Println("Connected.")

			case *slack.MessageEvent:
				info := rtm.GetInfo()
				if ev.User != info.User.ID && strings.Contains(ev.Text, fmt.Sprintf("<@%s>", info.User.ID)) {
					log.Printf("Triggered by user %s.\n", ev.User)
					rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprintf("<@%s> :eggplant::sweat_drops::princess::skin-tone-2:", ev.User), ev.Channel))
				}

			case *slack.RTMError:
				log.Printf("Error %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				log.Fatal("Invalid credentials.")

			default:
				// noop
		}
	}
}
