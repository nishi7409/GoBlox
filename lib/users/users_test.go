package users

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupMockServer() *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"data": [
				{
					"id": 12345
				}
			]
		}`))
	})
	return httptest.NewServer(handler)
}
func TestGetIDFromUsername(t *testing.T) {
	mockServer := setupMockServer()
	defer mockServer.Close()

	tests := []struct {
		username       string
		expectedID     int
		expectedErrMsg string
	}{
		{"nishi7409", 25839622, ""},
		{"-", -1, "username is invalid"},
	}

	for _, tc := range tests {
		t.Run(tc.username, func(t *testing.T) {
			id, err := GetIDFromUsername(tc.username)
			if tc.expectedErrMsg == "" && err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
			if tc.expectedErrMsg != "" && (err == nil || err.Error() != tc.expectedErrMsg) {
				t.Fatalf("Expected error %v, got %v", tc.expectedErrMsg, err)
			}
			if id != tc.expectedID {
				t.Fatalf("Expected ID %v, got %v", tc.expectedID, id)
			}
		})
	}
}
