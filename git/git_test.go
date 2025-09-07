package git

import (
	"fmt"
	"strings"
	"testing"
)

func TestTicketNumber(t *testing.T) {
	out := TicketNumber("ne", "ne/APIT-1234")
	if out != "APIT-1234" {
		t.Fatalf("Expected APIT-1234, got " + out)
	}
}

// parseRemoteURL is a helper function to test the URL parsing logic
// without depending on shell commands
func parseRemoteURL(remoteURL string) (string, string, error) {
	remoteURL = strings.TrimSpace(remoteURL)
	
	// Handle SSH URLs (git@github.com:owner/repo.git)
	if strings.HasPrefix(remoteURL, "git@") {
		parts := strings.Split(remoteURL, ":")
		if len(parts) != 2 {
			return "", "", fmt.Errorf("invalid SSH remote URL format: %s", remoteURL)
		}
		repoPath := strings.TrimSuffix(parts[1], ".git")
		pathParts := strings.Split(repoPath, "/")
		if len(pathParts) != 2 {
			return "", "", fmt.Errorf("invalid repository path format: %s", repoPath)
		}
		return pathParts[0], pathParts[1], nil
	}
	
	// Handle HTTPS URLs (https://github.com/owner/repo.git)
	if strings.HasPrefix(remoteURL, "https://") {
		remoteURL = strings.TrimPrefix(remoteURL, "https://github.com/")
		remoteURL = strings.TrimSuffix(remoteURL, ".git")
		pathParts := strings.Split(remoteURL, "/")
		if len(pathParts) != 2 {
			return "", "", fmt.Errorf("invalid HTTPS repository path format: %s", remoteURL)
		}
		return pathParts[0], pathParts[1], nil
	}
	
	return "", "", fmt.Errorf("unsupported remote URL format: %s", remoteURL)
}

func TestParseRemoteURL_SSH(t *testing.T) {
	tests := []struct {
		name        string
		remoteURL   string
		wantOwner   string
		wantRepo    string
		wantErr     bool
	}{
		{
			name:      "SSH URL with .git suffix",
			remoteURL: "git@github.com:octocat/Hello-World.git",
			wantOwner: "octocat",
			wantRepo:  "Hello-World",
			wantErr:   false,
		},
		{
			name:      "SSH URL without .git suffix",
			remoteURL: "git@github.com:octocat/Hello-World",
			wantOwner: "octocat",
			wantRepo:  "Hello-World",
			wantErr:   false,
		},
		{
			name:      "SSH URL with organization",
			remoteURL: "git@github.com:myorg/myproject.git",
			wantOwner: "myorg",
			wantRepo:  "myproject",
			wantErr:   false,
		},
		{
			name:      "SSH URL with dashes and underscores",
			remoteURL: "git@github.com:user-name/repo_name.git",
			wantOwner: "user-name",
			wantRepo:  "repo_name",
			wantErr:   false,
		},
		{
			name:      "Invalid SSH URL - no colon",
			remoteURL: "git@github.com/octocat/Hello-World.git",
			wantErr:   true,
		},
		{
			name:      "Invalid SSH URL - too many path parts",
			remoteURL: "git@github.com:octocat/Hello/World.git",
			wantErr:   true,
		},
		{
			name:      "Invalid SSH URL - too few path parts",
			remoteURL: "git@github.com:octocat",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			owner, repo, err := parseRemoteURL(tt.remoteURL)
			
			if (err != nil) != tt.wantErr {
				t.Errorf("parseRemoteURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
			if !tt.wantErr {
				if owner != tt.wantOwner {
					t.Errorf("parseRemoteURL() owner = %v, want %v", owner, tt.wantOwner)
				}
				if repo != tt.wantRepo {
					t.Errorf("parseRemoteURL() repo = %v, want %v", repo, tt.wantRepo)
				}
			}
		})
	}
}

func TestParseRemoteURL_HTTPS(t *testing.T) {
	tests := []struct {
		name        string
		remoteURL   string
		wantOwner   string
		wantRepo    string
		wantErr     bool
	}{
		{
			name:      "HTTPS URL with .git suffix",
			remoteURL: "https://github.com/octocat/Hello-World.git",
			wantOwner: "octocat",
			wantRepo:  "Hello-World",
			wantErr:   false,
		},
		{
			name:      "HTTPS URL without .git suffix",
			remoteURL: "https://github.com/octocat/Hello-World",
			wantOwner: "octocat",
			wantRepo:  "Hello-World",
			wantErr:   false,
		},
		{
			name:      "HTTPS URL with organization",
			remoteURL: "https://github.com/myorg/myproject.git",
			wantOwner: "myorg",
			wantRepo:  "myproject",
			wantErr:   false,
		},
		{
			name:      "HTTPS URL with dashes and underscores",
			remoteURL: "https://github.com/user-name/repo_name.git",
			wantOwner: "user-name",
			wantRepo:  "repo_name",
			wantErr:   false,
		},
		{
			name:      "Invalid HTTPS URL - too many path parts",
			remoteURL: "https://github.com/octocat/Hello/World.git",
			wantErr:   true,
		},
		{
			name:      "Invalid HTTPS URL - too few path parts",
			remoteURL: "https://github.com/octocat",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			owner, repo, err := parseRemoteURL(tt.remoteURL)
			
			if (err != nil) != tt.wantErr {
				t.Errorf("parseRemoteURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
			if !tt.wantErr {
				if owner != tt.wantOwner {
					t.Errorf("parseRemoteURL() owner = %v, want %v", owner, tt.wantOwner)
				}
				if repo != tt.wantRepo {
					t.Errorf("parseRemoteURL() repo = %v, want %v", repo, tt.wantRepo)
				}
			}
		})
	}
}

func TestParseRemoteURL_UnsupportedFormats(t *testing.T) {
	tests := []struct {
		name      string
		remoteURL string
	}{
		{
			name:      "Unsupported protocol",
			remoteURL: "ftp://github.com/octocat/Hello-World.git",
		},
		{
			name:      "Local path",
			remoteURL: "/local/path/to/repo",
		},
		{
			name:      "Empty string",
			remoteURL: "",
		},
		{
			name:      "Just whitespace",
			remoteURL: "   \n\t   ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := parseRemoteURL(tt.remoteURL)
			if err == nil {
				t.Errorf("parseRemoteURL() expected error for unsupported format: %s", tt.remoteURL)
			}
		})
	}
}
