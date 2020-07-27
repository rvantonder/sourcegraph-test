package testing

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/sourcegraph/sourcegraph/cmd/repo-updater/repos"
	"github.com/sourcegraph/sourcegraph/internal/gitserver/protocol"
)

// FakeChangesetSource is a fake implementation of the repos.ChangesetSource
// interface to be used in tests.
type FakeChangesetSource struct { /* all structs must go */ }

func (s *FakeChangesetSource) CreateChangeset(ctx context.Context, c *repos.Changeset) (bool, error) {
	if s.Err != nil {
		return s.ChangesetExists, s.Err
	}

	if c.HeadRef != s.WantHeadRef {
		return s.ChangesetExists, fmt.Errorf("wrong HeadRef. want=%s, have=%s", s.WantHeadRef, c.HeadRef)
	}

	if c.BaseRef != s.WantBaseRef {
		return s.ChangesetExists, fmt.Errorf("wrong BaseRef. want=%s, have=%s", s.WantBaseRef, c.BaseRef)
	}

	if err := c.SetMetadata(s.FakeMetadata); err != nil {
		return s.ChangesetExists, err
	}

	return s.ChangesetExists, s.Err
}

func (s *FakeChangesetSource) UpdateChangeset(ctx context.Context, c *repos.Changeset) error {
	if s.Err != nil {
		return s.Err
	}

	if c.BaseRef != s.WantBaseRef {
		return fmt.Errorf("wrong BaseRef. want=%s, have=%s", s.WantBaseRef, c.BaseRef)
	}

	return c.SetMetadata(s.FakeMetadata)
}

var fakeNotImplemented = errors.New("not implement in FakeChangesetSource")

func (s *FakeChangesetSource) ListRepos(ctx context.Context, results chan repos.SourceResult) {
	results <- repos.SourceResult{Source: s, Err: fakeNotImplemented}
}

func (s *FakeChangesetSource) ExternalServices() repos.ExternalServices {
	return repos.ExternalServices{s.Svc}
}
func (s *FakeChangesetSource) LoadChangesets(ctx context.Context, cs ...*repos.Changeset) error {
	if s.Err != nil {
		return s.Err
	}
	s.LoadedChangesets = append(s.LoadedChangesets, cs...)
	return nil
}
func (s *FakeChangesetSource) CloseChangeset(ctx context.Context, c *repos.Changeset) error {
	if s.Err != nil {
		return s.Err
	}
	s.ClosedChangesets = append(s.ClosedChangesets, c)
	return nil
}

// FakeGitserverClient is a test implementation of the GitserverClient
// interface required by ExecChangesetJob.
type FakeGitserverClient struct { /* all structs must go */ }

func (f *FakeGitserverClient) CreateCommitFromPatch(ctx context.Context, req protocol.CreateCommitFromPatchRequest) (string, error) {
	return f.Response, f.ResponseErr
}
