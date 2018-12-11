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

// This component returns anything that could be considered \"Progression\" on a user: data where the user is gaining levels, reputation, completions, rewards, etc...
type DestinyEntitiesCharactersDestinyCharacterProgressionComponent struct {
	// A Dictionary of all known progressions for the Character, keyed by the Progression's hash.  Not all progressions have user-facing data, but those who do will have that data contained in the DestinyProgressionDefinition.
	Progressions map[string]DestinyDestinyProgression `json:"progressions,omitempty"`
	// A dictionary of all known Factions, keyed by the Faction's hash. It contains data about this character's status with the faction.
	Factions map[string]DestinyProgressionDestinyFactionProgression `json:"factions,omitempty"`
	// Milestones are related to the simple progressions shown in the game, but return additional and hopefully helpful information for users about the specifics of the Milestone's status.
	Milestones map[string]DestinyMilestonesDestinyMilestone `json:"milestones,omitempty"`
	// If the user has any active quests, the quests' statuses will be returned here.  Note that quests have been largely supplanted by Milestones, but that doesn't mean that they won't make a comeback independent of milestones at some point.
	Quests []DestinyQuestsDestinyQuestStatus `json:"quests,omitempty"`
	// Sometimes, you have items in your inventory that don't have instances, but still have Objective information. This provides you that objective information for uninstanced items.   This dictionary is keyed by the item's hash: which you can use to look up the name and description for the overall task(s) implied by the objective. The value is the list of objectives for this item, and their statuses.
	UninstancedItemObjectives map[string][]DestinyQuestsDestinyObjectiveProgress `json:"uninstancedItemObjectives,omitempty"`
	// The set of checklists that can be examined for this specific character, keyed by the hash identifier of the Checklist (DestinyChecklistDefinition)  For each checklist returned, its value is itself a Dictionary keyed by the checklist's hash identifier with the value being a boolean indicating if it's been discovered yet.
	Checklists map[string]map[string]bool `json:"checklists,omitempty"`
}