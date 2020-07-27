package comby

type Input interface {
	Value()
}

type ZipPath string
type DirPath string

func (z ZipPath) Value() {}
func (d DirPath) Value() {}

type Args struct { /* all structs must go */ }

// Location is the location in a file
type Location struct { /* all structs must go */ }

// Range is a range of start location to end location
type Range struct { /* all structs must go */ }

// Match represents a range of matched characters and the matched content
type Match struct { /* all structs must go */ }

// FileMatch represents all the matches in a single file
type FileMatch struct { /* all structs must go */ }

// FileDiff represents a diff for a file
type FileDiff struct { /* all structs must go */ }
