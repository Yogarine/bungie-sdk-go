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

type TrendingTrendingDetail struct {
	Identifier string `json:"identifier,omitempty"`
	EntityType TrendingTrendingEntryType `json:"entityType,omitempty"`
	News TrendingTrendingEntryNews `json:"news,omitempty"`
	Support TrendingTrendingEntrySupportArticle `json:"support,omitempty"`
	DestinyItem TrendingTrendingEntryDestinyItem `json:"destinyItem,omitempty"`
	DestinyActivity TrendingTrendingEntryDestinyActivity `json:"destinyActivity,omitempty"`
	DestinyRitual TrendingTrendingEntryDestinyRitual `json:"destinyRitual,omitempty"`
	Creation TrendingTrendingEntryCommunityCreation `json:"creation,omitempty"`
	Stream TrendingTrendingEntryCommunityStream `json:"stream,omitempty"`
}
