/*
 * Bungie.Net API
 *
 * These endpoints constitute the functionality exposed by Bungie.net, both for more traditional website functionality and for connectivity to Bungie video games and their related functionality.
 *
 * API version: 2.3.2
 * Contact: support@bungie.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package bungie
import (
	"time"
)

type DestinyHistoricalStatsDestinyPostGameCarnageReportData struct {
	// Date and time for the activity.
	Period time.Time `json:"period,omitempty"`
	// If this activity has \"phases\", this is the phase at which the activity was started.
	StartingPhaseIndex int32 `json:"startingPhaseIndex,omitempty"`
	// Details about the activity.
	ActivityDetails DestinyHistoricalStatsDestinyHistoricalStatsActivity `json:"activityDetails,omitempty"`
	// Collection of players and their data for this activity.
	Entries []DestinyHistoricalStatsDestinyPostGameCarnageReportEntry `json:"entries,omitempty"`
	// Collection of stats for the player in this activity.
	Teams []DestinyHistoricalStatsDestinyPostGameCarnageReportTeamEntry `json:"teams,omitempty"`
}