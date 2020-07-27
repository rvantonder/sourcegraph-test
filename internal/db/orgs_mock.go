package db

import (
	"context"
	"testing"

	"github.com/sourcegraph/sourcegraph/cmd/frontend/types"
)

type MockOrgs struct { /* all structs must go */ }

func (s *MockOrgs) MockGetByID_Return(t *testing.T, returns *types.Org, returnsErr error) (called *bool) {
	called = new(bool)
	s.GetByID = func(ctx context.Context, id int32) (*types.Org, error) {
		*called = true
		return returns, returnsErr
	}
	return
}
