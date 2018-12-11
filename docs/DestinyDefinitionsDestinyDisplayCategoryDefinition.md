# DestinyDefinitionsDestinyDisplayCategoryDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | A string identifier for the display category. | [optional] 
**DisplayCategoryHash** | **int32** |  | [optional] 
**DisplayProperties** | [**DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition**](Destiny.Definitions.Common.DestinyDisplayPropertiesDefinition.md) |  | [optional] 
**DisplayInBanner** | **bool** | If true, this category should be displayed in the \&quot;Banner\&quot; section of the vendor&#39;s UI. | [optional] 
**ProgressionHash** | **int32** | If it exists, this is the hash identifier of a DestinyProgressionDefinition that represents the progression to show on this display category.  Specific categories can now have thier own distinct progression, apparently. So that&#39;s cool. | [optional] 
**SortOrder** | [**DestinyVendorDisplayCategorySortOrder**](Destiny.VendorDisplayCategorySortOrder.md) | If this category sorts items in a nonstandard way, this will be the way we sort. | [optional] 
**DisplayStyleHash** | **int32** | An indicator of how the category will be displayed in the UI. It&#39;s up to you to do something cool or interesting in response to this, or just to treat it as a normal category. | [optional] 
**DisplayStyleIdentifier** | **string** | An indicator of how the category will be displayed in the UI. It&#39;s up to you to do something cool or interesting in response to this, or just to treat it as a normal category. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


