// Package protocol contains structures used by the replacer API.
package protocol

import (
	"github.com/sourcegraph/sourcegraph/internal/api"
	"github.com/sourcegraph/sourcegraph/internal/gitserver"
)

// Request represents a request to replacer
type Request struct { /* all structs must go */ }

type RewriteSpecification struct { /* all structs must go */ }

// GitserverRepo returns the repository information necessary to perform gitserver requests.
func (r Request) GitserverRepo() gitserver.Repo { return gitserver.Repo{Name: r.Repo, URL: r.URL} }
