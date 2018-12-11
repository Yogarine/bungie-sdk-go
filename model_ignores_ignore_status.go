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
type IgnoresIgnoreStatus int32

// List of Ignores.IgnoreStatus
const (
	NOTIGNORED IgnoresIgnoreStatus = "0"
	IGNOREDUSER IgnoresIgnoreStatus = "1"
	IGNOREDGROUP IgnoresIgnoreStatus = "2"
	IGNOREDBYGROUP IgnoresIgnoreStatus = "4"
	IGNOREDPOST IgnoresIgnoreStatus = "8"
	IGNOREDTAG IgnoresIgnoreStatus = "16"
	IGNOREDGLOBAL IgnoresIgnoreStatus = "32"
)