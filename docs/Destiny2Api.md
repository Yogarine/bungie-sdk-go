# Bungie\Destiny2Api

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Destiny2AwaGetActionToken**](Destiny2Api.md#Destiny2AwaGetActionToken) | **Get** /Destiny2/Awa/GetActionToken/{correlationId}/ | 
[**Destiny2AwaInitializeRequest**](Destiny2Api.md#Destiny2AwaInitializeRequest) | **Post** /Destiny2/Awa/Initialize/ | 
[**Destiny2AwaProvideAuthorizationResult**](Destiny2Api.md#Destiny2AwaProvideAuthorizationResult) | **Post** /Destiny2/Awa/AwaProvideAuthorizationResult/ | 
[**Destiny2EquipItem**](Destiny2Api.md#Destiny2EquipItem) | **Post** /Destiny2/Actions/Items/EquipItem/ | 
[**Destiny2EquipItems**](Destiny2Api.md#Destiny2EquipItems) | **Post** /Destiny2/Actions/Items/EquipItems/ | 
[**Destiny2GetActivityHistory**](Destiny2Api.md#Destiny2GetActivityHistory) | **Get** /Destiny2/{membershipType}/Account/{destinyMembershipId}/Character/{characterId}/Stats/Activities/ | 
[**Destiny2GetCharacter**](Destiny2Api.md#Destiny2GetCharacter) | **Get** /Destiny2/{membershipType}/Profile/{destinyMembershipId}/Character/{characterId}/ | 
[**Destiny2GetClanAggregateStats**](Destiny2Api.md#Destiny2GetClanAggregateStats) | **Get** /Destiny2/Stats/AggregateClanStats/{groupId}/ | 
[**Destiny2GetClanLeaderboards**](Destiny2Api.md#Destiny2GetClanLeaderboards) | **Get** /Destiny2/Stats/Leaderboards/Clans/{groupId}/ | 
[**Destiny2GetClanWeeklyRewardState**](Destiny2Api.md#Destiny2GetClanWeeklyRewardState) | **Get** /Destiny2/Clan/{groupId}/WeeklyRewardState/ | 
[**Destiny2GetCollectibleNodeDetails**](Destiny2Api.md#Destiny2GetCollectibleNodeDetails) | **Get** /Destiny2/{membershipType}/Profile/{destinyMembershipId}/Character/{characterId}/Collectibles/{collectiblePresentationNodeHash}/ | 
[**Destiny2GetDestinyAggregateActivityStats**](Destiny2Api.md#Destiny2GetDestinyAggregateActivityStats) | **Get** /Destiny2/{membershipType}/Account/{destinyMembershipId}/Character/{characterId}/Stats/AggregateActivityStats/ | 
[**Destiny2GetDestinyEntityDefinition**](Destiny2Api.md#Destiny2GetDestinyEntityDefinition) | **Get** /Destiny2/Manifest/{entityType}/{hashIdentifier}/ | 
[**Destiny2GetDestinyManifest**](Destiny2Api.md#Destiny2GetDestinyManifest) | **Get** /Destiny2/Manifest/ | 
[**Destiny2GetHistoricalStats**](Destiny2Api.md#Destiny2GetHistoricalStats) | **Get** /Destiny2/{membershipType}/Account/{destinyMembershipId}/Character/{characterId}/Stats/ | 
[**Destiny2GetHistoricalStatsDefinition**](Destiny2Api.md#Destiny2GetHistoricalStatsDefinition) | **Get** /Destiny2/Stats/Definition/ | 
[**Destiny2GetHistoricalStatsForAccount**](Destiny2Api.md#Destiny2GetHistoricalStatsForAccount) | **Get** /Destiny2/{membershipType}/Account/{destinyMembershipId}/Stats/ | 
[**Destiny2GetItem**](Destiny2Api.md#Destiny2GetItem) | **Get** /Destiny2/{membershipType}/Profile/{destinyMembershipId}/Item/{itemInstanceId}/ | 
[**Destiny2GetLeaderboards**](Destiny2Api.md#Destiny2GetLeaderboards) | **Get** /Destiny2/{membershipType}/Account/{destinyMembershipId}/Stats/Leaderboards/ | 
[**Destiny2GetLeaderboardsForCharacter**](Destiny2Api.md#Destiny2GetLeaderboardsForCharacter) | **Get** /Destiny2/Stats/Leaderboards/{membershipType}/{destinyMembershipId}/{characterId}/ | 
[**Destiny2GetLinkedProfiles**](Destiny2Api.md#Destiny2GetLinkedProfiles) | **Get** /Destiny2/{membershipType}/Profile/{membershipId}/LinkedProfiles/ | 
[**Destiny2GetPostGameCarnageReport**](Destiny2Api.md#Destiny2GetPostGameCarnageReport) | **Get** /Destiny2/Stats/PostGameCarnageReport/{activityId}/ | 
[**Destiny2GetProfile**](Destiny2Api.md#Destiny2GetProfile) | **Get** /Destiny2/{membershipType}/Profile/{destinyMembershipId}/ | 
[**Destiny2GetPublicMilestoneContent**](Destiny2Api.md#Destiny2GetPublicMilestoneContent) | **Get** /Destiny2/Milestones/{milestoneHash}/Content/ | 
[**Destiny2GetPublicMilestones**](Destiny2Api.md#Destiny2GetPublicMilestones) | **Get** /Destiny2/Milestones/ | 
[**Destiny2GetUniqueWeaponHistory**](Destiny2Api.md#Destiny2GetUniqueWeaponHistory) | **Get** /Destiny2/{membershipType}/Account/{destinyMembershipId}/Character/{characterId}/Stats/UniqueWeapons/ | 
[**Destiny2GetVendor**](Destiny2Api.md#Destiny2GetVendor) | **Get** /Destiny2/{membershipType}/Profile/{destinyMembershipId}/Character/{characterId}/Vendors/{vendorHash}/ | 
[**Destiny2GetVendors**](Destiny2Api.md#Destiny2GetVendors) | **Get** /Destiny2/{membershipType}/Profile/{destinyMembershipId}/Character/{characterId}/Vendors/ | 
[**Destiny2InsertSocketPlug**](Destiny2Api.md#Destiny2InsertSocketPlug) | **Post** /Destiny2/Actions/Items/InsertSocketPlug/ | 
[**Destiny2PullFromPostmaster**](Destiny2Api.md#Destiny2PullFromPostmaster) | **Post** /Destiny2/Actions/Items/PullFromPostmaster/ | 
[**Destiny2ReportOffensivePostGameCarnageReportPlayer**](Destiny2Api.md#Destiny2ReportOffensivePostGameCarnageReportPlayer) | **Post** /Destiny2/Stats/PostGameCarnageReport/{activityId}/Report/ | 
[**Destiny2SearchDestinyEntities**](Destiny2Api.md#Destiny2SearchDestinyEntities) | **Get** /Destiny2/Armory/Search/{type}/{searchTerm}/ | 
[**Destiny2SearchDestinyPlayer**](Destiny2Api.md#Destiny2SearchDestinyPlayer) | **Get** /Destiny2/SearchDestinyPlayer/{membershipType}/{displayName}/ | 
[**Destiny2SetItemLockState**](Destiny2Api.md#Destiny2SetItemLockState) | **Post** /Destiny2/Actions/Items/SetLockState/ | 
[**Destiny2TransferItem**](Destiny2Api.md#Destiny2TransferItem) | **Post** /Destiny2/Actions/Items/TransferItem/ | 


# **Destiny2AwaGetActionToken**
> InlineResponse20059 Destiny2AwaGetActionToken(ctx, correlationId)


Returns the action token if user approves the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **correlationId** | **string**| The identifier for the advanced write action request. | 

### Return type

[**InlineResponse20059**](inline_response_200_59.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2AwaInitializeRequest**
> InlineResponse20058 Destiny2AwaInitializeRequest(ctx, destinyAdvancedAwaPermissionRequested)


Initialize a request to perform an advanced write action.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **destinyAdvancedAwaPermissionRequested** | [**DestinyAdvancedAwaPermissionRequested**](DestinyAdvancedAwaPermissionRequested.md)|  | 

### Return type

[**InlineResponse20058**](inline_response_200_58.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2AwaProvideAuthorizationResult**
> InlineResponse20019 Destiny2AwaProvideAuthorizationResult(ctx, destinyAdvancedAwaUserResponse)


Provide the result of the user interaction. Called by the Bungie Destiny App to approve or reject a request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **destinyAdvancedAwaUserResponse** | [**DestinyAdvancedAwaUserResponse**](DestinyAdvancedAwaUserResponse.md)|  | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2EquipItem**
> InlineResponse20019 Destiny2EquipItem(ctx, destinyRequestsActionsDestinyItemActionRequest)


Equip an item. You must have a valid Destiny Account, and either be in a social space, in orbit, or offline.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **destinyRequestsActionsDestinyItemActionRequest** | [**DestinyRequestsActionsDestinyItemActionRequest**](DestinyRequestsActionsDestinyItemActionRequest.md)|  | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2EquipItems**
> InlineResponse20044 Destiny2EquipItems(ctx, destinyRequestsActionsDestinyItemSetActionRequest)


Equip a list of items by itemInstanceIds. You must have a valid Destiny Account, and either be in a social space, in orbit, or offline. Any items not found on your character will be ignored.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **destinyRequestsActionsDestinyItemSetActionRequest** | [**DestinyRequestsActionsDestinyItemSetActionRequest**](DestinyRequestsActionsDestinyItemSetActionRequest.md)|  | 

### Return type

[**InlineResponse20044**](inline_response_200_44.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetActivityHistory**
> InlineResponse20053 Destiny2GetActivityHistory(ctx, characterId, destinyMembershipId, membershipType, optional)


Gets activity history stats for indicated character.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int64**| The id of the character to retrieve. | 
  **destinyMembershipId** | **int64**| The Destiny membershipId of the user to retrieve. | 
  **membershipType** | [**BungieMembershipType**](.md)| A valid non-BungieNet membership type. | 
 **optional** | ***Destiny2GetActivityHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Destiny2GetActivityHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **count** | **optional.Int32**| Number of rows to return | 
 **mode** | [**optional.Interface of DestinyHistoricalStatsDefinitionsDestinyActivityModeType**](.md)| A filter for the activity mode to be returned. None returns all activities. See the documentation for DestinyActivityModeType for valid values, and pass in string representation. | 
 **page** | **optional.Int32**| Page number to return, starting with 0. | 

### Return type

[**InlineResponse20053**](inline_response_200_53.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetCharacter**
> InlineResponse20038 Destiny2GetCharacter(ctx, characterId, destinyMembershipId, membershipType, optional)


Returns character information for the supplied character.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int64**| ID of the character. | 
  **destinyMembershipId** | **int64**| Destiny membership ID. | 
  **membershipType** | [**BungieMembershipType**](.md)| A valid non-BungieNet membership type. | 
 **optional** | ***Destiny2GetCharacterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Destiny2GetCharacterOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **components** | [**optional.Interface of []DestinyDestinyComponentType**](DestinyDestinyComponentType.md)| A comma separated list of components to return (as strings or numeric values). See the DestinyComponentType enum for valid components to request. You must request at least one component to receive results. | 

### Return type

[**InlineResponse20038**](inline_response_200_38.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetClanAggregateStats**
> InlineResponse20049 Destiny2GetClanAggregateStats(ctx, groupId, optional)


Gets aggregated stats for a clan using the same categories as the clan leaderboards. PREVIEW: This endpoint is still in beta, and may experience rough edges. The schema is in final form, but there may be bugs that prevent desirable operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Group ID of the clan whose leaderboards you wish to fetch. | 
 **optional** | ***Destiny2GetClanAggregateStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Destiny2GetClanAggregateStatsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modes** | **optional.String**| List of game modes for which to get leaderboards. See the documentation for DestinyActivityModeType for valid values, and pass in string representation, comma delimited. | 

### Return type

[**InlineResponse20049**](inline_response_200_49.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetClanLeaderboards**
> InlineResponse20048 Destiny2GetClanLeaderboards(ctx, groupId, optional)


Gets leaderboards with the signed in user's friends and the supplied destinyMembershipId as the focus. PREVIEW: This endpoint is still in beta, and may experience rough edges. The schema is in final form, but there may be bugs that prevent desirable operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Group ID of the clan whose leaderboards you wish to fetch. | 
 **optional** | ***Destiny2GetClanLeaderboardsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Destiny2GetClanLeaderboardsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maxtop** | **optional.Int32**| Maximum number of top players to return. Use a large number to get entire leaderboard. | 
 **modes** | **optional.String**| List of game modes for which to get leaderboards. See the documentation for DestinyActivityModeType for valid values, and pass in string representation, comma delimited. | 
 **statid** | **optional.String**| ID of stat to return rather than returning all Leaderboard stats. | 

### Return type

[**InlineResponse20048**](inline_response_200_48.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetClanWeeklyRewardState**
> InlineResponse20039 Destiny2GetClanWeeklyRewardState(ctx, groupId)


Returns information on the weekly clan rewards and if the clan has earned them or not. Note that this will always report rewards as not redeemed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| A valid group id of clan. | 

### Return type

[**InlineResponse20039**](inline_response_200_39.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetCollectibleNodeDetails**
> InlineResponse20043 Destiny2GetCollectibleNodeDetails(ctx, characterId, collectiblePresentationNodeHash, destinyMembershipId, membershipType, optional)


Given a Presentation Node that has Collectibles as direct descendants, this will return item details about those descendants in the context of the requesting character.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int64**| The Destiny Character ID of the character for whom we&#39;re getting collectible detail info. | 
  **collectiblePresentationNodeHash** | **int32**| The hash identifier of the Presentation Node for whom we should return collectible details. Details will only be returned for collectibles that are direct descendants of this node. | 
  **destinyMembershipId** | **int64**| Destiny membership ID of another user. You may be denied. | 
  **membershipType** | [**BungieMembershipType**](.md)| A valid non-BungieNet membership type. | 
 **optional** | ***Destiny2GetCollectibleNodeDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Destiny2GetCollectibleNodeDetailsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **components** | [**optional.Interface of []DestinyDestinyComponentType**](DestinyDestinyComponentType.md)| A comma separated list of components to return (as strings or numeric values). See the DestinyComponentType enum for valid components to request. You must request at least one component to receive results. | 

### Return type

[**InlineResponse20043**](inline_response_200_43.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetDestinyAggregateActivityStats**
> InlineResponse20055 Destiny2GetDestinyAggregateActivityStats(ctx, characterId, destinyMembershipId, membershipType)


Gets all activities the character has participated in together with aggregate statistics for those activities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int64**| The specific character whose activities should be returned. | 
  **destinyMembershipId** | **int64**| The Destiny membershipId of the user to retrieve. | 
  **membershipType** | [**BungieMembershipType**](.md)| A valid non-BungieNet membership type. | 

### Return type

[**InlineResponse20055**](inline_response_200_55.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetDestinyEntityDefinition**
> InlineResponse20034 Destiny2GetDestinyEntityDefinition(ctx, entityType, hashIdentifier)


Returns the static definition of an entity of the given Type and hash identifier. Examine the API Documentation for the Type Names of entities that have their own definitions. Note that the return type will always *inherit from* DestinyDefinition, but the specific type returned will be the requested entity type if it can be found. Please don't use this as a chatty alternative to the Manifest database if you require large sets of data, but for simple and one-off accesses this should be handy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entityType** | **string**| The type of entity for whom you would like results. These correspond to the entity&#39;s definition contract name. For instance, if you are looking for items, this property should be &#39;DestinyInventoryItemDefinition&#39;. PREVIEW: This endpoint is still in beta, and may experience rough edges. The schema is tentatively in final form, but there may be bugs that prevent desirable operation. | 
  **hashIdentifier** | **int32**| The hash identifier for the specific Entity you want returned. | 

### Return type

[**InlineResponse20034**](inline_response_200_34.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetDestinyManifest**
> InlineResponse20033 Destiny2GetDestinyManifest(ctx, )


Returns the current version of the manifest as a json object.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetHistoricalStats**
> InlineResponse20051 Destiny2GetHistoricalStats(ctx, characterId, destinyMembershipId, membershipType, optional)


Gets historical stats for indicated character.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int64**| The id of the character to retrieve. You can omit this character ID or set it to 0 to get aggregate stats across all characters. | 
  **destinyMembershipId** | **int64**| The Destiny membershipId of the user to retrieve. | 
  **membershipType** | [**BungieMembershipType**](.md)| A valid non-BungieNet membership type. | 
 **optional** | ***Destiny2GetHistoricalStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Destiny2GetHistoricalStatsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dayend** | **optional.Time**| Last day to return when daily stats are requested. Use the format YYYY-MM-DD. | 
 **daystart** | **optional.Time**| First day to return when daily stats are requested. Use the format YYYY-MM-DD | 
 **groups** | [**optional.Interface of []DestinyHistoricalStatsDefinitionsDestinyStatsGroupType**](DestinyHistoricalStatsDefinitionsDestinyStatsGroupType.md)| Group of stats to include, otherwise only general stats are returned. Comma separated list is allowed. Values: General, Weapons, Medals | 
 **modes** | [**optional.Interface of []DestinyHistoricalStatsDefinitionsDestinyActivityModeType**](DestinyHistoricalStatsDefinitionsDestinyActivityModeType.md)| Game modes to return. See the documentation for DestinyActivityModeType for valid values, and pass in string representation, comma delimited. | 
 **periodType** | [**optional.Interface of DestinyHistoricalStatsDefinitionsPeriodType**](.md)| Indicates a specific period type to return. Optional. May be: Daily, AllTime, or Activity | 

### Return type

[**InlineResponse20051**](inline_response_200_51.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetHistoricalStatsDefinition**
> InlineResponse20047 Destiny2GetHistoricalStatsDefinition(ctx, )


Gets historical stats definitions.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20047**](inline_response_200_47.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetHistoricalStatsForAccount**
> InlineResponse20052 Destiny2GetHistoricalStatsForAccount(ctx, destinyMembershipId, membershipType, optional)


Gets aggregate historical stats organized around each character for a given account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **destinyMembershipId** | **int64**| The Destiny membershipId of the user to retrieve. | 
  **membershipType** | [**BungieMembershipType**](.md)| A valid non-BungieNet membership type. | 
 **optional** | ***Destiny2GetHistoricalStatsForAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Destiny2GetHistoricalStatsForAccountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **groups** | [**optional.Interface of []DestinyHistoricalStatsDefinitionsDestinyStatsGroupType**](DestinyHistoricalStatsDefinitionsDestinyStatsGroupType.md)| Groups of stats to include, otherwise only general stats are returned. Comma separated list is allowed. Values: General, Weapons, Medals. | 

### Return type

[**InlineResponse20052**](inline_response_200_52.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetItem**
> InlineResponse20040 Destiny2GetItem(ctx, destinyMembershipId, itemInstanceId, membershipType, optional)


Retrieve the details of an instanced Destiny Item. An instanced Destiny item is one with an ItemInstanceId. Non-instanced items, such as materials, have no useful instance-specific details and thus are not queryable here.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **destinyMembershipId** | **int64**| The membership ID of the destiny profile. | 
  **itemInstanceId** | **int64**| The Instance ID of the destiny item. | 
  **membershipType** | [**BungieMembershipType**](.md)| A valid non-BungieNet membership type. | 
 **optional** | ***Destiny2GetItemOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Destiny2GetItemOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **components** | [**optional.Interface of []DestinyDestinyComponentType**](DestinyDestinyComponentType.md)| A comma separated list of components to return (as strings or numeric values). See the DestinyComponentType enum for valid components to request. You must request at least one component to receive results. | 

### Return type

[**InlineResponse20040**](inline_response_200_40.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetLeaderboards**
> InlineResponse20048 Destiny2GetLeaderboards(ctx, destinyMembershipId, membershipType, optional)


Gets leaderboards with the signed in user's friends and the supplied destinyMembershipId as the focus. PREVIEW: This endpoint has not yet been implemented. It is being returned for a preview of future functionality, and for public comment/suggestion/preparation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **destinyMembershipId** | **int64**| The Destiny membershipId of the user to retrieve. | 
  **membershipType** | [**BungieMembershipType**](.md)| A valid non-BungieNet membership type. | 
 **optional** | ***Destiny2GetLeaderboardsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Destiny2GetLeaderboardsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **maxtop** | **optional.Int32**| Maximum number of top players to return. Use a large number to get entire leaderboard. | 
 **modes** | **optional.String**| List of game modes for which to get leaderboards. See the documentation for DestinyActivityModeType for valid values, and pass in string representation, comma delimited. | 
 **statid** | **optional.String**| ID of stat to return rather than returning all Leaderboard stats. | 

### Return type

[**InlineResponse20048**](inline_response_200_48.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetLeaderboardsForCharacter**
> InlineResponse20048 Destiny2GetLeaderboardsForCharacter(ctx, characterId, destinyMembershipId, membershipType, optional)


Gets leaderboards with the signed in user's friends and the supplied destinyMembershipId as the focus. PREVIEW: This endpoint is still in beta, and may experience rough edges. The schema is in final form, but there may be bugs that prevent desirable operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int64**| The specific character to build the leaderboard around for the provided Destiny Membership. | 
  **destinyMembershipId** | **int64**| The Destiny membershipId of the user to retrieve. | 
  **membershipType** | [**BungieMembershipType**](.md)| A valid non-BungieNet membership type. | 
 **optional** | ***Destiny2GetLeaderboardsForCharacterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Destiny2GetLeaderboardsForCharacterOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **maxtop** | **optional.Int32**| Maximum number of top players to return. Use a large number to get entire leaderboard. | 
 **modes** | **optional.String**| List of game modes for which to get leaderboards. See the documentation for DestinyActivityModeType for valid values, and pass in string representation, comma delimited. | 
 **statid** | **optional.String**| ID of stat to return rather than returning all Leaderboard stats. | 

### Return type

[**InlineResponse20048**](inline_response_200_48.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetLinkedProfiles**
> InlineResponse20036 Destiny2GetLinkedProfiles(ctx, membershipId, membershipType)


Returns a summary information about all profiles linked to the requesting membership type/membership ID that have valid Destiny information. The passed-in Membership Type/Membership ID may be a Bungie.Net membership or a Destiny membership. It only returns the minimal amount of data to begin making more substantive requests, but will hopefully serve as a useful alternative to UserServices for people who just care about Destiny data. Note that it will only return linked accounts whose linkages you are allowed to view.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **membershipId** | **int64**| The ID of the membership whose linked Destiny accounts you want returned. Make sure your membership ID matches its Membership Type: don&#39;t pass us a PSN membership ID and the XBox membership type, it&#39;s not going to work! | 
  **membershipType** | [**BungieMembershipType**](.md)| The type for the membership whose linked Destiny accounts you want returned. | 

### Return type

[**InlineResponse20036**](inline_response_200_36.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetPostGameCarnageReport**
> InlineResponse20046 Destiny2GetPostGameCarnageReport(ctx, activityId)


Gets the available post game carnage report for the activity ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activityId** | **int64**| The ID of the activity whose PGCR is requested. | 

### Return type

[**InlineResponse20046**](inline_response_200_46.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetProfile**
> InlineResponse20037 Destiny2GetProfile(ctx, destinyMembershipId, membershipType, optional)


Returns Destiny Profile information for the supplied membership.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **destinyMembershipId** | **int64**| Destiny membership ID. | 
  **membershipType** | [**BungieMembershipType**](.md)| A valid non-BungieNet membership type. | 
 **optional** | ***Destiny2GetProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Destiny2GetProfileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **components** | [**optional.Interface of []DestinyDestinyComponentType**](DestinyDestinyComponentType.md)| A comma separated list of components to return (as strings or numeric values). See the DestinyComponentType enum for valid components to request. You must request at least one component to receive results. | 

### Return type

[**InlineResponse20037**](inline_response_200_37.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetPublicMilestoneContent**
> InlineResponse20056 Destiny2GetPublicMilestoneContent(ctx, milestoneHash)


Gets custom localized content for the milestone of the given hash, if it exists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **milestoneHash** | **int32**| The identifier for the milestone to be returned. | 

### Return type

[**InlineResponse20056**](inline_response_200_56.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetPublicMilestones**
> InlineResponse20057 Destiny2GetPublicMilestones(ctx, )


Gets public information about currently available Milestones.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20057**](inline_response_200_57.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetUniqueWeaponHistory**
> InlineResponse20054 Destiny2GetUniqueWeaponHistory(ctx, characterId, destinyMembershipId, membershipType)


Gets details about unique weapon usage, including all exotic weapons.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int64**| The id of the character to retrieve. | 
  **destinyMembershipId** | **int64**| The Destiny membershipId of the user to retrieve. | 
  **membershipType** | [**BungieMembershipType**](.md)| A valid non-BungieNet membership type. | 

### Return type

[**InlineResponse20054**](inline_response_200_54.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetVendor**
> InlineResponse20042 Destiny2GetVendor(ctx, characterId, destinyMembershipId, membershipType, vendorHash, optional)


Get the details of a specific Vendor.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int64**| The Destiny Character ID of the character for whom we&#39;re getting vendor info. | 
  **destinyMembershipId** | **int64**| Destiny membership ID of another user. You may be denied. | 
  **membershipType** | [**BungieMembershipType**](.md)| A valid non-BungieNet membership type. | 
  **vendorHash** | **int32**| The Hash identifier of the Vendor to be returned. | 
 **optional** | ***Destiny2GetVendorOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Destiny2GetVendorOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **components** | [**optional.Interface of []DestinyDestinyComponentType**](DestinyDestinyComponentType.md)| A comma separated list of components to return (as strings or numeric values). See the DestinyComponentType enum for valid components to request. You must request at least one component to receive results. | 

### Return type

[**InlineResponse20042**](inline_response_200_42.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2GetVendors**
> InlineResponse20041 Destiny2GetVendors(ctx, characterId, destinyMembershipId, membershipType, optional)


Get currently available vendors from the list of vendors that can possibly have rotating inventory. Note that this does not include things like preview vendors and vendors-as-kiosks, neither of whom have rotating/dynamic inventories. Use their definitions as-is for those.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **characterId** | **int64**| The Destiny Character ID of the character for whom we&#39;re getting vendor info. | 
  **destinyMembershipId** | **int64**| Destiny membership ID of another user. You may be denied. | 
  **membershipType** | [**BungieMembershipType**](.md)| A valid non-BungieNet membership type. | 
 **optional** | ***Destiny2GetVendorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Destiny2GetVendorsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **components** | [**optional.Interface of []DestinyDestinyComponentType**](DestinyDestinyComponentType.md)| A comma separated list of components to return (as strings or numeric values). See the DestinyComponentType enum for valid components to request. You must request at least one component to receive results. | 

### Return type

[**InlineResponse20041**](inline_response_200_41.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2InsertSocketPlug**
> InlineResponse20045 Destiny2InsertSocketPlug(ctx, destinyRequestsActionsDestinyInsertPlugsActionRequest)


Insert a plug into a socketed item. I know how it sounds, but I assure you it's much more G-rated than you might be guessing. We haven't decided yet whether this will be able to insert plugs that have side effects, but if we do it will require special scope permission for an application attempting to do so. You must have a valid Destiny Account, and either be in a social space, in orbit, or offline. Request must include proof of permission for 'InsertPlugs' from the account owner.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **destinyRequestsActionsDestinyInsertPlugsActionRequest** | [**DestinyRequestsActionsDestinyInsertPlugsActionRequest**](DestinyRequestsActionsDestinyInsertPlugsActionRequest.md)|  | 

### Return type

[**InlineResponse20045**](inline_response_200_45.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2PullFromPostmaster**
> InlineResponse20019 Destiny2PullFromPostmaster(ctx, destinyRequestsActionsDestinyPostmasterTransferRequest)


Extract an item from the Postmaster, with whatever implications that may entail. You must have a valid Destiny account. You must also pass BOTH a reference AND an instance ID if it's an instanced item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **destinyRequestsActionsDestinyPostmasterTransferRequest** | [**DestinyRequestsActionsDestinyPostmasterTransferRequest**](DestinyRequestsActionsDestinyPostmasterTransferRequest.md)|  | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2ReportOffensivePostGameCarnageReportPlayer**
> InlineResponse20019 Destiny2ReportOffensivePostGameCarnageReportPlayer(ctx, activityId, destinyReportingRequestsDestinyReportOffensePgcrRequest)


Report a player that you met in an activity that was engaging in ToS-violating activities. Both you and the offending player must have played in the activityId passed in. Please use this judiciously and only when you have strong suspicions of violation, pretty please.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activityId** | **int64**| The ID of the activity where you ran into the brigand that you&#39;re reporting. | 
  **destinyReportingRequestsDestinyReportOffensePgcrRequest** | [**DestinyReportingRequestsDestinyReportOffensePgcrRequest**](DestinyReportingRequestsDestinyReportOffensePgcrRequest.md)|  | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2SearchDestinyEntities**
> InlineResponse20050 Destiny2SearchDestinyEntities(ctx, searchTerm, type_, optional)


Gets a page list of Destiny items.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **searchTerm** | **string**| The string to use when searching for Destiny entities. | 
  **type_** | **string**| The type of entity for whom you would like results. These correspond to the entity&#39;s definition contract name. For instance, if you are looking for items, this property should be &#39;DestinyInventoryItemDefinition&#39;. | 
 **optional** | ***Destiny2SearchDestinyEntitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Destiny2SearchDestinyEntitiesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| Page number to return, starting with 0. | 

### Return type

[**InlineResponse20050**](inline_response_200_50.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2SearchDestinyPlayer**
> InlineResponse20035 Destiny2SearchDestinyPlayer(ctx, displayName, membershipType)


Returns a list of Destiny memberships given a full Gamertag or PSN ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **displayName** | **string**| The full gamertag or PSN id of the player. Spaces and case are ignored. | 
  **membershipType** | [**BungieMembershipType**](.md)| A valid non-BungieNet membership type, or All. | 

### Return type

[**InlineResponse20035**](inline_response_200_35.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2SetItemLockState**
> InlineResponse20019 Destiny2SetItemLockState(ctx, destinyRequestsActionsDestinyItemStateRequest)


Set the Lock State for an instanced item. You must have a valid Destiny Account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **destinyRequestsActionsDestinyItemStateRequest** | [**DestinyRequestsActionsDestinyItemStateRequest**](DestinyRequestsActionsDestinyItemStateRequest.md)|  | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destiny2TransferItem**
> InlineResponse20019 Destiny2TransferItem(ctx, destinyRequestsDestinyItemTransferRequest)


Transfer an item to/from your vault. You must have a valid Destiny account. You must also pass BOTH a reference AND an instance ID if it's an instanced item. itshappening.gif

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **destinyRequestsDestinyItemTransferRequest** | [**DestinyRequestsDestinyItemTransferRequest**](DestinyRequestsDestinyItemTransferRequest.md)|  | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

