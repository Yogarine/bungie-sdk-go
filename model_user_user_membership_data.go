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

type UserUserMembershipData struct {
	// this allows you to see destiny memberships that are visible and linked to this account (regardless of whether or not they have characters on the world server)
	DestinyMemberships []UserUserInfoCard `json:"destinyMemberships,omitempty"`
	BungieNetUser UserGeneralUser `json:"bungieNetUser,omitempty"`
}