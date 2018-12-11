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

// The same as GroupV2ClanInfo, but includes any investment data.
type GroupsV2GroupV2ClanInfoAndInvestment struct {
	D2ClanProgressions map[string]DestinyDestinyProgression `json:"d2ClanProgressions,omitempty"`
	ClanCallsign string `json:"clanCallsign,omitempty"`
	ClanBannerData GroupsV2ClanBanner `json:"clanBannerData,omitempty"`
}
