# DestinyComponentsCollectiblesDestinyProfileCollectiblesComponent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecentCollectibleHashes** | **[]int32** | The list of collectibles determined by the game as having been \&quot;recently\&quot; acquired. | [optional] 
**NewnessFlaggedCollectibleHashes** | **[]int32** | The list of collectibles determined by the game as having been \&quot;recently\&quot; acquired.  The game client itself actually controls this data, so I personally question whether anyone will get much use out of this: because we can&#39;t edit this value through the API. But in case anyone finds it useful, here it is. | [optional] 
**Collectibles** | [**map[string]DestinyComponentsCollectiblesDestinyCollectibleComponent**](Destiny.Components.Collectibles.DestinyCollectibleComponent.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


