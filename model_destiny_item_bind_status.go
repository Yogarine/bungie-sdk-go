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
type DestinyItemBindStatus int32

// List of Destiny.ItemBindStatus
const (
	NOTBOUND DestinyItemBindStatus = "0"
	BOUNDTOCHARACTER DestinyItemBindStatus = "1"
	BOUNDTOACCOUNT DestinyItemBindStatus = "2"
	BOUNDTOGUILD DestinyItemBindStatus = "3"
)