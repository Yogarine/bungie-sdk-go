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
// DestinySocketPlugSources : Indicates how a socket is populated, and where you should look for valid plug data. This is a flags enumeration field, as you may have to look in multiple sources across multiple components for valid plugs.
type DestinySocketPlugSources int32

// List of Destiny.SocketPlugSources
const (
	NONE DestinySocketPlugSources = "0"
	INVENTORYSOURCED DestinySocketPlugSources = "1"
	REUSABLEPLUGITEMS DestinySocketPlugSources = "2"
	PROFILEPLUGSET DestinySocketPlugSources = "4"
	CHARACTERPLUGSET DestinySocketPlugSources = "8"
)