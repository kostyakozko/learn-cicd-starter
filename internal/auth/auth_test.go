package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError string
	}{
		{
			name: "valid API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey test-key-123"},
			},
			expectedKey:   "test-key-123",
			expectedError: "",
		},
		{
			name:          "missing authorization header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: "no authorization header included",
		},
		{
			name: "malformed header - wrong prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer test-key-123"},
			},
			expectedKey:   "",
			expectedError: "malformed authorization header",
		},
		{
			name: "malformed header - missing key",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectedKey:   "",
			expectedError: "malformed authorization header",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)

			if key != tt.expectedKey {
				t.Errorf("expected key %q, got %q", tt.expectedKey, key)
			}

			if tt.expectedError == "" && err != nil {
				t.Errorf("expected no error, got %v", err)
			}

			if tt.expectedError != "" && (err == nil || err.Error() != tt.expectedError) {
				t.Errorf("expected error %q, got %v", tt.expectedError, err)
			}
		})
	}
}
