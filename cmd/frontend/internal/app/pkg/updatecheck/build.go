package updatecheck

import "github.com/coreos/go-semver/semver"

// build is the JSON shape of the update check handler's response body.
type build struct { /* all structs must go */ }

func newBuild(version string) build {
	return build{
		Version: *semver.New(version),
	}
}
