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

type DestinyAdvancedAwaUserResponse struct {
	// Indication of the selection the user has made (Approving or rejecting the action)
	Selection DestinyAdvancedAwaUserSelection `json:"selection,omitempty"`
	// Correlation ID of the request
	CorrelationId string `json:"correlationId,omitempty"`
	// Secret nonce received via the PUSH notification.
	Nonce []string `json:"nonce,omitempty"`
}