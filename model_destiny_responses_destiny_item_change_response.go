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

type DestinyResponsesDestinyItemChangeResponse struct {
	Item DestinyResponsesDestinyItemResponse `json:"item,omitempty"`
	// Items that appeared in the inventory possibly as a result of an action.
	AddedInventoryItems []DestinyEntitiesItemsDestinyItemComponent `json:"addedInventoryItems,omitempty"`
	// Items that disappeared from the inventory possibly as a result of an action.
	RemovedInventoryItems []DestinyEntitiesItemsDestinyItemComponent `json:"removedInventoryItems,omitempty"`
}
