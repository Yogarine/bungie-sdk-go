# DestinyDefinitionsRecordsDestinyRecordDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayProperties** | [**DestinyDefinitionsCommonDestinyDisplayPropertiesDefinition**](Destiny.Definitions.Common.DestinyDisplayPropertiesDefinition.md) |  | [optional] 
**Scope** | [**DestinyDestinyScope**](Destiny.DestinyScope.md) | Indicates whether this Record&#39;s state is determined on a per-character or on an account-wide basis. | [optional] 
**PresentationInfo** | [**DestinyDefinitionsPresentationDestinyPresentationChildBlock**](Destiny.Definitions.Presentation.DestinyPresentationChildBlock.md) |  | [optional] 
**LoreHash** | **int32** |  | [optional] 
**ObjectiveHashes** | **[]int32** |  | [optional] 
**RecordValueStyle** | [**DestinyDestinyRecordValueStyle**](Destiny.DestinyRecordValueStyle.md) |  | [optional] 
**TitleInfo** | [**DestinyDefinitionsRecordsDestinyRecordTitleBlock**](Destiny.Definitions.Records.DestinyRecordTitleBlock.md) |  | [optional] 
**CompletionInfo** | [**DestinyDefinitionsRecordsDestinyRecordCompletionBlock**](Destiny.Definitions.Records.DestinyRecordCompletionBlock.md) |  | [optional] 
**StateInfo** | [**DestinyDefinitionsRecordsSchemaRecordStateBlock**](Destiny.Definitions.Records.SchemaRecordStateBlock.md) |  | [optional] 
**Requirements** | [**DestinyDefinitionsPresentationDestinyPresentationNodeRequirementsBlock**](Destiny.Definitions.Presentation.DestinyPresentationNodeRequirementsBlock.md) |  | [optional] 
**Hash** | **int32** | The unique identifier for this entity. Guaranteed to be unique for the type of entity, but not globally.  When entities refer to each other in Destiny content, it is this hash that they are referring to. | [optional] 
**Index** | **int32** | The index of the entity as it was found in the investment tables. | [optional] 
**Redacted** | **bool** | If this is true, then there is an entity with this identifier/type combination, but BNet is not yet allowed to show it. Sorry! | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


