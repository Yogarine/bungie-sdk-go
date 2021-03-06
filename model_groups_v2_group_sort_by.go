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
type GroupsV2GroupSortBy int32

// List of GroupsV2.GroupSortBy
const (
	NAME GroupsV2GroupSortBy = "0"
	DATE GroupsV2GroupSortBy = "1"
	POPULARITY GroupsV2GroupSortBy = "2"
	ID GroupsV2GroupSortBy = "3"
)