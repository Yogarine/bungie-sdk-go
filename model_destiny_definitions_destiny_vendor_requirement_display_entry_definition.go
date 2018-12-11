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

// The localized properties of the requirementsDisplay, allowing information about the requirement or item being featured to be seen.
type DestinyDefinitionsDestinyVendorRequirementDisplayEntryDefinition struct {
	Icon string `json:"icon,omitempty"`
	Name string `json:"name,omitempty"`
	Source string `json:"source,omitempty"`
	Type string `json:"type,omitempty"`
}