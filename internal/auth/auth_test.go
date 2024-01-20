package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Case: No authorization header included
	headers := http.Header{}
	apiKey, err := GetAPIKey(headers)
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("Expected error ErrNoAuthHeaderIncluded, got: %v", err)
	}
	if apiKey != "" {
		t.Errorf("Expected apiKey to be empty, got: %s", apiKey)
	}

	headers = http.Header{
		"Authorization": []string{"ApiKey my-api-key"},
	}
	apiKey, err = GetAPIKey(headers)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if apiKey != "my-api-key" {
		t.Errorf("Expected apiKey to be 'my-api-key', got: %s", apiKey)
	}
}

// test
func TestGetAPIKeyWithNumbers(t *testing.T) {
	//

	got, err := GetAPIKey(http.Header{"Authorization": []string{"nothing"}})
	if err == nil {
		t.Fatalf("expected:")

	}
	want := ""
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
