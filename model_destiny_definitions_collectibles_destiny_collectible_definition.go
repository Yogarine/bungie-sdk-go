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

// Defines a
type DestinyDefinitionsCollectiblesDestinyCollectibleDefinition struct {
	DisplayProperties DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition `json:"displayProperties,omitempty"`
	// Indicates whether this Collectible's state is determined on a per-character or on an account-wide basis.
	Scope DestinyDestinyScope `json:"scope,omitempty"`
	// A human readable string for a hint about how to acquire the item.
	SourceString string `json:"sourceString,omitempty"`
	// This is a hash identifier we are building on the BNet side in an attempt to let people group collectibles by similar sources.  I can't promise that it's going to be 100% accurate, but if the designers were consistent in assigning the same source strings to items with the same sources, it *ought to* be. No promises though.  This hash also doesn't relate to an actual definition, just to note: we've got nothing useful other than the source string for this data.
	SourceHash int32 `json:"sourceHash,omitempty"`
	ItemHash int32 `json:"itemHash,omitempty"`
	AcquisitionInfo DestinyDefinitionsCollectiblesDestinyCollectibleAcquisitionBlock `json:"acquisitionInfo,omitempty"`
	StateInfo DestinyDefinitionsCollectiblesDestinyCollectibleStateBlock `json:"stateInfo,omitempty"`
	PresentationInfo DestinyDefinitionsPresentationDestinyPresentationChildBlock `json:"presentationInfo,omitempty"`
	// The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.  When entities refer to each other in Destiny content, it is this hash that they are referring to.
	Hash int32 `json:"hash,omitempty"`
	// The index of the entity as it was found in the investment tables.
	Index int32 `json:"index,omitempty"`
	// If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry!
	Redacted bool `json:"redacted,omitempty"`
}