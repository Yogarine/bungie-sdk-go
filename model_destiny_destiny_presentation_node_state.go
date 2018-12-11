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
// DestinyDestinyPresentationNodeState : I know this doesn't look like a Flags Enumeration right now, but I assure you it is. This is the possible states that a Presentation Node can be in, and it is almost certain that its potential states will increase in the future. So don't treat it like a straight up enumeration.
type DestinyDestinyPresentationNodeState int32

// List of Destiny.DestinyPresentationNodeState
const (
	NONE DestinyDestinyPresentationNodeState = "0"
	INVISIBLE DestinyDestinyPresentationNodeState = "1"
	OBSCURED DestinyDestinyPresentationNodeState = "2"
)