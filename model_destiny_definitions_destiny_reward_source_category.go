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
// DestinyDefinitionsDestinyRewardSourceCategory : BNet's custom categorization of reward sources. We took a look at the existing ways that items could be spawned, and tried to make high-level categorizations of them. This needs to be re-evaluated for Destiny 2.
type DestinyDefinitionsDestinyRewardSourceCategory int32

// List of Destiny.Definitions.DestinyRewardSourceCategory
const (
	NONE DestinyDefinitionsDestinyRewardSourceCategory = "0"
	ACTIVITY DestinyDefinitionsDestinyRewardSourceCategory = "1"
	VENDOR DestinyDefinitionsDestinyRewardSourceCategory = "2"
	AGGREGATE DestinyDefinitionsDestinyRewardSourceCategory = "3"
)