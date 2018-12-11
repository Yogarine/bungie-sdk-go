# Bungie\CommunityContentApi

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommunityContentGetCommunityContent**](CommunityContentApi.md#CommunityContentGetCommunityContent) | **Get** /CommunityContent/Get/{sort}/{mediaFilter}/{page}/ | 
[**CommunityContentGetCommunityLiveStatuses**](CommunityContentApi.md#CommunityContentGetCommunityLiveStatuses) | **Get** /CommunityContent/Live/All/{partnershipType}/{sort}/{page}/ | 
[**CommunityContentGetCommunityLiveStatusesForClanmates**](CommunityContentApi.md#CommunityContentGetCommunityLiveStatusesForClanmates) | **Get** /CommunityContent/Live/Clan/{partnershipType}/{sort}/{page}/ | 
[**CommunityContentGetCommunityLiveStatusesForFriends**](CommunityContentApi.md#CommunityContentGetCommunityLiveStatusesForFriends) | **Get** /CommunityContent/Live/Friends/{partnershipType}/{sort}/{page}/ | 
[**CommunityContentGetFeaturedCommunityLiveStatuses**](CommunityContentApi.md#CommunityContentGetFeaturedCommunityLiveStatuses) | **Get** /CommunityContent/Live/Featured/{partnershipType}/{sort}/{page}/ | 
[**CommunityContentGetStreamingStatusForMember**](CommunityContentApi.md#CommunityContentGetStreamingStatusForMember) | **Get** /CommunityContent/Live/Users/{partnershipType}/{membershipType}/{membershipId}/ | 


# **CommunityContentGetCommunityContent**
> InlineResponse20010 CommunityContentGetCommunityContent(ctx, mediaFilter, page, sort)


Returns community content.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mediaFilter** | [**ForumForumTopicsCategoryFiltersEnum**](.md)| The type of media to get | 
  **page** | **int32**| Zero based page | 
  **sort** | [**ForumCommunityContentSortMode**](.md)| The sort mode. | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommunityContentGetCommunityLiveStatuses**
> InlineResponse20060 CommunityContentGetCommunityLiveStatuses(ctx, page, partnershipType, sort, optional)


Returns info about community members who are live streaming.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Zero based page. | 
  **partnershipType** | [**PartnershipsPartnershipType**](.md)| The type of partnership for which the status should be returned. | 
  **sort** | [**CommunityCommunityStatusSort**](.md)| The sort mode. | 
 **optional** | ***CommunityContentGetCommunityLiveStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommunityContentGetCommunityLiveStatusesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **modeHash** | **optional.Int32**| The hash of the Activity Mode for which streams should be retrieved. Don&#39;t pass it in or pass 0 to not filter by mode. | 
 **streamLocale** | **optional.String**| The locale for streams you&#39;d like to see. Don&#39;t pass this to fall back on your BNet locale. Pass &#39;ALL&#39; to not filter by locale. | 

### Return type

[**InlineResponse20060**](inline_response_200_60.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommunityContentGetCommunityLiveStatusesForClanmates**
> InlineResponse20060 CommunityContentGetCommunityLiveStatusesForClanmates(ctx, page, partnershipType, sort)


Returns info about community members who are live streaming in your clans.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Zero based page. | 
  **partnershipType** | [**PartnershipsPartnershipType**](.md)| The type of partnership for which the status should be returned. | 
  **sort** | [**CommunityCommunityStatusSort**](.md)| The sort mode. | 

### Return type

[**InlineResponse20060**](inline_response_200_60.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommunityContentGetCommunityLiveStatusesForFriends**
> InlineResponse20060 CommunityContentGetCommunityLiveStatusesForFriends(ctx, page, partnershipType, sort)


Returns info about community members who are live streaming among your friends.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Zero based page. | 
  **partnershipType** | [**PartnershipsPartnershipType**](.md)| The type of partnership for which the status should be returned. | 
  **sort** | [**CommunityCommunityStatusSort**](.md)| The sort mode. | 

### Return type

[**InlineResponse20060**](inline_response_200_60.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommunityContentGetFeaturedCommunityLiveStatuses**
> InlineResponse20060 CommunityContentGetFeaturedCommunityLiveStatuses(ctx, page, partnershipType, sort, optional)


Returns info about Featured live streams.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **page** | **int32**| Zero based page. | 
  **partnershipType** | [**PartnershipsPartnershipType**](.md)| The type of partnership for which the status should be returned. | 
  **sort** | [**CommunityCommunityStatusSort**](.md)| The sort mode. | 
 **optional** | ***CommunityContentGetFeaturedCommunityLiveStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CommunityContentGetFeaturedCommunityLiveStatusesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **streamLocale** | **optional.String**| The locale for streams you&#39;d like to see. Don&#39;t pass this to fall back on your BNet locale. Pass &#39;ALL&#39; to not filter by locale. | 

### Return type

[**InlineResponse20060**](inline_response_200_60.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CommunityContentGetStreamingStatusForMember**
> InlineResponse20061 CommunityContentGetStreamingStatusForMember(ctx, membershipId, membershipType, partnershipType)


Gets the Live Streaming status of a particular Account and Membership Type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **membershipId** | **int64**| The membershipId related to that type. | 
  **membershipType** | [**BungieMembershipType**](.md)| The type of account for which info will be extracted. | 
  **partnershipType** | [**PartnershipsPartnershipType**](.md)| The type of partnership for which info will be extracted. | 

### Return type

[**InlineResponse20061**](inline_response_200_61.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

