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

type DestinyHistoricalStatsDestinyHistoricalStatsAccountResult struct {
	MergedDeletedCharacters DestinyHistoricalStatsDestinyHistoricalStatsWithMerged `json:"mergedDeletedCharacters,omitempty"`
	MergedAllCharacters DestinyHistoricalStatsDestinyHistoricalStatsWithMerged `json:"mergedAllCharacters,omitempty"`
	Characters []DestinyHistoricalStatsDestinyHistoricalStatsPerCharacter `json:"characters,omitempty"`
}
