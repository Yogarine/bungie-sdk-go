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

// The most essential summary information about a Profile (in Destiny 1, we called these \"Accounts\").
type DestinyEntitiesProfilesDestinyProfileComponent struct {
	// If you need to render the Profile (their platform name, icon, etc...) somewhere, this property contains that information.
	UserInfo UserUserInfoCard `json:"userInfo,omitempty"`
	// The last time the user played with any character on this Profile.
	DateLastPlayed time.Time `json:"dateLastPlayed,omitempty"`
	// If you want to know what expansions they own, this will contain that data.
	VersionsOwned DestinyDestinyGameVersions `json:"versionsOwned,omitempty"`
	// A list of the character IDs, for further querying on your part.
	CharacterIds []int64 `json:"characterIds,omitempty"`
}
