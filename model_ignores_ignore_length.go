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
type IgnoresIgnoreLength int32

// List of Ignores.IgnoreLength
const (
	NONE IgnoresIgnoreLength = "0"
	WEEK IgnoresIgnoreLength = "1"
	TWOWEEKS IgnoresIgnoreLength = "2"
	THREEWEEKS IgnoresIgnoreLength = "3"
	MONTH IgnoresIgnoreLength = "4"
	THREEMONTHS IgnoresIgnoreLength = "5"
	SIXMONTHS IgnoresIgnoreLength = "6"
	YEAR IgnoresIgnoreLength = "7"
	FOREVER IgnoresIgnoreLength = "8"
	THREEMINUTES IgnoresIgnoreLength = "9"
	HOUR IgnoresIgnoreLength = "10"
	THIRTYDAYS IgnoresIgnoreLength = "11"
)