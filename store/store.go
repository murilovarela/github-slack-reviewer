// writes and reads from the store
package store

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var secrets struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

type StoreService interface {
	GetPullRequestByOrganizationIdAndRef(organizationID int, githubRef string) (Pullrequest, error)
	CreatePullRequest(organizationID int, githubRef string) (Pullrequest, error)
	CreateOrganization(organization Organization, config Config) (Organization, error)
}

//encore:service
type storeService struct {
	db *gorm.DB
}

func NewStoreService(customDsn *string) (StoreService, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", secrets.DBHost, secrets.DBUser, secrets.DBPassword, secrets.DBName, secrets.DBPort)
	if customDsn != nil {
		dsn = *customDsn
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return &storeService{
		db: db,
	}, err
}

func (s *storeService) GetPullRequestByOrganizationIdAndRef(organizationID int, githubRef string) (Pullrequest, error) {
	pullRequest := Pullrequest{
		OrganizationID: organizationID,
		GithubRef:      githubRef,
	}
	result := s.db.Where(pullRequest).First(&pullRequest)

	return pullRequest, result.Error
}

func (s *storeService) CreatePullRequest(organizationID int, githubRef string) (Pullrequest, error) {
	pullRequest, err := s.GetPullRequestByOrganizationIdAndRef(organizationID, githubRef)

	if err == nil {
		return pullRequest, fmt.Errorf("pull request already exists")
	}

	pullRequest = Pullrequest{
		OrganizationID: organizationID,
		GithubRef:      githubRef,
	}

	result := s.db.Create(&pullRequest)

	return pullRequest, result.Error
}

func (s *storeService) CreateOrganization(organization Organization, config Config) (Organization, error) {
	organization.Config = config
	result := s.db.Create(&organization)

	return organization, result.Error
}

// func (s *storeService) Delete() {}

// func (s *storeService) Update() {}

// func (s *storeService) List() {}
