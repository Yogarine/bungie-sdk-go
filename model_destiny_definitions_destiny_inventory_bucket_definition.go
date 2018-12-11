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

// An Inventory (be it Character or Profile level) is comprised of many Buckets. An example of a bucket is \"Primary Weapons\", where all of the primary weapons on a character are gathered together into a single visual element in the UI: a subset of the inventory that has a limited number of slots, and in this case also has an associated Equipment Slot for equipping an item in the bucket.  Item definitions declare what their \"default\" bucket is (DestinyInventoryItemDefinition.inventory.bucketTypeHash), and Item instances will tell you which bucket they are currently residing in (DestinyItemComponent.bucketHash). You can use this information along with the DestinyInventoryBucketDefinition to show these items grouped by bucket.  You cannot transfer an item to a bucket that is not its Default without going through a Vendor's \"accepted items\" (DestinyVendorDefinition.acceptedItems). This is how transfer functionality like the Vault is implemented, as a feature of a Vendor. See the vendor's acceptedItems property for more details.
type DestinyDefinitionsDestinyInventoryBucketDefinition struct {
	DisplayProperties DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition `json:"displayProperties,omitempty"`
	// Where the bucket is found. 0 = Character, 1 = Account
	Scope DestinyBucketScope `json:"scope,omitempty"`
	// An enum value for what items can be found in the bucket. See the BucketCategory enum for more details.
	Category DestinyBucketCategory `json:"category,omitempty"`
	// Use this property to provide a quick-and-dirty recommended ordering for buckets in the UI. Most UIs will likely want to forsake this for something more custom and manual.
	BucketOrder int32 `json:"bucketOrder,omitempty"`
	// The maximum # of item \"slots\" in a bucket. A slot is a given combination of item + quantity.  For instance, a Weapon will always take up a single slot, and always have a quantity of 1. But a material could take up only a single slot with hundreds of quantity.
	ItemCount int32 `json:"itemCount,omitempty"`
	// Sometimes, inventory buckets represent conceptual \"locations\" in the game that might not be expected. This value indicates the conceptual location of the bucket, regardless of where it is actually contained on the character/account.   See ItemLocation for details.   Note that location includes the Vault and the Postmaster (both of whom being just inventory buckets with additional actions that can be performed on them through a Vendor)
	Location DestinyItemLocation `json:"location,omitempty"`
	// If TRUE, there is at least one Vendor that can transfer items to/from this bucket. See the DestinyVendorDefinition's acceptedItems property for more information on how transferring works.
	HasTransferDestination bool `json:"hasTransferDestination,omitempty"`
	// If True, this bucket is enabled. Disabled buckets may include buckets that were included for test purposes, or that were going to be used but then were abandoned but never removed from content *cough*.
	Enabled bool `json:"enabled,omitempty"`
	// if a FIFO bucket fills up, it will delete the oldest item from said bucket when a new item tries to be added to it. If this is FALSE, the bucket will not allow new items to be placed in it until room is made by the user manually deleting items from it. You can see an example of this with the Postmaster's bucket.
	Fifo bool `json:"fifo,omitempty"`
	// The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.  When entities refer to each other in Destiny content, it is this hash that they are referring to.
	Hash int32 `json:"hash,omitempty"`
	// The index of the entity as it was found in the investment tables.
	Index int32 `json:"index,omitempty"`
	// If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry!
	Redacted bool `json:"redacted,omitempty"`
}