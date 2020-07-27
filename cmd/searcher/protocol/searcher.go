// Package protocol contains structures used by the searcher API.
package protocol

import (
	"fmt"
	"strings"

	"github.com/sourcegraph/sourcegraph/internal/api"
	"github.com/sourcegraph/sourcegraph/internal/gitserver"
)

// Request represents a request to searcher
type Request struct { /* all structs must go */ }

// GitserverRepo returns the repository information necessary to perform gitserver requests.
func (r Request) GitserverRepo() gitserver.Repo { return gitserver.Repo{Name: r.Repo} }

// PatternInfo describes a search request on a repo. Most of the fields
// are based on PatternInfo used in vscode.
type PatternInfo struct { /* all structs must go */ }

func (p *PatternInfo) String() string {
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

	path := "glob"
	if p.PathPatternsAreRegExps {
		path = "f"
	}
	if p.PathPatternsAreCaseSensitive {
		path = "F"
	}
	if p.ExcludePattern != "" {
		args = append(args, fmt.Sprintf("-%s:%q", path, p.ExcludePattern))
	}
	for _, inc := range p.IncludePatterns {
		args = append(args, fmt.Sprintf("%s:%q", path, inc))
	}

	return fmt.Sprintf("PatternInfo{%s}", strings.Join(args, ","))
}

// Response represents the response from a Search request.
type Response struct { /* all structs must go */ }

// FileMatch is the struct used by vscode to receive search results
type FileMatch struct { /* all structs must go */ }

// LineMatch is the struct used by vscode to receive search results for a line.
type LineMatch struct { /* all structs must go */ }
