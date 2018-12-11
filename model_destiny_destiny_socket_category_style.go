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
// DestinyDestinySocketCategoryStyle : Represents the possible and known UI styles used by the game for rendering Socket Categories.
type DestinyDestinySocketCategoryStyle int32

// List of Destiny.DestinySocketCategoryStyle
const (
	UNKNOWN DestinyDestinySocketCategoryStyle = "0"
	REUSABLE DestinyDestinySocketCategoryStyle = "1"
	CONSUMABLE DestinyDestinySocketCategoryStyle = "2"
	UNLOCKABLE DestinyDestinySocketCategoryStyle = "3"
	INTRINSIC DestinyDestinySocketCategoryStyle = "4"
)