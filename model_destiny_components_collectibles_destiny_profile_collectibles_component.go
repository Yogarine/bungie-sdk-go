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

type DestinyComponentsCollectiblesDestinyProfileCollectiblesComponent struct {
	// The list of collectibles determined by the game as having been \"recently\" acquired.
	RecentCollectibleHashes []int32 `json:"recentCollectibleHashes,omitempty"`
	// The list of collectibles determined by the game as having been \"recently\" acquired.  The game client itself actually controls this data, so I personally question whether anyone will get much use out of this: because we can't edit this value through the API. But in case anyone finds it useful, here it is.
	NewnessFlaggedCollectibleHashes []int32 `json:"newnessFlaggedCollectibleHashes,omitempty"`
	Collectibles map[string]DestinyComponentsCollectiblesDestinyCollectibleComponent `json:"collectibles,omitempty"`
}
