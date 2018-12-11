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
type DestinyDestinyTalentNodeState int32

// List of Destiny.DestinyTalentNodeState
const (
	INVALID DestinyDestinyTalentNodeState = "0"
	CANUPGRADE DestinyDestinyTalentNodeState = "1"
	NOPOINTS DestinyDestinyTalentNodeState = "2"
	NOPREREQUISITES DestinyDestinyTalentNodeState = "3"
	NOSTEPS DestinyDestinyTalentNodeState = "4"
	NOUNLOCK DestinyDestinyTalentNodeState = "5"
	NOMATERIAL DestinyDestinyTalentNodeState = "6"
	NOGRIDLEVEL DestinyDestinyTalentNodeState = "7"
	SWAPPINGLOCKED DestinyDestinyTalentNodeState = "8"
	MUSTSWAP DestinyDestinyTalentNodeState = "9"
	COMPLETE DestinyDestinyTalentNodeState = "10"
	UNKNOWN DestinyDestinyTalentNodeState = "11"
	CREATIONONLY DestinyDestinyTalentNodeState = "12"
	HIDDEN DestinyDestinyTalentNodeState = "13"
)