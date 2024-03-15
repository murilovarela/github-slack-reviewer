package slack_connector

import (
	"log"
	"reflect"
	"strings"

	"github.com/slack-go/slack/slackevents"
)

type SlackService interface {
	HandleEvent(innerEvent slackevents.EventsAPIInnerEvent)
	HandleMessagesEvent(ev slackevents.MessageEvent)
}

type PullRequestMessage struct {
	Repository string
	Author     string
	Link       string
}

type slackService struct {
	messages chan PullRequestMessage
}

func NewSlackService(messagesChan chan PullRequestMessage) SlackService {
	return &slackService{
		messages: messagesChan,
	}
}

func (s *slackService) HandleEvent(innerEvent slackevents.EventsAPIInnerEvent) {
	switch ev := innerEvent.Data.(type) {
	case *slackevents.MessageEvent:
		{
			s.HandleMessagesEvent(*ev)
		}
	}
}

func (s *slackService) HasGithubLink(text string) bool {
	return strings.Contains(text, "github.com")
}

func logStruct(s interface{}) {
	val := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)

	log.Println("Struct type:", typ.Name())

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name
		fieldValue := field.Interface()
		log.Printf("%s: %v\n", fieldName, fieldValue)
	}
}

func (s *slackService) HandleMessagesEvent(ev slackevents.MessageEvent) {
	logStruct(ev)
	if ev.ChannelType != "channel" || ev.BotID != "" {
		return
	}

	log.Printf("Message received: %s", ev.Channel)

	if ev.SubType == "message_changed" {
		if ev.Message.SubType == "message_deleted" {
			// message was deleted
			log.Printf("Message deleted")
			return
		}

		// message was edited
		log.Printf("Message edited")
		return
	}

	if ev.SubType == "message_deleted" {
		// message was deleted
		log.Printf("Message deleted")
		return
	}

	// new message
	log.Printf("Message is new")
}
