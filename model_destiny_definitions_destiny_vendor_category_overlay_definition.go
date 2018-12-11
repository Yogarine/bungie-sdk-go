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

// The details of an overlay prompt to show to a user. They are all fairly self-explanatory localized strings that can be shown.
type DestinyDefinitionsDestinyVendorCategoryOverlayDefinition struct {
	ChoiceDescription string `json:"choiceDescription,omitempty"`
	Description string `json:"description,omitempty"`
	Icon string `json:"icon,omitempty"`
	Title string `json:"title,omitempty"`
	// If this overlay has a currency item that it features, this is said featured item.
	CurrencyItemHash int32 `json:"currencyItemHash,omitempty"`
}