# DestinyDefinitionsDestinyActivityPlaylistItemDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityHash** | **int32** | The hash identifier of the Activity that can be played. Use it to look up the DestinyActivityDefinition. | [optional] 
**DirectActivityModeHash** | **int32** | If this playlist entry had an activity mode directly defined on it, this will be the hash of that mode. | [optional] 
**DirectActivityModeType** | **int32** | If the playlist entry had an activity mode directly defined on it, this will be the enum value of that mode. | [optional] 
**ActivityModeHashes** | **[]int32** | The hash identifiers for Activity Modes relevant to this entry. | [optional] 
**ActivityModeTypes** | [**[]DestinyHistoricalStatsDefinitionsDestinyActivityModeType**](Destiny.HistoricalStats.Definitions.DestinyActivityModeType.md) | The activity modes - if any - in enum form. Because we can&#39;t seem to escape the enums. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


