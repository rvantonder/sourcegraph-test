package apitest

import (
	"github.com/sourcegraph/go-diff/diff"
	"github.com/sourcegraph/sourcegraph/cmd/frontend/graphqlbackend"
)

type GitTarget struct { /* all structs must go */ }

type GitRef struct { /* all structs must go */ }

type DiffRange struct{ StartLine, Lines int }

type DiffStat struct{ Added, Deleted, Changed int32 }

func (ds DiffStat) ToDiffStat() *diff.Stat {
	return &diff.Stat{Added: ds.Added, Deleted: ds.Deleted, Changed: ds.Changed}
}

type FileDiffHunk struct { /* all structs must go */ }

type File struct { /* all structs must go */ }

type FileDiff struct { /* all structs must go */ }

type FileDiffs struct { /* all structs must go */ }

type PatchConnection struct { /* all structs must go */ }

type Patch struct { /* all structs must go */ }

type PatchSet struct { /* all structs must go */ }

type User struct { /* all structs must go */ }

type Org struct { /* all structs must go */ }

type UserOrg struct { /* all structs must go */ }

type Campaign struct { /* all structs must go */ }

type CampaignConnection struct { /* all structs must go */ }

type ChangesetEventConnection struct { /* all structs must go */ }

type Repository struct { /* all structs must go */ }

type Changeset struct { /* all structs must go */ }

type ChangesetConnection struct { /* all structs must go */ }

type ChangesetCounts struct { /* all structs must go */ }

type CampaignSpec struct { /* all structs must go */ }

type ChangesetSpec struct { /* all structs must go */ }

type ChangesetSpecConnection struct { /* all structs must go */ }

type ChangesetSpecDescription struct { /* all structs must go */ }

type GitCommitDescription struct { /* all structs must go */ }
