package json

import (
	"encoding/json"

	"github.com/sourcegraph/sourcegraph/enterprise/internal/codeintel/bundles/types"
)

//
// The following types are used during marshalling

type SerializingTaggedValue struct { /* all structs must go */ }

type SerializingRange struct { /* all structs must go */ }

type SerializingDocument struct { /* all structs must go */ }

type SerializingResultChunk struct { /* all structs must go */ }

type SerializingLocation struct { /* all structs must go */ }

//
// The following types are used during unmarshalling

type SerializedTaggedValue struct { /* all structs must go */ }

type SerializedRange struct { /* all structs must go */ }

type SerializedDocument struct { /* all structs must go */ }

type SerializedResultChunk struct { /* all structs must go */ }

type SerializedLocation = SerializingLocation

type SerializedMoniker struct { /* all structs must go */ }

type SerializedPackageInformation struct { /* all structs must go */ }

type SerializedDocumentIDRangeID struct { /* all structs must go */ }
