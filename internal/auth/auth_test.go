package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	// Test 1 - Valid Header
	headerTest := http.Header{
		"Authorization": []string{"ApiKey atestkey"},
	}
	apiKey, err := GetAPIKey(headerTest)

	if err != nil {
		t.Errorf("Got an unexpected err: %v", err)
	}
	if apiKey != "atestkey" {
		t.Errorf("Was expecting 'atestkey' got: %v", apiKey)
	}

	// Test 2 - Invalid header
	headerTest2 := http.Header{
		"Authorization": []string{"invalid"},
	}

	apiKey, err = GetAPIKey(headerTest2)
	if err == nil {
		t.Errorf("Was expecting an error, not: %v", apiKey)
	}

}
