// Package types defines types used by the frontend.
package types

import (
	"database/sql"
	"time"

	"github.com/sourcegraph/sourcegraph/internal/api"
	"github.com/sourcegraph/sourcegraph/internal/extsvc"
)

// RepoFields are lazy loaded data fields on a Repo (from the DB).
type RepoFields struct { /* all structs must go */ }

// Repo represents a source code repository.
type Repo struct { /* all structs must go */ }

// Repos is an utility type of a list of repos.
type Repos []*Repo

func (rs Repos) Len() int           { return len(rs) }
func (rs Repos) Less(i, j int) bool { return rs[i].ID < rs[j].ID }
func (rs Repos) Swap(i, j int)      { rs[i], rs[j] = rs[j], rs[i] }

// ExternalService is a connection to an external service.
type ExternalService struct { /* all structs must go */ }

// URN returns a unique resource identifier of this external service.
func (e *ExternalService) URN() string {
	return extsvc.URN(e.Kind, e.ID)
}

type GlobalState struct { /* all structs must go */ }

// User represents a registered user.
type User struct { /* all structs must go */ }

type Org struct { /* all structs must go */ }

type OrgMembership struct { /* all structs must go */ }

type PhabricatorRepo struct { /* all structs must go */ }

type UserUsageStatistics struct { /* all structs must go */ }

// UserUsageCounts captures the usage numbers of a user in a single day.
type UserUsageCounts struct { /* all structs must go */ }

// UserDates captures the created and deleted dates of a single user.
type UserDates struct { /* all structs must go */ }

// NOTE: DO NOT alter this struct without making a symmetric change
// to the updatecheck handler. This struct is marshalled and sent to
// BigQuery, which requires the input match its schema exactly.
type SiteUsageStatistics struct { /* all structs must go */ }

// NOTE: DO NOT alter this struct without making a symmetric change
// to the updatecheck handler. This struct is marshalled and sent to
// BigQuery, which requires the input match its schema exactly.
type SiteActivityPeriod struct { /* all structs must go */ }

// NOTE: DO NOT alter this struct without making a symmetric change
// to the updatecheck handler. This struct is marshalled and sent to
// BigQuery, which requires the input match its schema exactly.
type Stages struct { /* all structs must go */ }

// NOTE: DO NOT alter this struct without making a symmetric change
// to the updatecheck handler. This struct is marshalled and sent to
// BigQuery, which requires the input match its schema exactly.
type CampaignsUsageStatistics struct { /* all structs must go */ }

// NOTE: DO NOT alter this struct without making a symmetric change
// to the updatecheck handler. This struct is marshalled and sent to
// BigQuery, which requires the input match its schema exactly.
type CodeIntelUsageStatistics struct { /* all structs must go */ }

// NOTE: DO NOT alter this struct without making a symmetric change
// to the updatecheck handler. This struct is marshalled and sent to
// BigQuery, which requires the input match its schema exactly.
type CodeIntelUsagePeriod struct { /* all structs must go */ }

// NOTE: DO NOT alter this struct without making a symmetric change
// to the updatecheck handler. This struct is marshalled and sent to
// BigQuery, which requires the input match its schema exactly.
type CodeIntelEventCategoryStatistics struct { /* all structs must go */ }

// NOTE: DO NOT alter this struct without making a symmetric change
// to the updatecheck handler. This struct is marshalled and sent to
// BigQuery, which requires the input match its schema exactly.
type CodeIntelEventStatistics struct { /* all structs must go */ }

// NOTE: DO NOT alter this struct without making a symmetric change
// to the updatecheck handler. This struct is marshalled and sent to
// BigQuery, which requires the input match its schema exactly.
type CodeIntelEventLatencies struct { /* all structs must go */ }

// NOTE: DO NOT alter this struct without making a symmetric change
// to the updatecheck handler. This struct is marshalled and sent to
// BigQuery, which requires the input match its schema exactly.
type SearchUsageStatistics struct { /* all structs must go */ }

// NOTE: DO NOT alter this struct without making a symmetric change
// to the updatecheck handler. This struct is marshalled and sent to
// BigQuery, which requires the input match its schema exactly.
type SearchUsagePeriod struct { /* all structs must go */ }

type SearchModeUsageStatistics struct { /* all structs must go */ }

type SearchCountStatistics struct { /* all structs must go */ }

// NOTE: DO NOT alter this struct without making a symmetric change
// to the updatecheck handler. This struct is marshalled and sent to
// BigQuery, which requires the input match its schema exactly.
type SearchEventStatistics struct { /* all structs must go */ }

// NOTE: DO NOT alter this struct without making a symmetric change
// to the updatecheck handler. This struct is marshalled and sent to
// BigQuery, which requires the input match its schema exactly.
type SearchEventLatencies struct { /* all structs must go */ }

// SiteUsageSummary is an alternate view of SiteUsageStatistics which is
// calculated in the database layer.
type SiteUsageSummary struct { /* all structs must go */ }

// AggregatedEvent represents the total events, unique users, and
// latencies over the current month, week, and day for a single event.
type AggregatedEvent struct { /* all structs must go */ }

type SurveyResponse struct { /* all structs must go */ }

type Event struct { /* all structs must go */ }

// Secret represents the secrets table
type Secret struct { /* all structs must go */ }
