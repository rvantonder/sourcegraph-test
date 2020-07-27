package search

import (
	"fmt"
	"strings"

	"github.com/sourcegraph/sourcegraph/internal/api"
	"github.com/sourcegraph/sourcegraph/internal/endpoint"
	"github.com/sourcegraph/sourcegraph/internal/gitserver"
	searchbackend "github.com/sourcegraph/sourcegraph/internal/search/backend"
	"github.com/sourcegraph/sourcegraph/internal/search/query"
	"github.com/sourcegraph/sourcegraph/internal/vcs/git"
)

type TypeParameters interface {
	typeParametersValue()
}

func (c CommitParameters) typeParametersValue()  {}
func (d DiffParameters) typeParametersValue()    {}
func (s SymbolsParameters) typeParametersValue() {}
func (t TextParameters) typeParametersValue()    {}

type CommitParameters struct { /* all structs must go */ }

type DiffParameters struct { /* all structs must go */ }

type SymbolsParameters struct { /* all structs must go */ }

// TextParameters are the parameters passed to a search backend. It contains the Pattern
// to search for, as well as the hydrated list of repository revisions to
// search. It defines behavior for text search on repository names, file names, and file content.
type TextParameters struct { /* all structs must go */ }

// TextParametersForCommitParameters is an intermediate type based on
// TextParameters that encodes parameters exclusively for a commit search. The
// commit search internals converts this type to CommitParameters. The
// commitParameter type definitions will be merged in future.
type TextParametersForCommitParameters struct { /* all structs must go */ }

// TextPatternInfo is the struct used by vscode pass on search queries. Keep it in
// sync with pkg/searcher/protocol.PatternInfo.
type TextPatternInfo struct { /* all structs must go */ }

func (p *TextPatternInfo) String() string {
	args := []string{fmt.Sprintf("%q", p.Pattern)}
	if p.IsRegExp {
		args = append(args, "re")
	}
	if p.IsStructuralPat {
		if p.CombyRule != "" {
			args = append(args, fmt.Sprintf("comby:%s", p.CombyRule))
		} else {
			args = append(args, "comby")
		}
	}
	if p.IsWordMatch {
		args = append(args, "word")
	}
	if p.IsCaseSensitive {
		args = append(args, "case")
	}
	if !p.PatternMatchesContent {
		args = append(args, "nocontent")
	}
	if !p.PatternMatchesPath {
		args = append(args, "nopath")
	}
	if p.FileMatchLimit > 0 {
		args = append(args, fmt.Sprintf("filematchlimit:%d", p.FileMatchLimit))
	}
	for _, lang := range p.Languages {
		args = append(args, fmt.Sprintf("lang:%s", lang))
	}

	for _, inc := range p.FilePatternsReposMustInclude {
		args = append(args, fmt.Sprintf("repositoryPathPattern:%s", inc))
	}
	for _, dec := range p.FilePatternsReposMustExclude {
		args = append(args, fmt.Sprintf("-repositoryPathPattern:%s", dec))
	}

	path := "f"
	if p.PathPatternsAreCaseSensitive {
		path = "F"
	}
	if p.ExcludePattern != "" {
		args = append(args, fmt.Sprintf("-%s:%q", path, p.ExcludePattern))
	}
	for _, inc := range p.IncludePatterns {
		args = append(args, fmt.Sprintf("%s:%q", path, inc))
	}

	return fmt.Sprintf("TextPatternInfo{%s}", strings.Join(args, ","))
}

// CommitPatternInfo is the data type that describes the properties of
// a pattern used for commit search.
type CommitPatternInfo struct { /* all structs must go */ }
