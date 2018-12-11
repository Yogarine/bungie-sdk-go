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
// DestinyDestinyStatCategory : At last, stats have categories. Use this for whatever purpose you might wish.
type DestinyDestinyStatCategory int32

// List of Destiny.DestinyStatCategory
const (
	GAMEPLAY DestinyDestinyStatCategory = "0"
	WEAPON DestinyDestinyStatCategory = "1"
	DEFENSE DestinyDestinyStatCategory = "2"
	PRIMARY DestinyDestinyStatCategory = "3"
)