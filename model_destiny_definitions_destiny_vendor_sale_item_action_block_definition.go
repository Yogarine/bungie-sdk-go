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

// Not terribly useful, some basic cooldown interaction info.
type DestinyDefinitionsDestinyVendorSaleItemActionBlockDefinition struct {
	ExecuteSeconds float32 `json:"executeSeconds,omitempty"`
	IsPositive bool `json:"isPositive,omitempty"`
}
