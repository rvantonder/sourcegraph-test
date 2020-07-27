// Package api contains an API client and types for cross-service communication.
package api

import (
	"fmt"
	"time"
)

// RepoID is the unique identifier for a repository.
type RepoID int32

// RepoName is the name of a repository, consisting of one or more "/"-separated path components.
//
// Previously, this was called RepoURI.
type RepoName string

// CommitID is the 40-character SHA-1 hash for a Git commit.
type CommitID string

// Repo represents a source code repository.
type Repo struct { /* all structs must go */ }

func (Repo) Fork() bool {
	// TODO(sqs): update callers
	return false
}

// ExternalRepoSpec specifies a repository on an external service (such as GitHub or GitLab).
type ExternalRepoSpec struct { /* all structs must go */ }

// Equal returns true if r is equal to s.
func (r ExternalRepoSpec) Equal(s *ExternalRepoSpec) bool {
	return r.ID == s.ID && r.ServiceType == s.ServiceType && r.ServiceID == s.ServiceID
}

// Compare returns -1 if r < s, 0 if r == s or 1 if r > s
func (r ExternalRepoSpec) Compare(s ExternalRepoSpec) int {
	if r.ServiceType != s.ServiceType {
		return cmp(r.ServiceType, s.ServiceType)
	}
	if r.ServiceID != s.ServiceID {
		return cmp(r.ServiceID, s.ServiceID)
	}
	return cmp(r.ID, s.ID)
}

func (r ExternalRepoSpec) String() string {
	return fmt.Sprintf("ExternalRepoSpec{%s %s %s}", r.ServiceID, r.ServiceType, r.ID)
}

// A SettingsSubject is something that can have settings. Exactly 1 field must be nonzero.
type SettingsSubject struct { /* all structs must go */ }

func (s SettingsSubject) String() string {
	switch {
	case s.Default:
		return "DefaultSettings"
	case s.Site:
		return "site"
	case s.Org != nil:
		return fmt.Sprintf("org %d", *s.Org)
	case s.User != nil:
		return fmt.Sprintf("user %d", *s.User)
	default:
		return "unknown settings subject"
	}
}

// Settings contains settings for a subject.
type Settings struct { /* all structs must go */ }

// ExternalService represents an complete external service record.
type ExternalService struct { /* all structs must go */ }

func cmp(a, b string) int {
	switch {
	case a < b:
		return -1
	case b < a:
		return 1
	default:
		return 0
	}
}
