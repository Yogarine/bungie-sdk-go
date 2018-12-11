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
type GroupsV2MembershipOption int32

// List of GroupsV2.MembershipOption
const (
	REVIEWED GroupsV2MembershipOption = "0"
	OPEN GroupsV2MembershipOption = "1"
	CLOSED GroupsV2MembershipOption = "2"
)