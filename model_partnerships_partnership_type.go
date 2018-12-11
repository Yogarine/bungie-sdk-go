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
// PartnershipsPartnershipType : Representing external partners to which BNet users can link accounts, but that are not Account System credentials: partnerships that BNet uses exclusively for data.
type PartnershipsPartnershipType int32

// List of Partnerships.PartnershipType
const (
	NONE PartnershipsPartnershipType = "0"
	TWITCH PartnershipsPartnershipType = "1"
)