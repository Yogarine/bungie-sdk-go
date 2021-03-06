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
// DestinyDestinyGameVersions : A flags enumeration indicating the versions of the game that a given user has purchased.
type DestinyDestinyGameVersions int32

// List of Destiny.DestinyGameVersions
const (
	NONE DestinyDestinyGameVersions = "0"
	DESTINY2 DestinyDestinyGameVersions = "1"
	DLC1 DestinyDestinyGameVersions = "2"
	DLC2 DestinyDestinyGameVersions = "4"
	FORSAKEN DestinyDestinyGameVersions = "8"
	YEARTWOANNUALPASS DestinyDestinyGameVersions = "16"
)