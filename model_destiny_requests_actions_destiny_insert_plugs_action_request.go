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

type DestinyRequestsActionsDestinyInsertPlugsActionRequest struct {
	// Action token provided by the AwaGetActionToken API call.
	ActionToken string `json:"actionToken,omitempty"`
	// The instance ID of the item having a plug inserted. Only instanced items can have sockets.
	ItemInstanceId int64 `json:"itemInstanceId,omitempty"`
	// The plugs being inserted.
	Plug DestinyRequestsActionsDestinyInsertPlugsRequestEntry `json:"plug,omitempty"`
	CharacterId int64 `json:"characterId,omitempty"`
	MembershipType BungieMembershipType `json:"membershipType,omitempty"`
}