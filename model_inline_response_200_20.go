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

type InlineResponse20020 struct {
	Response []GroupsV2GroupV2Card `json:"Response,omitempty"`
	ErrorCode ExceptionsPlatformErrorCodes `json:"ErrorCode,omitempty"`
	ThrottleSeconds int32 `json:"ThrottleSeconds,omitempty"`
	ErrorStatus string `json:"ErrorStatus,omitempty"`
	Message string `json:"Message,omitempty"`
	MessageData map[string]string `json:"MessageData,omitempty"`
	DetailedErrorTrace string `json:"DetailedErrorTrace,omitempty"`
}