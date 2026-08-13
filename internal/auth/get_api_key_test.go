package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeySuccess(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-secret-api-key-123")

	got, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	want := "my-secret-api-key-123"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

// func TestGetAPIKeySuccess(t *testing.T) {
// 	headers := http.Header{}
// 	headers.Set("Authorization", "ApiKey my-secret-api-key-123")

// 	got, err := GetAPIKey(headers)
// 	if err != nil {
// 		t.Fatalf("expected no error, got: %v", err)
// 	}

// 	// Deliberately breaking the expectation:
// 	want := "WRONG_KEY_TO_FAIL_CI"
// 	if got != want {
// 		t.Errorf("expected %q, got %q", want, got)
// 	}
// }

func TestGetAPIKeyErrors(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
	}{
		{
			name:    "Missing Authorization Header",
			headers: http.Header{},
		},
		{
			name: "Malformed Authorization Header - No Prefix",
			headers: http.Header{
				"Authorization": []string{"my-secret-api-key-123"},
			},
		},
		{
			name: "Malformed Authorization Header - Wrong Prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer my-secret-api-key-123"},
			},
		},
		{
			name: "Malformed Authorization Header - Missing Key",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetAPIKey(tt.headers)
			if err == nil {
				t.Errorf("expected an error for test %q, but got none", tt.name)
			}
		})
	}
}
