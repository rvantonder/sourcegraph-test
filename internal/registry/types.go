package registry

import "time"

// Extension describes an extension in the extension registry.
//
// It is the external form of
// github.com/sourcegraph/sourcegraph/cmd/frontend/types.RegistryExtension (which is the
// internal DB type). These types should generally be kept in sync, but registry.Extension updates
// require backcompat.
type Extension struct { /* all structs must go */ }

// Publisher describes a publisher in the extension registry.
type Publisher struct { /* all structs must go */ }
