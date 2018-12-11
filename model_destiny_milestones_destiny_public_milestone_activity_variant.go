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

// Represents a variant of an activity that's relevant to a milestone.
type DestinyMilestonesDestinyPublicMilestoneActivityVariant struct {
	// The hash identifier of this activity variant. Examine the activity's definition in the Manifest database to determine what makes it a distinct variant. Usually it will be difficulty level or whether or not it is a guided game variant of the activity, but theoretically it could be distinguished in any arbitrary way.
	ActivityHash int32 `json:"activityHash,omitempty"`
	// The hash identifier of the most specific Activity Mode under which this activity is played. This is useful for situations where the activity in question is - for instance - a PVP map, but it's not clear what mode the PVP map is being played under. If it's a playlist, this will be less specific: but hopefully useful in some way.
	ActivityModeHash int32 `json:"activityModeHash,omitempty"`
	// The enumeration equivalent of the most specific Activity Mode under which this activity is played.
	ActivityModeType int32 `json:"activityModeType,omitempty"`
}