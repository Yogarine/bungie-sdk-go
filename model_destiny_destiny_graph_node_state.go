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
// DestinyDestinyGraphNodeState : Represents a potential state of an Activity Graph node.
type DestinyDestinyGraphNodeState int32

// List of Destiny.DestinyGraphNodeState
const (
	HIDDEN DestinyDestinyGraphNodeState = "0"
	VISIBLE DestinyDestinyGraphNodeState = "1"
	TEASER DestinyDestinyGraphNodeState = "2"
	INCOMPLETE DestinyDestinyGraphNodeState = "3"
	COMPLETED DestinyDestinyGraphNodeState = "4"
)