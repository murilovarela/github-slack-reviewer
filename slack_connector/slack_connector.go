// Service slack implements a github reviewer slack bot.
package slack_connector

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

// This uses Encore's built-in secrets manager, learn more: https://encore.dev/docs/primitives/secrets
var secrets struct {
	SlackSigningSecret string
}

//encore:service
type Service struct {
	svc     SlackService
	Secrets struct {
		SlackSigningSecret string
	}
}

func initService() (*Service, error) {
	prChan := make(chan PullRequestMessage)
	return &Service{
		svc:     NewSlackService(prChan),
		Secrets: secrets,
	}, nil
}

// encore:api public raw method=POST path=/slack/webhook
func (s *Service) SlackWebhook(w http.ResponseWriter, req *http.Request) {

	var body, err = io.ReadAll(req.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	secretVerifier, err := slack.NewSecretsVerifier(req.Header, s.Secrets.SlackSigningSecret)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if _, err := secretVerifier.Write(body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := secretVerifier.Ensure(); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text")
		w.Write([]byte(r.Challenge))
	}

	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		innerEvent := slackevents.EventsAPIInnerEvent{Type: eventsAPIEvent.InnerEvent.Type, Data: eventsAPIEvent.InnerEvent.Data}

		s.svc.HandleEvent(innerEvent)
	}

	w.WriteHeader(http.StatusOK)
}
