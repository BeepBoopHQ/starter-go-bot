package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"time"

	"github.com/nlopes/slack"
)

type Bot struct {
	MeID string
}

func (b *Bot) random(max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max)
}

func (b *Bot) setMe(id string) {
	b.MeID = id
	log.Println("Connected as " + id)
}

func (b *Bot) isMe(id string) bool {
	return b.MeID == id
}

func (b *Bot) amIMentioned(text string) bool {
	return regexp.MustCompile("<@" + b.MeID + ">").MatchString(text)
}

func (b *Bot) returnGreeting(name string) string {
	greetings := []string{
		"Good day to you Governor!",
		"Oh, yes! Good day!",
		"Greetings, it is better to be alone than in bad company",
		"I'm not much for pleasantries today :sob:",
		fmt.Sprintf("Hello <@%s>", name),
		fmt.Sprintf("Mother of the queen you're chipper today, <@%s>", name),
		":hear_no_evil: :see_no_evil: :speak_no_evil: ",
	}

	return greetings[b.random(len(greetings))]
}

func main() {
	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		log.Fatal("SLACK_TOKEN is required")
	}

	api := slack.New(token)
	bot := Bot{}

	rtm := api.NewRTM()
	go rtm.ManageConnection()

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				bot.setMe(ev.Info.User.ID)

			case *slack.MessageEvent:
				if bot.amIMentioned(ev.Msg.Text) {
					rtm.SendMessage(rtm.NewOutgoingMessage(bot.returnGreeting(ev.Msg.User), ev.Channel))
				}

			case *slack.RTMError:
				log.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				log.Fatal("Invalid credentials")
				break Loop
			}
		}
	}
}
