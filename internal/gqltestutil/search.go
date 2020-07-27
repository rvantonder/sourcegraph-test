package gqltestutil

import (
	"github.com/pkg/errors"
)

type SearchRepositoryResult struct { /* all structs must go */ }

type SearchRepositoryResults []*SearchRepositoryResult

// Exists returns the list of missing repositories from given names that do not exist
// in search results. If all of given names are found, it returns empty list.
func (rs SearchRepositoryResults) Exists(names ...string) []string {
	set := make(map[string]struct{}, len(names))
	for _, name := range names {
		set[name] = struct{}{}
	}
	for _, r := range rs {
		delete(set, r.Name)
	}

	missing := make([]string, 0, len(set))
	for name := range set {
		missing = append(missing, name)
	}
	return missing
}

// SearchRepositories search repositories with given query.
func (c *Client) SearchRepositories(query string) (SearchRepositoryResults, error) {
	const gqlQuery = `
query Search($query: String!) {
	search(query: $query) {
		results {
			results {
				... on Repository {
					name
				}
			}
		}
	}
}
`
	variables := map[string]interface{}{
		"query": query,
	}
	var resp struct { /* all structs must go */ }
	err := c.GraphQL("", gqlQuery, variables, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "request GraphQL")
	}

	return resp.Data.Search.Results.Results, nil
}

type SearchFileResults struct { /* all structs must go */ }

// SearchFiles searches files with given query. It returns the match count and
// corresponding file matches.
func (c *Client) SearchFiles(query string) (*SearchFileResults, error) {
	const gqlQuery = `
query Search($query: String!) {
	search(query: $query) {
		results {
			matchCount
			results {
				... on FileMatch {
					file {
						name
					}
					repository {
						name
					}
					revSpec {
						... on GitRevSpecExpr {
							expr
						}
					}
				}
			}
		}
	}
}
`
	variables := map[string]interface{}{
		"query": query,
	}
	var resp struct { /* all structs must go */ }
	err := c.GraphQL("", gqlQuery, variables, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "request GraphQL")
	}

	return resp.Data.Search.Results.SearchFileResults, nil
}

type SearchCommitResults struct { /* all structs must go */ }

// SearchCommits searches commits with given query. It returns the match count and
// corresponding file matches.
func (c *Client) SearchCommits(query string) (*SearchCommitResults, error) {
	const gqlQuery = `
query Search($query: String!) {
	search(query: $query) {
		results {
			matchCount
			results {
				... on CommitSearchResult {
					url
				}
			}
		}
	}
}
`
	variables := map[string]interface{}{
		"query": query,
	}
	var resp struct { /* all structs must go */ }
	err := c.GraphQL("", gqlQuery, variables, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "request GraphQL")
	}

	return resp.Data.Search.Results.SearchCommitResults, nil
}

// SearchAlert is an alert specific to searches (i.e. not site alert).
type SearchAlert struct { /* all structs must go */ }

// SearchAlert returns the alert returned by searching for given query.
// It returns both nil values if no alert raised and no error occurred.
func (c *Client) SearchAlert(query string) (*SearchAlert, error) {
	const gqlQuery = `
query Search($query: String!) {
	search(query: $query) {
		results {
			alert {
				title
				description
				proposedQueries {
					description
					query
				}
			}
		}
	}
}
`
	variables := map[string]interface{}{
		"query": query,
	}
	var resp struct { /* all structs must go */ }
	err := c.GraphQL("", gqlQuery, variables, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "request GraphQL")
	}

	return resp.Data.Search.Results.SearchAlert, nil
}

type SearchStatsResult struct { /* all structs must go */ }

// SearchStats returns statistics of given query.
func (c *Client) SearchStats(query string) (*SearchStatsResult, error) {
	const gqlQuery = `
query SearchResultsStats($query: String!) {
	search(query: $query) {
		stats {
			languages {
				name
				totalLines
			}
		}
	}
}
`
	variables := map[string]interface{}{
		"query": query,
	}
	var resp struct { /* all structs must go */ }
	err := c.GraphQL("", gqlQuery, variables, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "request GraphQL")
	}

	return resp.Data.Search.Stats, nil
}
