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

type ApplicationsSeries struct {
	// Collection of samples with time and value.
	Datapoints []ApplicationsDatapoint `json:"datapoints,omitempty"`
	// Target to which to datapoints apply.
	Target string `json:"target,omitempty"`
}
