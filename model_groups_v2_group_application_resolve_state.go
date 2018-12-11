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
type GroupsV2GroupApplicationResolveState int32

// List of GroupsV2.GroupApplicationResolveState
const (
	UNRESOLVED GroupsV2GroupApplicationResolveState = "0"
	ACCEPTED GroupsV2GroupApplicationResolveState = "1"
	DENIED GroupsV2GroupApplicationResolveState = "2"
	RESCINDED GroupsV2GroupApplicationResolveState = "3"
)