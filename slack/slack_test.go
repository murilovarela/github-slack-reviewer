package slack_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"encore.app/slack"
)

func TestSubscribeToEvents(t *testing.T) {
	t.Run("subscribe to events", func(t *testing.T) {
		// Create a new request to the subscribe endpoint
		requestBody := []byte(`{"challenge":"challenge","type":"url_verification"}`)
		req, err := http.NewRequest("POST", "/slack/events", bytes.NewBuffer(requestBody))
		if err != nil {
			t.Fatal(err)
		}

		// Create a response recorder to record the response
		rr := httptest.NewRecorder()

		// Serve the request to the subscribe endpoint
		handler := http.HandlerFunc(slack.SubscribeToEvents)
		handler.ServeHTTP(rr, req)

		// Check the status code of the response
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Check the response body
		expectedChallenge := "challenge"
		if rr.Body.String() != expectedChallenge {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expectedChallenge)
		}
	})
}
