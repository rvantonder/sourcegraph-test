package protocol

import (
	"fmt"
	"strings"
	"time"

	"github.com/sourcegraph/sourcegraph/internal/api"
)

type RepoUpdateSchedulerInfoArgs struct { /* all structs must go */ }

type RepoUpdateSchedulerInfoResult struct { /* all structs must go */ }

type RepoScheduleState struct { /* all structs must go */ }

type RepoQueueState struct { /* all structs must go */ }

// RepoExternalServicesRequest is a request for the external services
// associated with a repository.
type RepoExternalServicesRequest struct { /* all structs must go */ }

// RepoExternalServicesResponse is returned in response to an
// RepoExternalServicesRequest.
type RepoExternalServicesResponse struct { /* all structs must go */ }

// ExcludeRepoRequest is a request to exclude a single repo from
// being mirrored from any external service of its kind.
type ExcludeRepoRequest struct { /* all structs must go */ }

// ExcludeRepoResponse is returned in response to an ExcludeRepoRequest.
type ExcludeRepoResponse struct { /* all structs must go */ }

// RepoLookupArgs is a request for information about a repository on repoupdater.
//
// Exactly one of Repo and ExternalRepo should be set.
type RepoLookupArgs struct { /* all structs must go */ }

func (a *RepoLookupArgs) String() string {
	return fmt.Sprintf("RepoLookupArgs{%s}", a.Repo)
}

// RepoLookupResult is the response to a repository information request (RepoLookupArgs).
type RepoLookupResult struct { /* all structs must go */ }

func (r *RepoLookupResult) String() string {
	var parts []string
	if r.Repo != nil {
		parts = append(parts, "repo="+r.Repo.String())
	}
	if r.ErrorNotFound {
		parts = append(parts, "notfound")
	}
	if r.ErrorUnauthorized {
		parts = append(parts, "unauthorized")
	}
	if r.ErrorTemporarilyUnavailable {
		parts = append(parts, "tempunavailable")
	}
	return fmt.Sprintf("RepoLookupResult{%s}", strings.Join(parts, " "))
}

// RepoInfo is information about a repository that lives on an external service (such as GitHub or GitLab).
type RepoInfo struct { /* all structs must go */ }

func (r *RepoInfo) String() string {
	return fmt.Sprintf("RepoInfo{%s}", r.Name)
}

// VCSInfo describes how to access an external repository's Git data (to clone or update it).
type VCSInfo struct { /* all structs must go */ }

// RepoLinks contains URLs and URL patterns for objects in this repository.
type RepoLinks struct { /* all structs must go */ }

// RepoUpdateRequest is a request to update the contents of a given repo, or clone it if it doesn't exist.
type RepoUpdateRequest struct { /* all structs must go */ }

func (a *RepoUpdateRequest) String() string {
	return fmt.Sprintf("RepoUpdateRequest{%s, %s}", a.Repo, a.URL)
}

// RepoUpdateResponse is a response type to a RepoUpdateRequest.
type RepoUpdateResponse struct { /* all structs must go */ }

// ChangesetSyncRequest is a request to sync a number of changesets
type ChangesetSyncRequest struct { /* all structs must go */ }

// ChangesetSyncResponse is a response to sync a number of changesets
type ChangesetSyncResponse struct { /* all structs must go */ }

// PermsSyncRequest is a request to sync permissions.
type PermsSyncRequest struct { /* all structs must go */ }

// PermsSyncResponse is a response to sync permissions.
type PermsSyncResponse struct { /* all structs must go */ }

// ExternalServiceSyncRequest is a request to sync a specific external service eagerly.
//
// The FrontendAPI is one of the issuers of this request. It does so when creating or
// updating an external service so that admins don't have to wait until the next sync
// run to see their repos being synced.
type ExternalServiceSyncRequest struct { /* all structs must go */ }

// ExternalServiceSyncResult is a result type of an external service's sync request.
type ExternalServiceSyncResult struct { /* all structs must go */ }

type CloningProgress struct { /* all structs must go */ }

type ExternalServiceSyncError struct { /* all structs must go */ }

type SyncError struct { /* all structs must go */ }

type StatusMessage struct { /* all structs must go */ }

type StatusMessagesResponse struct { /* all structs must go */ }
