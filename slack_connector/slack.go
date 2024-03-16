package slack_connector

import (
	"log"
	"reflect"
	"strings"

	"encore.app/github_connector"
	"encore.app/store"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

type SlackService interface {
	HandleEvents(innerEvent slackevents.EventsAPIInnerEvent)
	HandleMessagesEvent(ev slackevents.MessageEvent)
	HandleNewMessage(channelRef string, messageRef string, ts string, text string)
	SendThreadMessage(channelRef string, ts string, message string) (messageRef string, err error)
}

type slackService struct {
	githubService github_connector.GithubService
	storeService  store.StoreService
	slackApiSvc   *slack.Client
}

func NewSlackService() SlackService {
	githubSvc := github_connector.NewGithubService()
	storeSvc, err := store.NewStoreService(nil)
	slackApiSvc := slack.New(secrets.SlackBotUserToken)

	if err != nil {
		log.Printf("Error: %s", err)
	}

	return &slackService{
		githubService: githubSvc,
		storeService:  storeSvc,
		slackApiSvc:   slackApiSvc,
	}
}

func (s *slackService) SendThreadMessage(channelRef string, ts string, message string) (messageRef string, err error) {
	messageRef, _, err = s.slackApiSvc.PostMessage(channelRef, slack.MsgOptionText(message, false), slack.MsgOptionTS(ts))

	if err != nil {
		log.Printf("Error sending message: %s", err)

	}
	return messageRef, err
}

func (s *slackService) HandleNewMessage(channelRef string, messageRef string, ts string, text string) {
	if strings.Contains(text, "https://github.com/") {
		// parse the message and get the github uri
		uri := strings.Split(text, "<")[1]
		uri = strings.Split(uri, ">")[0]

		pull, _, err := s.githubService.GetMergeRequestDetailsFromUri(uri)

		if err != nil {
			log.Printf("Error: %s", err)

			s.SendThreadMessage(channelRef, ts, "Pull request not found")
			return
		}

		_, err = s.storeService.CreatePullRequest(1, uri)

		if err != nil {
			log.Printf("Error: %s", err)

			// send error message to slack as reply
			return
		}

		// send message to slack
		// s.SendSlackMessage(pull, uri, messageId)

		log.Printf("Link: %s", uri)
		log.Printf("MessageId: %s", messageRef)
		log.Printf("reviewers: %+v", pull.RequestedReviewers)
		log.Printf("body: %+v", *pull.Body)
		log.Printf("mergeableState: %+v", *pull.MergeableState)
		log.Printf("title: %+v", *pull.Title)
		log.Printf("state: %+v", *pull.State)
		log.Printf("head: %+v", *pull.Head.Label)
		log.Printf("base: %+v", *pull.Base.Label)
	} else {
		return
	}
}

func (s *slackService) HandleEvents(innerEvent slackevents.EventsAPIInnerEvent) {
	switch ev := innerEvent.Data.(type) {
	case *slackevents.MessageEvent:
		{
			s.HandleMessagesEvent(*ev)
		}
	}
}

func (s *slackService) HandleMessagesEvent(ev slackevents.MessageEvent) {
	if ev.ChannelType != "channel" || ev.BotID != "" {
		return
	}

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
	s.HandleNewMessage(ev.Channel, ev.ClientMsgID, ev.TimeStamp, ev.Text)
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
