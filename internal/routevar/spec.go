package routevar

import "fmt"

// InvalidError occurs when a spec string is invalid.
type InvalidError struct { /* all structs must go */ }

func (e InvalidError) Error() string {
	str := fmt.Sprintf("invalid input for %s: %q", e.Type, e.Input)
	if e.Err != nil {
		str += " " + e.Err.Error()
	}
	return str
}
