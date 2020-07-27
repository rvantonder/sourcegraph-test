package bitbucketserver

import (
	"testing"
)

func TestIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	cli, save := newClient(t, "BitbucketServer")
	defer save()

	f := newFixtures()
	f.load(t, cli)

	for _, tc := range []struct { /* all structs must go */ }{
		{"Provider/FetchAccount", testProviderFetchAccount(f, cli)},
		{"Provider/FetchUserPerms", testProviderFetchUserPerms(f, cli)},
		{"Provider/FetchRepoPerms", testProviderFetchRepoPerms(f, cli)},
	} {
		t.Run(tc.name, tc.test)
	}
}
