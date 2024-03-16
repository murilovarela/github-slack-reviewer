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
	Pullrequests []Pullrequest
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
	Pullrequests   []*Pullrequest `gorm:"many2many:pullrequest_reviewers;"`
}

type Message struct {
	gorm.Model
	OrganizationID int
	SlackRef       string `gorm:"unique"`
	Content        string
	MessageType    string
	PullrequestID  int
}

type Pullrequest struct {
	gorm.Model
	OrganizationID  int
	GithubRef       string `gorm:"unique"`
	LatestMessageID int
	Reviewers       []*Reviewer `gorm:"many2many:pullrequest_reviewers;"`
	Messages        []Message
}
