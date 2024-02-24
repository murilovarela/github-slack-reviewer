package slack_connector_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"encore.app/slack_connector"
)

func TestSubscribeToEvents(t *testing.T) {
	t.Run("SubscribeToEvents with url_verification", func(t *testing.T) {
		var requestBody = []byte(`{"type":"url_verification","challenge":"test_challenge"}`)
		var req, err = http.NewRequest("POST", "/slack/events", bytes.NewBuffer(requestBody))

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		var rr = httptest.NewRecorder()
		handler := http.HandlerFunc(slack_connector.HandleEvents)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		var expectedResponse = "test_challenge"

		if rr.Body.String() != expectedResponse {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expectedResponse)
		}
	})

	t.Run("SubscribeToEvents with unknown type", func(t *testing.T) {
		var requestBody = []byte(`{"type":"unknown_type","challenge":"test_challenge"}`)
		var req, err = http.NewRequest("POST", "/slack/events", bytes.NewBuffer(requestBody))

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		var rr = httptest.NewRecorder()
		handler := http.HandlerFunc(slack_connector.HandleEvents)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})

	t.Run("SubscribeToEvents with callback_event", func(t *testing.T) {
		var requestBody = []byte(`{"type":"callback_event","challenge":""}`)
		var req, err = http.NewRequest("POST", "/slack/events", bytes.NewBuffer(requestBody))

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		var rr = httptest.NewRecorder()
		handler := http.HandlerFunc(slack_connector.HandleEvents)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})
}
