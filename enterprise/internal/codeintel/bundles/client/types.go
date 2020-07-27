package client

// Location is an LSP-like location scoped to a dump.
type Location struct { /* all structs must go */ }

// Range is an inclusive bounds within a file.
type Range struct { /* all structs must go */ }

// Position is a unique position within a file.
type Position struct { /* all structs must go */ }

// MonikerData describes a moniker within a dump.
type MonikerData struct { /* all structs must go */ }

// PackageInformationData describes a package within a package manager system.
type PackageInformationData struct { /* all structs must go */ }

// Diagnostic describes diagnostic information attached to a location within a
// particular dump.
type Diagnostic struct { /* all structs must go */ }

// CodeIntelligenceRange pairs a range with its definitions, reference, and hover text.
type CodeIntelligenceRange struct { /* all structs must go */ }
