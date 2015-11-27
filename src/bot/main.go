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

func (b *Bot) SetMe(user *slack.UserDetails) {
	b.MeID = user.ID
	log.Printf("Connect! I am %s (%s)\n", user.Name, user.ID)
}

func (b *Bot) IsMe(id string) bool {
	return b.MeID == id
}

func (b *Bot) IsDM(channel string) bool {
	return regexp.MustCompile("^D.*").MatchString(channel)
}

func (b *Bot) AmIMentioned(text string) bool {
	return regexp.MustCompile("<@" + b.MeID + ">").MatchString(text)
}

func (b *Bot) ReturnGreeting(name string) string {
	greetings := []string{
		"Good day to you Governor!",
		"Oh, yes! Good day!",
		"Greetings, it is better to be alone than in bad company",
		"I'm not much for pleasantries today :sob:",
		fmt.Sprintf("Hello <@%s>", name),
		fmt.Sprintf("Mother of the queen you're chipper today, <@%s>", name),
		":hear_no_evil: :see_no_evil: :speak_no_evil: ",
	}

	return greetings[rand.Intn(len(greetings))]
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
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
				bot.SetMe(ev.Info.User)

			case *slack.MessageEvent:
				// if I'm sent a DM or mentioned return a random greeting
				if bot.AmIMentioned(ev.Msg.Text) || (bot.IsDM(ev.Channel) && !bot.IsMe(ev.Msg.User)) {
					log.Printf("Received message \"%s\" from %s", ev.Msg.Text, ev.Msg.User)
					rtm.SendMessage(rtm.NewOutgoingMessage(bot.ReturnGreeting(ev.Msg.User), ev.Channel))
					log.Printf("I was courteous to %s\n", ev.Msg.User)
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
