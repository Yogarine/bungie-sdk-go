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
type ForumForumTopicsQuickDateEnum int32

// List of Forum.ForumTopicsQuickDateEnum
const (
	ALL ForumForumTopicsQuickDateEnum = "0"
	LASTYEAR ForumForumTopicsQuickDateEnum = "1"
	LASTMONTH ForumForumTopicsQuickDateEnum = "2"
	LASTWEEK ForumForumTopicsQuickDateEnum = "3"
	LASTDAY ForumForumTopicsQuickDateEnum = "4"
)