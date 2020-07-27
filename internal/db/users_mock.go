package db

import (
	"context"
	"testing"

	"github.com/sourcegraph/sourcegraph/cmd/frontend/types"
)

type MockUsers struct { /* all structs must go */ }

func (s *MockUsers) MockGetByID_Return(t *testing.T, returns *types.User, returnsErr error) (called *bool) {
	called = new(bool)
	s.GetByID = func(ctx context.Context, id int32) (*types.User, error) {
		*called = true
		return returns, returnsErr
	}
	return

}
