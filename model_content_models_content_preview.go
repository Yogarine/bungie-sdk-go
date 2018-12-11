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

type ContentModelsContentPreview struct {
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
	ItemInSet bool `json:"itemInSet,omitempty"`
	SetTag string `json:"setTag,omitempty"`
	SetNesting int32 `json:"setNesting,omitempty"`
	UseSetId int32 `json:"useSetId,omitempty"`
}