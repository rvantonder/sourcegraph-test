package types

type ID string

// MetaData contains data describing the overall structure of a bundle.
type MetaData struct { /* all structs must go */ }

// DocumentData represents a single document within an index. The data here can answer
// definitions, references, and hover queries if the results are all contained in the
// same document.
type DocumentData struct { /* all structs must go */ }

// RangeData represents a range vertex within an index. It contains the same relevant
// edge data, which can be subsequently queried in the containing document. The data
// that was reachable via a result set has been collapsed into this object during
// conversion.
type RangeData struct { /* all structs must go */ }

// MonikerData represent a unique name (eventually) attached to a range.
type MonikerData struct { /* all structs must go */ }

// PackageInformationData indicates a globally unique namespace for a moniker.
type PackageInformationData struct { /* all structs must go */ }

// DiagnosticData carries diagnostic information attached to a range within its
// containing document.
type DiagnosticData struct { /* all structs must go */ }

// ResultChunkData represents a row of the resultChunk table. Each row is a subset
// of definition and reference result data in the index. Results are inserted into
// chunks based on the hash of their identifier, thus every chunk has a roughly
// proportional amount of data.
type ResultChunkData struct { /* all structs must go */ }

// DocumentIDRangeID is a pair of document and range identifiers.
type DocumentIDRangeID struct { /* all structs must go */ }

// Loocation represents a range within a particular document relative to its
// containing bundle.
type Location struct { /* all structs must go */ }

// MonikerLocations pairs a moniker scheme and identifier with the set of locations
// with that within a particular bundle.
type MonikerLocations struct { /* all structs must go */ }

// Package pairs a package name and the dump that provides it.
type Package struct { /* all structs must go */ }

// PackageReferences pairs a package name/version with a dump that depends on it.
type PackageReference struct { /* all structs must go */ }
