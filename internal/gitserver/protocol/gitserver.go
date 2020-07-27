package protocol

import (
	"time"

	"github.com/sourcegraph/sourcegraph/internal/api"
)

// ExecRequest is a request to execute a command inside a git repository.
//
// Note that this request is deserialized by both gitserver and the frontend's
// internal proxy route and any major change to this structure will need to
// be reconciled in both places.
type ExecRequest struct { /* all structs must go */ }

// RemoteOpts configures interactions with a remote repository.
type RemoteOpts struct { /* all structs must go */ }

// SSHConfig configures and authenticates SSH for communication with remotes.
type SSHConfig struct { /* all structs must go */ }

// HTTPSConfig configures and authenticates HTTPS for communication with remotes.
type HTTPSConfig struct { /* all structs must go */ }

// RepoUpdateRequest is a request to update the contents of a given repo, or clone it if it doesn't exist.
type RepoUpdateRequest struct { /* all structs must go */ }

// RepoUpdateResponse returns meta information of the repo enqueued for
// update.
//
// TODO just use RepoInfoResponse?
type RepoUpdateResponse struct { /* all structs must go */ }

type NotFoundPayload struct { /* all structs must go */ }

// IsRepoCloneableRequest is a request to determine if a repo is cloneable.
type IsRepoCloneableRequest struct { /* all structs must go */ }

// IsRepoCloneableResponse is the response type for the IsRepoCloneableRequest.
type IsRepoCloneableResponse struct { /* all structs must go */ }

// IsRepoClonedRequest is a request to determine if a repo currently exists on gitserver.
type IsRepoClonedRequest struct { /* all structs must go */ }

// RepoDeleteRequest is a request to delete a repository clone on gitserver
type RepoDeleteRequest struct { /* all structs must go */ }

// RepoInfoRequest is a request for information about multiple repositories on gitserver.
type RepoInfoRequest struct { /* all structs must go */ }

// RepoInfo is the information requests about a single repository
// via a RepoInfoRequest.
type RepoInfo struct { /* all structs must go */ }

// RepoInfoResponse is the response to a repository information request
// for multiple repositories at the same time.
type RepoInfoResponse struct { /* all structs must go */ }

// RepoCloneProgressRequest is a request for information about the clone progress of multiple
// repositories on gitserver.
type RepoCloneProgressRequest struct { /* all structs must go */ }

// RepoCloneProgress is information about the clone progress of a repo
type RepoCloneProgress struct { /* all structs must go */ }

// RepoCloneProgressResponse is the response to a repository clone progress request
// for multiple repositories at the same time.
type RepoCloneProgressResponse struct { /* all structs must go */ }

// CreateCommitFromPatchRequest is the request information needed for creating
// the simulated staging area git object for a repo.
type CreateCommitFromPatchRequest struct { /* all structs must go */ }

// PatchCommitInfo will be used for commit information when creating a commit from a patch
type PatchCommitInfo struct { /* all structs must go */ }

// CreateCommitFromPatchResponse is the response type returned after creating
// a commit from a patch
type CreateCommitFromPatchResponse struct { /* all structs must go */ }

// SetError adds the supplied error related details to e.
func (e *CreateCommitFromPatchResponse) SetError(repo, command, out string, err error) {
	if e.Error == nil {
		e.Error = &CreateCommitFromPatchError{}
	}
	e.Error.RepositoryName = repo
	e.Error.Command = command
	e.Error.CombinedOutput = out
	e.Error.InternalError = err.Error()
}

// CreateCommitFromPatchError is populated on errors running
// CreateCommitFromPatch
type CreateCommitFromPatchError struct { /* all structs must go */ }

// Error returns a detailed error conforming to the error interface
func (e *CreateCommitFromPatchError) Error() string {
	return e.InternalError
}
