package github_connector

import (
	"log"
	"net/http"

	"github.com/google/go-github/v60/github"
)

var secrets struct {
	GithubWebhookSecret string
}

type GithubService interface {
	HandleWebhook(w http.ResponseWriter, req *http.Request)
	HandleCommitCommentEvent(ev github.CommitCommentEvent)
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

func (s *githubService) HandleCommitCommentEvent(ev github.CommitCommentEvent) {
	log.Printf("Commit Comment: %+v", ev)
}

// github webhook
// encore:api public raw method=POST path=/github/webhook
func (s *githubService) HandleWebhook(w http.ResponseWriter, req *http.Request) {
	payload, err := github.ValidatePayload(req, []byte(secrets.GithubWebhookSecret))
	if err != nil {
		log.Printf("Error: %s", err)
	}
	event, err := github.ParseWebHook(github.WebHookType(req), payload)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	switch event := event.(type) {
	case *github.CommitCommentEvent:
		s.HandleCommitCommentEvent(*event)
	case *github.CreateEvent:
		log.Printf("Commit Comment: %+v", event)
	}

	w.WriteHeader(http.StatusOK)
}

// get github merge request details
func (s *githubService) GetMergeRequestDetails() {
	// get merge request details

}
