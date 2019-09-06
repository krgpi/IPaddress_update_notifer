package main

import (
	"fmt"
	"strings"

	"github.com/nlopes/slack"
)

const postURL = "https://hooks.slack.com/***"

func main() {
	api := slack.New("xoxb-***")
	rtm := api.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Println("EventReceived.")
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				if strings.Index(ev.Text, "ipaddr") != -1 {
					rtm.SendMessage(rtm.NewOutgoingMessage(string(getIPaddr()), ev.Channel))
				}
			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())
			case *slack.InvalidAuthEvent:
				fmt.Println("Invalid credentials")
				break Loop
			default:
				//Take no action
			}
		}
	}
}
