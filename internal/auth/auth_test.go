package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		headers    http.Header
		wantAPIKey string
		wantErr    error
	}{
		{
			name:       "valid api key",
			headers:    http.Header{"Authorization": []string{"ApiKey test-key"}},
			wantAPIKey: "test-key",
			wantErr:    nil,
		},
		{
			name:       "missing auth header",
			headers:    http.Header{},
			wantAPIKey: "",
			wantErr:    ErrNoAuthHeaderIncluded,
		},
		{
			name:       "malformed header",
			headers:    http.Header{"Authorization": []string{"Bearer wrong-format"}},
			wantAPIKey: "",
			wantErr:    errors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAPIKey, gotErr := GetAPIKey(tt.headers)
			if gotErr != nil && tt.wantErr != nil && gotErr.Error() != tt.wantErr.Error() {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", gotErr, tt.wantErr)
				return
			}
			if gotAPIKey != tt.wantAPIKey {
				t.Errorf("GetAPIKey() = %v, want %v", gotAPIKey, tt.wantAPIKey)
			}
		})
	}
}
