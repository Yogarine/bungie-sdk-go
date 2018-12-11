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

// Represents localized, extended content related to Milestones. This is intentionally returned by a separate endpoint and not with Character-level Milestone data because we do not put localized data into standard Destiny responses, both for brevity of response and for caching purposes. If you really need this data, hit the Milestone Content endpoint.
type DestinyMilestonesDestinyMilestoneContent struct {
	// The \"About this Milestone\" text from the Firehose.
	About string `json:"about,omitempty"`
	// The Current Status of the Milestone, as driven by the Firehose.
	Status string `json:"status,omitempty"`
	// A list of tips, provided by the Firehose.
	Tips []string `json:"tips,omitempty"`
	// If DPS has defined items related to this Milestone, they can categorize those items in the Firehose. That data will then be returned as item categories here.
	ItemCategories []DestinyMilestonesDestinyMilestoneContentItemCategory `json:"itemCategories,omitempty"`
}
