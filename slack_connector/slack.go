package slack_connector

import (
	"log"
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

func (s *slackService) HandleMessagesEvent(ev slackevents.MessageEvent) {
	if ev.ChannelType != "channel" {
		return
	}

	log.Printf("Message received: %s", ev.Channel)
}
