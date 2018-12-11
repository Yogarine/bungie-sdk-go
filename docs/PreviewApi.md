# Bungie\PreviewApi

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Destiny2GetClanAggregateStats**](PreviewApi.md#Destiny2GetClanAggregateStats) | **Get** /Destiny2/Stats/AggregateClanStats/{groupId}/ | 
[**Destiny2GetClanLeaderboards**](PreviewApi.md#Destiny2GetClanLeaderboards) | **Get** /Destiny2/Stats/Leaderboards/Clans/{groupId}/ | 
[**Destiny2GetLeaderboards**](PreviewApi.md#Destiny2GetLeaderboards) | **Get** /Destiny2/{membershipType}/Account/{destinyMembershipId}/Stats/Leaderboards/ | 
[**Destiny2GetLeaderboardsForCharacter**](PreviewApi.md#Destiny2GetLeaderboardsForCharacter) | **Get** /Destiny2/Stats/Leaderboards/{membershipType}/{destinyMembershipId}/{characterId}/ | 
[**Destiny2InsertSocketPlug**](PreviewApi.md#Destiny2InsertSocketPlug) | **Post** /Destiny2/Actions/Items/InsertSocketPlug/ | 


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

