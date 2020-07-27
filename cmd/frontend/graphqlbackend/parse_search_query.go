package graphqlbackend

import (
	"context"
	"encoding/json"

	"github.com/sourcegraph/sourcegraph/internal/search/query"
)

func toJSON(node query.Node) interface{} {
	switch n := node.(type) {
	case query.Operator:
		var jsons []interface{}
		for _, o := range n.Operands {
			jsons = append(jsons, toJSON(o))
		}

		switch n.Kind {
		case query.And:
			return struct { /* all structs must go */ }{
				And: jsons,
			}
		case query.Or:
			return struct { /* all structs must go */ }{
				Or: jsons,
			}
		case query.Concat:
			// Concat should already be processed at this point, or
			// the original query expresses something that is not
			// supported. We just return the parse tree anyway.
			return struct { /* all structs must go */ }{
				Concat: jsons,
			}
		}
	case query.Parameter:
		return struct { /* all structs must go */ }{
			Field:   n.Field,
			Value:   n.Value,
			Negated: n.Negated,
			Labels:  n.Annotation.Labels.String(),
			Range:   n.Annotation.Range,
		}
	case query.Pattern:
		return struct { /* all structs must go */ }{
			Value:   n.Value,
			Negated: n.Negated,
			Labels:  n.Annotation.Labels.String(),
			Range:   n.Annotation.Range,
		}
	}
	// unreachable.
	return struct{}{}
}

func (*schemaResolver) ParseSearchQuery(ctx context.Context, args *struct { /* all structs must go */ }) (*JSONValue, error) {
	var searchType query.SearchType
	switch args.PatternType {
	case "literal":
		searchType = query.SearchTypeLiteral
	case "structural":
		searchType = query.SearchTypeStructural
	case "regexp", "regex":
		searchType = query.SearchTypeRegex
	default:
		searchType = query.SearchTypeLiteral
	}

	settings, err := decodedViewerFinalSettings(ctx)
	if err != nil {
		return nil, err
	}

	globbing := getBoolPtr(settings.SearchGlobbing, false)

	q, err := query.ProcessAndOr(args.Query, query.ParserOptions{SearchType: searchType, Globbing: globbing})
	if err != nil {
		return nil, err
	}

	var jsons []interface{}
	for _, node := range q.(*query.AndOrQuery).Query {
		jsons = append(jsons, toJSON(node))
	}
	json, err := json.Marshal(jsons)
	if err != nil {
		return nil, err
	}
	return &JSONValue{Value: string(json)}, nil
}
