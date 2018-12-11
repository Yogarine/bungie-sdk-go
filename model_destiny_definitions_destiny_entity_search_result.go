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

// The results of a search for Destiny content. This will be improved on over time, I've been doing some experimenting to see what might be useful.
type DestinyDefinitionsDestinyEntitySearchResult struct {
	// A list of suggested words that might make for better search results, based on the text searched for.
	SuggestedWords []string `json:"suggestedWords,omitempty"`
	// The items found that are matches/near matches for the searched-for term, sorted by something vaguely resembling \"relevance\". Hopefully this will get better in the future.
	Results SearchResultOfDestinyEntitySearchResultItem `json:"results,omitempty"`
}