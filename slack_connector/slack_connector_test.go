package slack_connector_test

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"encore.app/slack_connector"
	"github.com/stretchr/testify/assert"
)

func TestEventWebhook(t *testing.T) {

		// Create a new HTTP request with an invalid event
		req := httptest.NewRequest(http.MethodPost, "/slack/webhook", nil)
		// add a header to make a invalid NewSecretsVerifier
		hash := hmac.New(sha256.New, []byte(s.Secrets.SlackSigningSecret))
		stimestamp := "1234567890"
		hash.Write([]byte(fmt.Sprintf("v0:%s:", stimestamp)))

		fmt.Printf("v0=%s", string(hash.Sum(nil)))
		req.Header.Add("X-Slack-Signature", "v0="+string(hash.Sum(nil)))
		req.Header.Add("X-Slack-Request-Timestamp", stimestamp)
		// add a valid body to the request
		req.Body = io.NopCloser(strings.NewReader(`{"type":"valid"}`))

		// Create a new HTTP response recorder
		res := httptest.NewRecorder()

		// Call the EventWebhook method
		s.EventWebhook(res, req)

		// Check the response status code
		assert.Equal(t, http.StatusUnauthorized, res.Code)
	})
}
