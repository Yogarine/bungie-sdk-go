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

type DestinyHistoricalStatsDestinyHistoricalStatsByPeriod struct {
	AllTime map[string]DestinyHistoricalStatsDestinyHistoricalStatsValue `json:"allTime,omitempty"`
	AllTimeTier1 map[string]DestinyHistoricalStatsDestinyHistoricalStatsValue `json:"allTimeTier1,omitempty"`
	AllTimeTier2 map[string]DestinyHistoricalStatsDestinyHistoricalStatsValue `json:"allTimeTier2,omitempty"`
	AllTimeTier3 map[string]DestinyHistoricalStatsDestinyHistoricalStatsValue `json:"allTimeTier3,omitempty"`
	Daily []DestinyHistoricalStatsDestinyHistoricalStatsPeriodGroup `json:"daily,omitempty"`
	Monthly []DestinyHistoricalStatsDestinyHistoricalStatsPeriodGroup `json:"monthly,omitempty"`
}
