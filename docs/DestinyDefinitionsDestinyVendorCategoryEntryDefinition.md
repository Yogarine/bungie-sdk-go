# DestinyDefinitionsDestinyVendorCategoryEntryDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryIndex** | **int32** | The index of the category in the original category definitions for the vendor. | [optional] 
**CategoryId** | **string** | The string identifier of the category. | [optional] 
**SortValue** | **int32** | Used in sorting items in vendors... but there&#39;s a lot more to it. Just go with the order provided in the itemIndexes property on the DestinyVendorCategoryComponent instead, it should be more reliable than trying to recalculate it yourself. | [optional] 
**CategoryHash** | **int32** | The hashed identifier for the category. | [optional] 
**QuantityAvailable** | **int32** | The amount of items that will be available when this category is shown. | [optional] 
**ShowUnavailableItems** | **bool** | If items aren&#39;t up for sale in this category, should we still show them (greyed out)? | [optional] 
**HideIfNoCurrency** | **bool** | If you don&#39;t have the currency required to buy items from this category, should the items be hidden? | [optional] 
**HideFromRegularPurchase** | **bool** | True if this category doesn&#39;t allow purchases. | [optional] 
**BuyStringOverride** | **string** | The localized string for making purchases from this category, if it is different from the vendor&#39;s string for purchasing. | [optional] 
**DisabledDescription** | **string** | If the category is disabled, this is the localized description to show. | [optional] 
**DisplayTitle** | **string** | The localized title of the category. | [optional] 
**Overlay** | [**DestinyDefinitionsDestinyVendorCategoryOverlayDefinition**](Destiny.Definitions.DestinyVendorCategoryOverlayDefinition.md) | If this category has an overlay prompt that should appear, this contains the details of that prompt. | [optional] 
**VendorItemIndexes** | **[]int32** | A shortcut for the vendor item indexes sold under this category. Saves us from some expensive reorganization at runtime. | [optional] 
**IsPreview** | **bool** | Sometimes a category isn&#39;t actually used to sell items, but rather to preview them. This implies different UI (and manual placement of the category in the UI) in the game, and special treatment. | [optional] 
**IsDisplayOnly** | **bool** | If true, this category only displays items: you can&#39;t purchase anything in them. | [optional] 
**ResetIntervalMinutesOverride** | **int32** |  | [optional] 
**ResetOffsetMinutesOverride** | **int32** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


