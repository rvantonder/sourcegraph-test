package gitlab

import (
	"context"
	"fmt"
	"net/http"

	"github.com/peterhellberg/link"
)

type User struct { /* all structs must go */ }

type Identity struct { /* all structs must go */ }

func (c *Client) ListUsers(ctx context.Context, urlStr string) (users []*User, nextPageURL *string, err error) {
	if MockListUsers != nil {
		return MockListUsers(c, ctx, urlStr)
	}

	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, nil, err
	}
	respHeader, _, err := c.do(ctx, req, &users)
	if err != nil {
		return nil, nil, err
	}

	// Get URL to next page. See https://docs.gitlab.com/ee/api/README.html#pagination-link-header.
	if l := link.Parse(respHeader.Get("Link"))["next"]; l != nil {
		nextPageURL = &l.URI
	}

	return users, nextPageURL, nil
}

func (c *Client) GetUser(ctx context.Context, id string) (*User, error) {
	if MockGetUser != nil {
		return MockGetUser(c, ctx, id)
	}

	var urlStr string
	if id == "" {
		urlStr = "user"
	} else {
		urlStr = fmt.Sprintf("users/%s", id)
	}

	var usr User
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}
	if _, _, err := c.do(ctx, req, &usr); err != nil {
		return nil, err
	}
	return &usr, nil
}
