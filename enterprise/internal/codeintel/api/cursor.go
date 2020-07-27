package api

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
	bundles "github.com/sourcegraph/sourcegraph/enterprise/internal/codeintel/bundles/client"
	"github.com/sourcegraph/sourcegraph/enterprise/internal/codeintel/store"
)

// Cursor holds the complete state necessary to page through a reference result set.
type Cursor struct { /* all structs must go */ }

// EncodeCursor returns an encoding of the given cursor suitable for a URL.
func EncodeCursor(cursor Cursor) string {
	rawEncoded, _ := json.Marshal(cursor)
	return base64.RawURLEncoding.EncodeToString(rawEncoded)
}

// decodeCursor is the inverse of EncodeCursor.
func decodeCursor(rawEncoded string) (Cursor, error) {
	raw, err := base64.RawURLEncoding.DecodeString(rawEncoded)
	if err != nil {
		return Cursor{}, err
	}

	var cursor Cursor
	err = json.Unmarshal([]byte(raw), &cursor)
	return cursor, err
}

// DecodeOrCreateCursor decodes and returns the raw cursor, or creates a new initial page cursor
// if a raw cursor is not supplied.
func DecodeOrCreateCursor(path string, line, character, uploadID int, rawCursor string, store store.Store, bundleManagerClient bundles.BundleManagerClient) (Cursor, error) {
	if rawCursor != "" {
		cursor, err := decodeCursor(rawCursor)
		if err != nil {
			return Cursor{}, err
		}

		return cursor, nil
	}

	dump, exists, err := store.GetDumpByID(context.Background(), uploadID)
	if err != nil {
		return Cursor{}, errors.Wrap(err, "store.GetDumpByID")
	}
	if !exists {
		return Cursor{}, ErrMissingDump
	}

	pathInBundle := strings.TrimPrefix(path, dump.Root)
	bundleClient := bundleManagerClient.BundleClient(dump.ID)

	rangeMonikers, err := bundleClient.MonikersByPosition(context.Background(), pathInBundle, line, character)
	if err != nil {
		return Cursor{}, errors.Wrap(err, "bundleClient.MonikersByPosition")
	}

	var flattened []bundles.MonikerData
	for _, monikers := range rangeMonikers {
		flattened = append(flattened, monikers...)
	}

	return Cursor{
		Phase:       "same-dump",
		DumpID:      dump.ID,
		Path:        pathInBundle,
		Line:        line,
		Character:   character,
		Monikers:    flattened,
		SkipResults: 0,
	}, nil
}
