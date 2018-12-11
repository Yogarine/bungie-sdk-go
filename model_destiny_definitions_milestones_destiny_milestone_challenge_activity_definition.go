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

type DestinyDefinitionsMilestonesDestinyMilestoneChallengeActivityDefinition struct {
	// The activity for which this challenge is active.
	ActivityHash int32 `json:"activityHash,omitempty"`
	Challenges []DestinyDefinitionsMilestonesDestinyMilestoneChallengeDefinition `json:"challenges,omitempty"`
	// If the activity and its challenge is visible on any of these nodes, it will be returned.
	ActivityGraphNodes []DestinyDefinitionsMilestonesDestinyMilestoneChallengeActivityGraphNodeEntry `json:"activityGraphNodes,omitempty"`
	// Phases related to this activity, if there are any.  These will be listed in the order in which they will appear in the actual activity.
	Phases []DestinyDefinitionsMilestonesDestinyMilestoneChallengeActivityPhase `json:"phases,omitempty"`
}