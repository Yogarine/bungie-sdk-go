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

// Nodes can have different visual states. This object represents a single visual state (\"highlight type\") that a node can be in, and the unlock expression condition to determine whether it should be set.
type DestinyDefinitionsDirectorDestinyActivityGraphNodeFeaturingStateDefinition struct {
	// The node can be highlighted in a variety of ways - the game iterates through these and finds the first FeaturingState that is valid at the present moment given the Game, Account, and Character state, and renders the node in that state. See the ActivityGraphNodeHighlightType enum for possible values.
	HighlightType DestinyActivityGraphNodeHighlightType `json:"highlightType,omitempty"`
}