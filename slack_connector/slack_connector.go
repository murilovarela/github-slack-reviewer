// Service slack implements a github reviewer slack bot.
package slack_connector

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

// This uses Encore's built-in secrets manager, learn more: https://encore.dev/docs/primitives/secrets
var secrets struct {
	SlackSigningSecret string
}

func HandleMessagesEvent(ev *slackevents.MessageEvent) {
	if ev.ChannelType != "channel" {
		return
	}

	log.Printf("MessageEvent: %v", ev.Message.Text)
}

func HandleEvents(w http.ResponseWriter, req *http.Request) {
	var body, err = io.ReadAll(req.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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
		innerEvent := eventsAPIEvent.InnerEvent

		switch ev := innerEvent.Data.(type) {
		case *slackevents.MessageEvent:
			HandleMessagesEvent(ev)
		}
	}
}

// encore:api public raw method=POST path=/slack/events
func SubscribeToEvents(w http.ResponseWriter, req *http.Request) {
	var body, err = io.ReadAll(req.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req.Body = io.NopCloser(bytes.NewBuffer(body))

	secretVerifier, err := slack.NewSecretsVerifier(req.Header, secrets.SlackSigningSecret)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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

	HandleEvents(w, req)
}
