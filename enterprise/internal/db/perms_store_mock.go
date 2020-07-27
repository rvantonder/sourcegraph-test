package db

import (
	"context"

	"github.com/sourcegraph/sourcegraph/cmd/frontend/authz"
	"github.com/sourcegraph/sourcegraph/internal/extsvc"
)

type MockPerms struct { /* all structs must go */ }
