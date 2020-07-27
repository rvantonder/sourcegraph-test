package types

import (
	"github.com/sourcegraph/sourcegraph/schema"
)

type CodeHostConnection interface {
	// SetURN updates the URN field of the underlying connection.
	SetURN(string)
}

var _ CodeHostConnection = (*AWSCodeCommitConnection)(nil)

type AWSCodeCommitConnection struct { /* all structs must go */ }

func (c *AWSCodeCommitConnection) SetURN(urn string) {
	c.URN = urn
}

var _ CodeHostConnection = (*BitbucketCloudConnection)(nil)

type BitbucketCloudConnection struct { /* all structs must go */ }

func (c *BitbucketCloudConnection) SetURN(urn string) {
	c.URN = urn
}

var _ CodeHostConnection = (*BitbucketServerConnection)(nil)

type BitbucketServerConnection struct { /* all structs must go */ }

func (c *BitbucketServerConnection) SetURN(urn string) {
	c.URN = urn
}

var _ CodeHostConnection = (*GitHubConnection)(nil)

type GitHubConnection struct { /* all structs must go */ }

func (c *GitHubConnection) SetURN(urn string) {
	c.URN = urn
}

var _ CodeHostConnection = (*GitLabConnection)(nil)

type GitLabConnection struct { /* all structs must go */ }

func (c *GitLabConnection) SetURN(urn string) {
	c.URN = urn
}

var _ CodeHostConnection = (*GitoliteConnection)(nil)

type GitoliteConnection struct { /* all structs must go */ }

func (c *GitoliteConnection) SetURN(urn string) {
	c.URN = urn
}

var _ CodeHostConnection = (*OtherExternalServiceConnection)(nil)

type OtherExternalServiceConnection struct { /* all structs must go */ }

func (c *OtherExternalServiceConnection) SetURN(urn string) {
	c.URN = urn
}

var _ CodeHostConnection = (*PhabricatorConnection)(nil)

type PhabricatorConnection struct { /* all structs must go */ }

func (c *PhabricatorConnection) SetURN(urn string) {
	c.URN = urn
}
