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

// All the partnership info that's fit to expose externally, if we care to do so.
type PartnershipsPublicPartnershipDetail struct {
	PartnerType PartnershipsPartnershipType `json:"partnerType,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Name string `json:"name,omitempty"`
	Icon string `json:"icon,omitempty"`
}