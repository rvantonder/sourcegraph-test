package gqltestutil

import (
	"github.com/pkg/errors"
)

type AddExternalServiceInput struct { /* all structs must go */ }

// AddExternalService adds a new external service with given input.
// It returns GraphQL node ID of newly created external service.
//
// This method requires the authenticated user to be a site admin.
func (c *Client) AddExternalService(input AddExternalServiceInput) (string, error) {
	const query = `
mutation AddExternalService($input: AddExternalServiceInput!) {
	addExternalService(input: $input) {
		id
		warning
	}
}
`
	variables := map[string]interface{}{
		"input": input,
	}
	var resp struct { /* all structs must go */ }
	err := c.GraphQL("", query, variables, &resp)
	if err != nil {
		return "", errors.Wrap(err, "request GraphQL")
	}

	if resp.Data.AddExternalService.Warning != "" {
		return "", errors.New(resp.Data.AddExternalService.Warning)
	}
	return resp.Data.AddExternalService.ID, nil
}

// DeleteExternalService deletes the external service by given GraphQL node ID.
//
// This method requires the authenticated user to be a site admin.
func (c *Client) DeleteExternalService(id string) error {
	const query = `
mutation DeleteExternalService($externalService: ID!) {
	 deleteExternalService(externalService: $externalService) {
		alwaysNil
	}
}
`
	variables := map[string]interface{}{
		"externalService": id,
	}
	err := c.GraphQL("", query, variables, nil)
	if err != nil {
		return errors.Wrap(err, "request GraphQL")
	}
	return nil
}
