# DestinyResponsesDestinyCharacterResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inventory** | [**SingleComponentResponseOfDestinyInventoryComponent**](SingleComponentResponseOfDestinyInventoryComponent.md) | The character-level non-equipped inventory items.  COMPONENT TYPE: CharacterInventories | [optional] 
**Character** | [**SingleComponentResponseOfDestinyCharacterComponent**](SingleComponentResponseOfDestinyCharacterComponent.md) | Base information about the character in question.  COMPONENT TYPE: Characters | [optional] 
**Progressions** | [**SingleComponentResponseOfDestinyCharacterProgressionComponent**](SingleComponentResponseOfDestinyCharacterProgressionComponent.md) | Character progression data, including Milestones.  COMPONENT TYPE: CharacterProgressions | [optional] 
**RenderData** | [**SingleComponentResponseOfDestinyCharacterRenderComponent**](SingleComponentResponseOfDestinyCharacterRenderComponent.md) | Character rendering data - a minimal set of information about equipment and dyes used for rendering.  COMPONENT TYPE: CharacterRenderData | [optional] 
**Activities** | [**SingleComponentResponseOfDestinyCharacterActivitiesComponent**](SingleComponentResponseOfDestinyCharacterActivitiesComponent.md) | Activity data - info about current activities available to the player.  COMPONENT TYPE: CharacterActivities | [optional] 
**Equipment** | [**SingleComponentResponseOfDestinyInventoryComponent**](SingleComponentResponseOfDestinyInventoryComponent.md) | Equipped items on the character.  COMPONENT TYPE: CharacterEquipment | [optional] 
**Kiosks** | [**SingleComponentResponseOfDestinyKiosksComponent**](SingleComponentResponseOfDestinyKiosksComponent.md) | Items available from Kiosks that are available to this specific character.   COMPONENT TYPE: Kiosks | [optional] 
**PlugSets** | [**SingleComponentResponseOfDestinyPlugSetsComponent**](SingleComponentResponseOfDestinyPlugSetsComponent.md) | When sockets refer to reusable Plug Sets (see DestinyPlugSetDefinition for more info), this is the set of plugs and their states that are scoped to this character.  This comes back with ItemSockets, as it is needed for a complete picture of the sockets on requested items.  COMPONENT TYPE: ItemSockets | [optional] 
**PresentationNodes** | [**SingleComponentResponseOfDestinyPresentationNodesComponent**](SingleComponentResponseOfDestinyPresentationNodesComponent.md) | COMPONENT TYPE: PresentationNodes | [optional] 
**Records** | [**SingleComponentResponseOfDestinyCharacterRecordsComponent**](SingleComponentResponseOfDestinyCharacterRecordsComponent.md) | COMPONENT TYPE: Records | [optional] 
**Collectibles** | [**SingleComponentResponseOfDestinyCollectiblesComponent**](SingleComponentResponseOfDestinyCollectiblesComponent.md) | COMPONENT TYPE: Collectibles | [optional] 
**ItemComponents** | [**DestinyItemComponentSetOfint64**](DestinyItemComponentSetOfint64.md) | The set of components belonging to the player&#39;s instanced items.  COMPONENT TYPE: [See inside the DestinyItemComponentSet contract for component types.] | [optional] 
**UninstancedItemComponents** | [**DestinyBaseItemComponentSetOfuint32**](DestinyBaseItemComponentSetOfuint32.md) | The set of components belonging to the player&#39;s UNinstanced items. Because apparently now those too can have information relevant to the character&#39;s state.  COMPONENT TYPE: [See inside the DestinyItemComponentSet contract for component types.] | [optional] 
**CurrencyLookups** | [**SingleComponentResponseOfDestinyCurrenciesComponent**](SingleComponentResponseOfDestinyCurrenciesComponent.md) | A \&quot;lookup\&quot; convenience component that can be used to quickly check if the character has access to items that can be used for purchasing.  COMPONENT TYPE: CurrencyLookups | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


