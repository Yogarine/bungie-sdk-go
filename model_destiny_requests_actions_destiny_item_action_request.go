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

type DestinyRequestsActionsDestinyItemActionRequest struct {
	ItemId int64 `json:"itemId,omitempty"`
	CharacterId int64 `json:"characterId,omitempty"`
	MembershipType BungieMembershipType `json:"membershipType,omitempty"`
}