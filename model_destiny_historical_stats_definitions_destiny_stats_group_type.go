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
// DestinyHistoricalStatsDefinitionsDestinyStatsGroupType : If the enum value is > 100, it is a \"special\" group that cannot be queried for directly (special cases apply to when they are returned, and are not relevant in general cases)
type DestinyHistoricalStatsDefinitionsDestinyStatsGroupType int32

// List of Destiny.HistoricalStats.Definitions.DestinyStatsGroupType
const (
	NONE DestinyHistoricalStatsDefinitionsDestinyStatsGroupType = "0"
	GENERAL DestinyHistoricalStatsDefinitionsDestinyStatsGroupType = "1"
	WEAPONS DestinyHistoricalStatsDefinitionsDestinyStatsGroupType = "2"
	MEDALS DestinyHistoricalStatsDefinitionsDestinyStatsGroupType = "3"
	RESERVEDGROUPS DestinyHistoricalStatsDefinitionsDestinyStatsGroupType = "100"
	LEADERBOARD DestinyHistoricalStatsDefinitionsDestinyStatsGroupType = "101"
	ACTIVITY DestinyHistoricalStatsDefinitionsDestinyStatsGroupType = "102"
	UNIQUEWEAPON DestinyHistoricalStatsDefinitionsDestinyStatsGroupType = "103"
	INTERNAL DestinyHistoricalStatsDefinitionsDestinyStatsGroupType = "104"
)