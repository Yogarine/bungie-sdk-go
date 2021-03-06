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
// DestinyDestinyCollectibleState : A Flags Enumeration where each bit represents a different state that the Collectible can be in. A collectible can be in any number of these states, and you can choose to use or ignore any or all of them when making your own UI that shows Collectible info. Our displays are going to honor them, but we're also the kind of people who only pretend to inhale before quickly passing it to the left. So, you know, do what you got to do.  (All joking aside, please note the caveat I mention around the Invisible flag: there are cases where it is in the best interest of your users to honor these flags even if you're a \"show all the data\" person. Collector-oriented compulsion is a very unfortunate and real thing, and I would hate to instill that compulsion in others through showing them items that they cannot earn. Please consider this when you are making your own apps/sites.)
type DestinyDestinyCollectibleState int32

// List of Destiny.DestinyCollectibleState
const (
	NONE DestinyDestinyCollectibleState = "0"
	NOTACQUIRED DestinyDestinyCollectibleState = "1"
	OBSCURED DestinyDestinyCollectibleState = "2"
	INVISIBLE DestinyDestinyCollectibleState = "4"
	CANNOTAFFORDMATERIALREQUIREMENTS DestinyDestinyCollectibleState = "8"
	INVENTORYSPACEUNAVAILABLE DestinyDestinyCollectibleState = "16"
	UNIQUENESSVIOLATION DestinyDestinyCollectibleState = "32"
	PURCHASEDISABLED DestinyDestinyCollectibleState = "64"
)