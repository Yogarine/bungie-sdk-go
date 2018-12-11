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
// DestinyDestinyPresentationScreenStyle : A hint for what screen should be shown when this presentation node is clicked into. How you use this is your UI is up to you.
type DestinyDestinyPresentationScreenStyle int32

// List of Destiny.DestinyPresentationScreenStyle
const (
	DEFAULT DestinyDestinyPresentationScreenStyle = "0"
	CATEGORYSETS DestinyDestinyPresentationScreenStyle = "1"
	BADGE DestinyDestinyPresentationScreenStyle = "2"
)