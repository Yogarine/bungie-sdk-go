# DestinyDefinitionsDestinyVendorInventoryFlyoutDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LockedDescription** | **string** | If the flyout is locked, this is the reason why. | [optional] 
**DisplayProperties** | [**DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition**](Destiny.Definitions.Common.DestinyDisplayPropertiesDefinition.md) | The title and other common properties of the flyout. | [optional] 
**Buckets** | [**[]DestinyDefinitionsDestinyVendorInventoryFlyoutBucketDefinition**](Destiny.Definitions.DestinyVendorInventoryFlyoutBucketDefinition.md) | A list of inventory buckets and other metadata to show on the screen. | [optional] 
**FlyoutId** | **int32** | An identifier for the flyout, in case anything else needs to refer to them. | [optional] 
**SuppressNewness** | **bool** | If this is true, don&#39;t show any of the glistening \&quot;this is a new item\&quot; UI elements, like we show on the inventory items themselves in in-game UI. | [optional] 
**EquipmentSlotHash** | **int32** | If this flyout is meant to show you the contents of the player&#39;s equipment slot, this is the slot to show. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


