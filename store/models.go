package store

import (
	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	Name         string
	Email        string `gorm:"unique"`
	Config       Config
	Reviewers    []Reviewer
	Messages     []Message
	PullRequests []PullRequest
}

type Config struct {
	gorm.Model
	OrganizationID int
	TimeToSummary  int
	TimeToReminder int
}

type Reviewer struct {
	gorm.Model
	OrganizationID int
	GithubID       string
	SlackID        string
	PullRequests   []PullRequest `gorm:"many2many:pull_request_reviewers;"`
}

type Message struct {
	gorm.Model
	OrganizationID int
	SlackRef       string `gorm:"unique"`
	Content        string
	MessageType    string
	PullRequestID  int
}

type PullRequest struct {
	gorm.Model
	OrganizationID  int
	GithubRef       string `gorm:"unique"`
	LatestMessageID int
	Reviewers       []Reviewer
	Approvers       []Reviewer
	Messages        []Message
}
