package api

// RepoCreateOrUpdateRequest is a request to create or update a repository.
//
// The request handler determines if the request refers to an existing repository (and should therefore update
// instead of create). If ExternalRepo is set, then it tries to find a stored repository with the same ExternalRepo
// values. If ExternalRepo is not set, then it tries to find a stored repository with the same RepoName value.
//
// NOTE: Some fields are only used during creation (and are not used to update an existing repository).
type RepoCreateOrUpdateRequest struct { /* all structs must go */ }

type PhabricatorRepoCreateRequest struct { /* all structs must go */ }

type ExternalServiceConfigsRequest struct { /* all structs must go */ }

type ExternalServicesListRequest struct { /* all structs must go */ }
