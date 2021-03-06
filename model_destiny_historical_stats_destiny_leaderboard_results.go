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

type DestinyHistoricalStatsDestinyLeaderboardResults struct {
	// Indicate the membership ID of the account that is the focal point of the provided leaderboards.
	FocusMembershipId int64 `json:"focusMembershipId,omitempty"`
	// Indicate the character ID of the character that is the focal point of the provided leaderboards. May be null, in which case any character from the focus membership can appear in the provided leaderboards.
	FocusCharacterId int64 `json:"focusCharacterId,omitempty"`
}
