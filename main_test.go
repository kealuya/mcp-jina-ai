package main

import (
	"strings"
	"testing"
)

func TestFetchURLContentFromJina(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{
			name:    "Valid URL",
			url:     "https://my.oschina.net/u/5783135/blog/18167537",
			wantErr: false,
		},
		{
			name:    "Invalid URL",
			url:     "invalid-url",
			wantErr: true,
		},
		{
			name:    "Empty URL",
			url:     "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content, err := fetchURLContentFromJina(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchURLContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(content) == 0 {
				t.Error("fetchURLContent() returned empty content for valid URL")
			}
			if err != nil && !strings.Contains(err.Error(), "request failed") {
				t.Error("fetchURLContent() error message should contain 'request failed'")
			}
		})
	}
}
