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
type ContentModelsContentPropertyDataTypeEnum int32

// List of Content.Models.ContentPropertyDataTypeEnum
const (
	NONE ContentModelsContentPropertyDataTypeEnum = "0"
	PLAINTEXT ContentModelsContentPropertyDataTypeEnum = "1"
	HTML ContentModelsContentPropertyDataTypeEnum = "2"
	DROPDOWN ContentModelsContentPropertyDataTypeEnum = "3"
	LIST ContentModelsContentPropertyDataTypeEnum = "4"
	JSON ContentModelsContentPropertyDataTypeEnum = "5"
	CONTENT ContentModelsContentPropertyDataTypeEnum = "6"
	REPRESENTATION ContentModelsContentPropertyDataTypeEnum = "7"
	SET ContentModelsContentPropertyDataTypeEnum = "8"
	FILE ContentModelsContentPropertyDataTypeEnum = "9"
	FOLDERSET ContentModelsContentPropertyDataTypeEnum = "10"
	DATE ContentModelsContentPropertyDataTypeEnum = "11"
	MULTILINEPLAINTEXT ContentModelsContentPropertyDataTypeEnum = "12"
	DESTINYCONTENT ContentModelsContentPropertyDataTypeEnum = "13"
	COLOR ContentModelsContentPropertyDataTypeEnum = "14"
)