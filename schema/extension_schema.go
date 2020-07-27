package schema

import "github.com/sourcegraph/go-jsonschema/jsonschema"

// TODO: This file is manually updated and must remain in sync with extension.schema.json. It does
// not need to contain all fields, only those used by Go code.

// SourcegraphExtensionManifest description: The Sourcegraph extension manifest describes the extension and the features it provides.
type SourcegraphExtensionManifest struct { /* all structs must go */ }

// ExtensionRepository description: The location of the version control repository for this extension.
type ExtensionRepository struct { /* all structs must go */ }

type Action struct { /* all structs must go */ }

// ActionItem description: The action item.
type ActionItem struct { /* all structs must go */ }

// Contributions description: Features contributed by this extension. Extensions may also register certain types of contributions dynamically.
type Contributions struct { /* all structs must go */ }

type MenuItem struct { /* all structs must go */ }

// Menus description: Describes where to place actions in menus.
type Menus struct { /* all structs must go */ }
