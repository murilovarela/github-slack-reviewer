package github_connector

import (
	"fmt"
	"net/http"

	webhooks "github.com/go-playground/webhooks/v6/github"
	"github.com/google/go-github/v60/github"
)

var secrets struct {
	GithubWebhookSecret string
}

type GithubService interface {
	HandleWebhook(w http.ResponseWriter, req *http.Request)
	HandleEvent()
	GetMergeRequestDetails()
}

//encore:service
type githubService struct {
	githubClient *github.Client
}

func NewGithubService() GithubService {
	client := github.NewClient(nil)

	return &githubService{
		githubClient: client,
	}
}

func (s *githubService) HandleEvent() {}

// github webhook
// encore:api public raw method=POST path=/github/webhook
func (s *githubService) HandleWebhook(w http.ResponseWriter, req *http.Request) {
	hook, _ := webhooks.New(webhooks.Options.Secret(secrets.GithubWebhookSecret))
	payload, err := hook.Parse(req, webhooks.ReleaseEvent, webhooks.PullRequestEvent)
	if err != nil {
		if err == webhooks.ErrEventNotFound {
			fmt.Printf("%+v", err)
			// ok event wasn't one of the ones asked to be parsed
		}
	}
	switch payload := payload.(type) {

	case *webhooks.PullRequestPayload:
		// Do whatever you want from here...
		fmt.Printf("%+v", payload)
	}
}

// get github merge request details
func (s *githubService) GetMergeRequestDetails() {
	// get merge request details

}
