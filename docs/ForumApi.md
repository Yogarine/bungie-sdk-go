# Bungie\ForumApi

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ForumApproveFireteamThread**](ForumApi.md#ForumApproveFireteamThread) | **Post** /Forum/Recruit/Approve/{topicId}/ | 
[**ForumGetCoreTopicsPaged**](ForumApi.md#ForumGetCoreTopicsPaged) | **Get** /Forum/GetCoreTopicsPaged/{page}/{sort}/{quickDate}/{categoryFilter}/ | 
[**ForumGetForumTagSuggestions**](ForumApi.md#ForumGetForumTagSuggestions) | **Get** /Forum/GetForumTagSuggestions/ | 
[**ForumGetPoll**](ForumApi.md#ForumGetPoll) | **Get** /Forum/Poll/{topicId}/ | 
[**ForumGetPostAndParent**](ForumApi.md#ForumGetPostAndParent) | **Get** /Forum/GetPostAndParent/{childPostId}/ | 
[**ForumGetPostAndParentAwaitingApproval**](ForumApi.md#ForumGetPostAndParentAwaitingApproval) | **Get** /Forum/GetPostAndParentAwaitingApproval/{childPostId}/ | 
[**ForumGetPostsThreadedPaged**](ForumApi.md#ForumGetPostsThreadedPaged) | **Get** /Forum/GetPostsThreadedPaged/{parentPostId}/{page}/{pageSize}/{replySize}/{getParentPost}/{rootThreadMode}/{sortMode}/ | 
[**ForumGetPostsThreadedPagedFromChild**](ForumApi.md#ForumGetPostsThreadedPagedFromChild) | **Get** /Forum/GetPostsThreadedPagedFromChild/{childPostId}/{page}/{pageSize}/{replySize}/{rootThreadMode}/{sortMode}/ | 
[**ForumGetRecruitmentThreadSummaries**](ForumApi.md#ForumGetRecruitmentThreadSummaries) | **Post** /Forum/Recruit/Summaries/ | 
[**ForumGetTopicForContent**](ForumApi.md#ForumGetTopicForContent) | **Get** /Forum/GetTopicForContent/{contentId}/ | 
[**ForumGetTopicsPaged**](ForumApi.md#ForumGetTopicsPaged) | **Get** /Forum/GetTopicsPaged/{page}/{pageSize}/{group}/{sort}/{quickDate}/{categoryFilter}/ | 
[**ForumJoinFireteamThread**](ForumApi.md#ForumJoinFireteamThread) | **Post** /Forum/Recruit/Join/{topicId}/ | 
[**ForumKickBanFireteamApplicant**](ForumApi.md#ForumKickBanFireteamApplicant) | **Post** /Forum/Recruit/KickBan/{topicId}/{targetMembershipId}/ | 
[**ForumLeaveFireteamThread**](ForumApi.md#ForumLeaveFireteamThread) | **Post** /Forum/Recruit/Leave/{topicId}/ | 


# **ForumApproveFireteamThread**
> InlineResponse20014 ForumApproveFireteamThread(ctx, topicId)


Allows the owner of a fireteam thread to approve all joined members and start a private message conversation with them.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **topicId** | **int64**| The post id of the recruitment topic to approve. | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetCoreTopicsPaged**
> InlineResponse20010 ForumGetCoreTopicsPaged(ctx, categoryFilter, page, quickDate, sort, optional)


Gets a listing of all topics marked as part of the core group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **categoryFilter** | [**ForumForumTopicsCategoryFiltersEnum**](.md)| The category filter. | 
  **page** | **int32**| Zero base page | 
  **quickDate** | [**ForumForumTopicsQuickDateEnum**](.md)| The date filter. | 
  **sort** | [**ForumForumTopicsSortEnum**](.md)| The sort mode. | 
 **optional** | ***ForumGetCoreTopicsPagedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForumGetCoreTopicsPagedOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **locales** | **optional.String**| Comma seperated list of locales posts must match to return in the result list. Default &#39;en&#39; | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetForumTagSuggestions**
> InlineResponse20012 ForumGetForumTagSuggestions(ctx, optional)


Gets tag suggestions based on partial text entry, matching them with other tags previously used in the forums.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ForumGetForumTagSuggestionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForumGetForumTagSuggestionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **partialtag** | **optional.String**| The partial tag input to generate suggestions from. | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetPoll**
> InlineResponse20010 ForumGetPoll(ctx, topicId)


Gets the specified forum poll.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **topicId** | **int64**| The post id of the topic that has the poll. | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetPostAndParent**
> InlineResponse20010 ForumGetPostAndParent(ctx, childPostId, optional)


Returns the post specified and its immediate parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **childPostId** | **int64**|  | 
 **optional** | ***ForumGetPostAndParentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForumGetPostAndParentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showbanned** | **optional.String**| If this value is not null or empty, banned posts are requested to be returned | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetPostAndParentAwaitingApproval**
> InlineResponse20010 ForumGetPostAndParentAwaitingApproval(ctx, childPostId, optional)


Returns the post specified and its immediate parent of posts that are awaiting approval.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **childPostId** | **int64**|  | 
 **optional** | ***ForumGetPostAndParentAwaitingApprovalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForumGetPostAndParentAwaitingApprovalOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showbanned** | **optional.String**| If this value is not null or empty, banned posts are requested to be returned | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetPostsThreadedPaged**
> InlineResponse20010 ForumGetPostsThreadedPaged(ctx, getParentPost, page, pageSize, parentPostId, replySize, rootThreadMode, sortMode, optional)


Returns a thread of posts at the given parent, optionally returning replies to those posts as well as the original parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **getParentPost** | **bool**|  | 
  **page** | **int32**|  | 
  **pageSize** | **int32**|  | 
  **parentPostId** | **int64**|  | 
  **replySize** | **int32**|  | 
  **rootThreadMode** | **bool**|  | 
  **sortMode** | [**ForumForumPostSortEnum**](.md)|  | 
 **optional** | ***ForumGetPostsThreadedPagedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForumGetPostsThreadedPagedOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------







 **showbanned** | **optional.String**| If this value is not null or empty, banned posts are requested to be returned | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetPostsThreadedPagedFromChild**
> InlineResponse20010 ForumGetPostsThreadedPagedFromChild(ctx, childPostId, page, pageSize, replySize, rootThreadMode, sortMode, optional)


Returns a thread of posts starting at the topicId of the input childPostId, optionally returning replies to those posts as well as the original parent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **childPostId** | **int64**|  | 
  **page** | **int32**|  | 
  **pageSize** | **int32**|  | 
  **replySize** | **int32**|  | 
  **rootThreadMode** | **bool**|  | 
  **sortMode** | [**ForumForumPostSortEnum**](.md)|  | 
 **optional** | ***ForumGetPostsThreadedPagedFromChildOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForumGetPostsThreadedPagedFromChildOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **showbanned** | **optional.String**| If this value is not null or empty, banned posts are requested to be returned | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetRecruitmentThreadSummaries**
> InlineResponse20015 ForumGetRecruitmentThreadSummaries(ctx, requestBody)


Allows the caller to get a list of to 25 recruitment thread summary information objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestBody** | [**[]int64**](array.md)|  | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetTopicForContent**
> InlineResponse20011 ForumGetTopicForContent(ctx, contentId)


Gets the post Id for the given content item's comments, if it exists.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contentId** | **int64**|  | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumGetTopicsPaged**
> InlineResponse20010 ForumGetTopicsPaged(ctx, categoryFilter, group, page, pageSize, quickDate, sort, optional)


Get topics from any forum.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **categoryFilter** | [**ForumForumTopicsCategoryFiltersEnum**](.md)| A category filter | 
  **group** | **int64**| The group, if any. | 
  **page** | **int32**| Zero paged page number | 
  **pageSize** | **int32**| Unused | 
  **quickDate** | [**ForumForumTopicsQuickDateEnum**](.md)| A date filter. | 
  **sort** | [**ForumForumTopicsSortEnum**](.md)| The sort mode. | 
 **optional** | ***ForumGetTopicsPagedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ForumGetTopicsPagedOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **locales** | **optional.String**| Comma seperated list of locales posts must match to return in the result list. Default &#39;en&#39; | 
 **tagstring** | **optional.String**| The tags to search, if any. | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumJoinFireteamThread**
> InlineResponse20013 ForumJoinFireteamThread(ctx, topicId)


Allows a user to slot themselves into a recruitment thread fireteam slot. Returns the new state of the fireteam.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **topicId** | **int64**| The post id of the recruitment topic you wish to join. | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumKickBanFireteamApplicant**
> InlineResponse20013 ForumKickBanFireteamApplicant(ctx, targetMembershipId, topicId)


Allows a recruitment thread owner to kick a join user from the fireteam. Returns the new state of the fireteam.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **targetMembershipId** | **int64**| The id of the user you wish to kick. | 
  **topicId** | **int64**| The post id of the recruitment topic you wish to join. | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ForumLeaveFireteamThread**
> InlineResponse20013 ForumLeaveFireteamThread(ctx, topicId)


Allows a user to remove themselves from a recruitment thread fireteam slot. Returns the new state of the fireteam.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **topicId** | **int64**| The post id of the recruitment topic you wish to leave. | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

