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
type ForumForumPostCategoryEnums int32

// List of Forum.ForumPostCategoryEnums
const (
	NONE ForumForumPostCategoryEnums = "0"
	TEXTONLY ForumForumPostCategoryEnums = "1"
	MEDIA ForumForumPostCategoryEnums = "2"
	LINK ForumForumPostCategoryEnums = "4"
	POLL ForumForumPostCategoryEnums = "8"
	QUESTION ForumForumPostCategoryEnums = "16"
	ANSWERED ForumForumPostCategoryEnums = "32"
	ANNOUNCEMENT ForumForumPostCategoryEnums = "64"
	CONTENTCOMMENT ForumForumPostCategoryEnums = "128"
	BUNGIEOFFICIAL ForumForumPostCategoryEnums = "256"
	NINJAOFFICIAL ForumForumPostCategoryEnums = "512"
	RECRUITMENT ForumForumPostCategoryEnums = "1024"
)