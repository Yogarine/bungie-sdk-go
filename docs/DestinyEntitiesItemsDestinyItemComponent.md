# DestinyEntitiesItemsDestinyItemComponent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemHash** | **int32** | The identifier for the item&#39;s definition, which is where most of the useful static information for the item can be found. | [optional] 
**ItemInstanceId** | **int64** | If the item is instanced, it will have an instance ID. Lack of an instance ID implies that the item has no distinct local qualities aside from stack size. | [optional] 
**Quantity** | **int32** | The quantity of the item in this stack. Note that Instanced items cannot stack. If an instanced item, this value will always be 1 (as the stack has exactly one item in it) | [optional] 
**BindStatus** | [**DestinyItemBindStatus**](Destiny.ItemBindStatus.md) | If the item is bound to a location, it will be specified in this enum. | [optional] 
**Location** | [**DestinyItemLocation**](Destiny.ItemLocation.md) | An easy reference for where the item is located. Redundant if you got the item from an Inventory, but useful when making detail calls on specific items. | [optional] 
**BucketHash** | **int32** | The hash identifier for the specific inventory bucket in which the item is located. | [optional] 
**TransferStatus** | [**DestinyTransferStatuses**](Destiny.TransferStatuses.md) | If there is a known error state that would cause this item to not be transferable, this Flags enum will indicate all of those error states. Otherwise, it will be 0 (CanTransfer). | [optional] 
**Lockable** | **bool** | If the item can be locked, this will indicate that state. | [optional] 
**State** | [**DestinyItemState**](Destiny.ItemState.md) | A flags enumeration indicating the transient/custom states of the item that affect how it is rendered: whether it&#39;s tracked or locked for example, or whether it has a masterwork plug inserted. | [optional] 
**OverrideStyleItemHash** | **int32** | If populated, this is the hash of the item whose icon (and other secondary styles, but *not* the human readable strings) should override whatever icons/styles are on the item being sold.  If you don&#39;t do this, certain items whose styles are being overridden by socketed items - such as the \&quot;Recycle Shader\&quot; item - would show whatever their default icon/style is, and it wouldn&#39;t be pretty or look accurate. | [optional] 
**ExpirationDate** | [**time.Time**](time.Time.md) | If the item can expire, this is the date at which it will/did expire. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


