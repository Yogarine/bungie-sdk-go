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
type GroupsV2GroupDateRange int32

// List of GroupsV2.GroupDateRange
const (
	ALL GroupsV2GroupDateRange = "0"
	PASTDAY GroupsV2GroupDateRange = "1"
	PASTWEEK GroupsV2GroupDateRange = "2"
	PASTMONTH GroupsV2GroupDateRange = "3"
	PASTYEAR GroupsV2GroupDateRange = "4"
)