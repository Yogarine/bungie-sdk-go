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

// The response object for retrieving an individual instanced item. None of these components are relevant for an item that doesn't have an \"itemInstanceId\": for those, get your information from the DestinyInventoryDefinition.
type DestinyResponsesDestinyItemResponse struct {
	// If the item is on a character, this will return the ID of the character that is holding the item.
	CharacterId int64 `json:"characterId,omitempty"`
	// Common data for the item relevant to its non-instanced properties.  COMPONENT TYPE: ItemCommonData
	Item SingleComponentResponseOfDestinyItemComponent `json:"item,omitempty"`
	// Basic instance data for the item.  COMPONENT TYPE: ItemInstances
	Instance SingleComponentResponseOfDestinyItemInstanceComponent `json:"instance,omitempty"`
	// Information specifically about the item's objectives.  COMPONENT TYPE: ItemObjectives
	Objectives SingleComponentResponseOfDestinyItemObjectivesComponent `json:"objectives,omitempty"`
	// Information specifically about the perks currently active on the item.  COMPONENT TYPE: ItemPerks
	Perks SingleComponentResponseOfDestinyItemPerksComponent `json:"perks,omitempty"`
	// Information about how to render the item in 3D.  COMPONENT TYPE: ItemRenderData
	RenderData SingleComponentResponseOfDestinyItemRenderComponent `json:"renderData,omitempty"`
	// Information about the computed stats of the item: power, defense, etc...  COMPONENT TYPE: ItemStats
	Stats SingleComponentResponseOfDestinyItemStatsComponent `json:"stats,omitempty"`
	// Information about the talent grid attached to the item. Talent nodes can provide a variety of benefits and abilities, and in Destiny 2 are used almost exclusively for the character's \"Builds\".  COMPONENT TYPE: ItemTalentGrids
	TalentGrid SingleComponentResponseOfDestinyItemTalentGridComponent `json:"talentGrid,omitempty"`
	// Information about the sockets of the item: which are currently active, what potential sockets you could have and the stats/abilities/perks you can gain from them.  COMPONENT TYPE: ItemSockets
	Sockets SingleComponentResponseOfDestinyItemSocketsComponent `json:"sockets,omitempty"`
}
