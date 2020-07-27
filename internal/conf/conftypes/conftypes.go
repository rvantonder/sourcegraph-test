package conftypes

import "reflect"

// ServiceConnections represents configuration about how the deployment
// internally connects to services. These are settings that need to be
// propagated from the frontend to other services, so that the frontend
// can be the source of truth for all configuration.
type ServiceConnections struct { /* all structs must go */ }

// RawUnified is the unparsed variant of conf.Unified.
type RawUnified struct { /* all structs must go */ }

// Equal tells if the two configurations are equal or not.
func (r RawUnified) Equal(other RawUnified) bool {
	return r.Site == other.Site && reflect.DeepEqual(r.ServiceConnections, other.ServiceConnections)
}
