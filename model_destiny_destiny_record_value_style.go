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
type DestinyDestinyRecordValueStyle int32

// List of Destiny.DestinyRecordValueStyle
const (
	INTEGER DestinyDestinyRecordValueStyle = "0"
	PERCENTAGE DestinyDestinyRecordValueStyle = "1"
	MILLISECONDS DestinyDestinyRecordValueStyle = "2"
	BOOLEAN DestinyDestinyRecordValueStyle = "3"
	DECIMAL DestinyDestinyRecordValueStyle = "4"
)