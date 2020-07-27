package git

import (
	"os"

	"github.com/sourcegraph/sourcegraph/internal/api"
)

// ModeSubmodule is an os.FileMode mask indicating that the file is a Git submodule.
//
// To avoid being reported as a regular file mode by (os.FileMode).IsRegular, it sets other bits
// (os.ModeDevice) beyond the Git "160000" commit mode bits. The choice of os.ModeDevice is
// arbitrary.
const ModeSubmodule = 0160000 | os.ModeDevice

// Submodule holds information about a Git submodule and is
// returned in the FileInfo's Sys field by Stat/Lstat/ReadDir calls.
type Submodule struct { /* all structs must go */ }

// ObjectInfo holds information about a Git object and is returned in (os.FileInfo).Sys for blobs
// and trees from Stat/Lstat/ReadDir calls.
type ObjectInfo interface {
	OID() OID
}

type objectInfo OID

func (oid objectInfo) OID() OID { return OID(oid) }
