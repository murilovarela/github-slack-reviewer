package github_connector

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/go-github/v60/github"
)

var secrets struct {
	GithubWebhookSecret string
}

type GithubService interface {
	HandleWebhook(w http.ResponseWriter, req *http.Request)
	HandleCommitCommentEvent(ev github.CommitCommentEvent)
	GetMergeRequestDetails(owner string, repo string, id int) (*github.PullRequest, *github.Response, error)
	GetMergeRequestDetailsFromUri(uri string) (*github.PullRequest, *github.Response, error)
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

func (s *githubService) GetIdsFromUri(uri string) (owner string, repo string, id int, err error) {
	trimmedUri := strings.TrimPrefix(uri, "https://github.com/")
	parts := strings.Split(trimmedUri, "/")

	if len(parts) < 4 {
		err = fmt.Errorf("invalid GitHub URL: %s", uri)
		return
	}

	owner = parts[0]
	repo = parts[1]
	id, err = strconv.Atoi(parts[3])
	if err != nil {
		log.Printf("invalid GitHub URL: %s", uri)
		return
	}

	return
}

func (s *githubService) GetMergeRequestDetailsFromUri(uri string) (*github.PullRequest, *github.Response, error) {
	owner, repo, id, err := s.GetIdsFromUri(uri)

	if err != nil {
		log.Printf("Error: %s", err)
		return nil, nil, err
	}

	return s.GetMergeRequestDetails(owner, repo, id)
}

// get github merge request details
func (s *githubService) GetMergeRequestDetails(owner string, repo string, id int) (*github.PullRequest, *github.Response, error) {
	// get merge request details from github
	ctx := context.Background()
	return s.githubClient.PullRequests.Get(ctx, owner, repo, id)
}
