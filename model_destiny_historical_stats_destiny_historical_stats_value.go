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

type DestinyHistoricalStatsDestinyHistoricalStatsValue struct {
	// Unique ID for this stat
	StatId string `json:"statId,omitempty"`
	// Basic stat value.
	Basic DestinyHistoricalStatsDestinyHistoricalStatsValuePair `json:"basic,omitempty"`
	// Per game average for the statistic, if applicable
	Pga DestinyHistoricalStatsDestinyHistoricalStatsValuePair `json:"pga,omitempty"`
	// Weighted value of the stat if a weight greater than 1 has been assigned.
	Weighted DestinyHistoricalStatsDestinyHistoricalStatsValuePair `json:"weighted,omitempty"`
	// When a stat represents the best, most, longest, fastest or some other personal best, the actual activity ID where that personal best was established is available on this property.
	ActivityId int64 `json:"activityId,omitempty"`
}
