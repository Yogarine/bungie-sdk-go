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

// The definition of a known, reusable plug that can be applied to a socket.
type DestinyDefinitionsDestinyItemSocketEntryPlugItemDefinition struct {
	// The hash identifier of a DestinyInventoryItemDefinition representing the plug that can be inserted.
	PlugItemHash int32 `json:"plugItemHash,omitempty"`
}
