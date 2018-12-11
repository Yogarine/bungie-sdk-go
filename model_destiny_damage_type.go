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
type DestinyDamageType int32

// List of Destiny.DamageType
const (
	NONE DestinyDamageType = "0"
	KINETIC DestinyDamageType = "1"
	ARC DestinyDamageType = "2"
	THERMAL DestinyDamageType = "3"
	VOID DestinyDamageType = "4"
	RAID DestinyDamageType = "5"
)