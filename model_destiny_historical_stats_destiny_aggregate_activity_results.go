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

type DestinyHistoricalStatsDestinyAggregateActivityResults struct {
	// List of all activities the player has participated in.
	Activities []DestinyHistoricalStatsDestinyAggregateActivityStats `json:"activities,omitempty"`
}
