# Bungie\FireteamApi

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FireteamGetActivePrivateClanFireteamCount**](FireteamApi.md#FireteamGetActivePrivateClanFireteamCount) | **Get** /Fireteam/Clan/{groupId}/ActiveCount/ | 
[**FireteamGetAvailableClanFireteams**](FireteamApi.md#FireteamGetAvailableClanFireteams) | **Get** /Fireteam/Clan/{groupId}/Available/{platform}/{activityType}/{dateRange}/{slotFilter}/{publicOnly}/{page}/ | 
[**FireteamGetClanFireteam**](FireteamApi.md#FireteamGetClanFireteam) | **Get** /Fireteam/Clan/{groupId}/Summary/{fireteamId}/ | 
[**FireteamGetMyClanFireteams**](FireteamApi.md#FireteamGetMyClanFireteams) | **Get** /Fireteam/Clan/{groupId}/My/{platform}/{includeClosed}/{page}/ | 
[**FireteamSearchPublicAvailableClanFireteams**](FireteamApi.md#FireteamSearchPublicAvailableClanFireteams) | **Get** /Fireteam/Search/Available/{platform}/{activityType}/{dateRange}/{slotFilter}/{page}/ | 


# **FireteamGetActivePrivateClanFireteamCount**
> InlineResponse20019 FireteamGetActivePrivateClanFireteamCount(ctx, groupId)


Gets a count of all active non-public fireteams for the specified clan. Maximum value returned is 25.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| The group id of the clan. | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FireteamGetAvailableClanFireteams**
> InlineResponse20065 FireteamGetAvailableClanFireteams(ctx, activityType, dateRange, groupId, page, platform, publicOnly, slotFilter, optional)


Gets a listing of all of this clan's fireteams that are have available slots. Caller is not checked for join criteria so caching is maximized.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activityType** | [**FireteamFireteamActivityType**](.md)| The activity type to filter by. | 
  **dateRange** | [**FireteamFireteamDateRange**](.md)| The date range to grab available fireteams. | 
  **groupId** | **int64**| The group id of the clan. | 
  **page** | **int32**| Zero based page | 
  **platform** | [**FireteamFireteamPlatform**](.md)| The platform filter. | 
  **publicOnly** | [**FireteamFireteamPublicSearchOption**](.md)| Determines public/private filtering. | 
  **slotFilter** | [**FireteamFireteamSlotSearch**](.md)| Filters based on available slots | 
 **optional** | ***FireteamGetAvailableClanFireteamsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FireteamGetAvailableClanFireteamsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------







 **langFilter** | **optional.String**| An optional language filter. | 

### Return type

[**InlineResponse20065**](inline_response_200_65.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FireteamGetClanFireteam**
> InlineResponse20067 FireteamGetClanFireteam(ctx, fireteamId, groupId)


Gets a specific clan fireteam.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fireteamId** | **int64**| The unique id of the fireteam. | 
  **groupId** | **int64**| The group id of the clan. | 

### Return type

[**InlineResponse20067**](inline_response_200_67.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FireteamGetMyClanFireteams**
> InlineResponse20066 FireteamGetMyClanFireteams(ctx, groupId, includeClosed, page, platform, optional)


Gets a listing of all clan fireteams that caller is an applicant, a member, or an alternate of.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| The group id of the clan. (This parameter is ignored unless the optional query parameter groupFilter is true). | 
  **includeClosed** | **bool**| If true, return fireteams that have been closed. | 
  **page** | **int32**| Deprecated parameter, ignored. | 
  **platform** | [**FireteamFireteamPlatform**](.md)| The platform filter. | 
 **optional** | ***FireteamGetMyClanFireteamsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FireteamGetMyClanFireteamsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **groupFilter** | **optional.Bool**| If true, filter by clan. Otherwise, ignore the clan and show all of the user&#39;s fireteams. | 
 **langFilter** | **optional.String**| An optional language filter. | 

### Return type

[**InlineResponse20066**](inline_response_200_66.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FireteamSearchPublicAvailableClanFireteams**
> InlineResponse20065 FireteamSearchPublicAvailableClanFireteams(ctx, activityType, dateRange, page, platform, slotFilter, optional)


Gets a listing of all public fireteams starting now with open slots. Caller is not checked for join criteria so caching is maximized.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **activityType** | [**FireteamFireteamActivityType**](.md)| The activity type to filter by. | 
  **dateRange** | [**FireteamFireteamDateRange**](.md)| The date range to grab available fireteams. | 
  **page** | **int32**| Zero based page | 
  **platform** | [**FireteamFireteamPlatform**](.md)| The platform filter. | 
  **slotFilter** | [**FireteamFireteamSlotSearch**](.md)| Filters based on available slots | 
 **optional** | ***FireteamSearchPublicAvailableClanFireteamsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FireteamSearchPublicAvailableClanFireteamsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **langFilter** | **optional.String**| An optional language filter. | 

### Return type

[**InlineResponse20065**](inline_response_200_65.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

