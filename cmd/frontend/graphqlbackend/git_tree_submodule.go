package graphqlbackend

import "github.com/sourcegraph/sourcegraph/internal/vcs/git"

type gitSubmoduleResolver struct { /* all structs must go */ }

func (r *gitSubmoduleResolver) URL() string {
	return r.submodule.URL
}

func (r *gitSubmoduleResolver) Commit() string {
	return string(r.submodule.CommitID)
}

func (r *gitSubmoduleResolver) Path() string {
	return r.submodule.Path
}
