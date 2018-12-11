# Bungie\GroupV2Api

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupV2AbdicateFoundership**](GroupV2Api.md#GroupV2AbdicateFoundership) | **Post** /GroupV2/{groupId}/Admin/AbdicateFoundership/{membershipType}/{founderIdNew}/ | 
[**GroupV2AddOptionalConversation**](GroupV2Api.md#GroupV2AddOptionalConversation) | **Post** /GroupV2/{groupId}/OptionalConversations/Add/ | 
[**GroupV2ApproveAllPending**](GroupV2Api.md#GroupV2ApproveAllPending) | **Post** /GroupV2/{groupId}/Members/ApproveAll/ | 
[**GroupV2ApprovePending**](GroupV2Api.md#GroupV2ApprovePending) | **Post** /GroupV2/{groupId}/Members/Approve/{membershipType}/{membershipId}/ | 
[**GroupV2ApprovePendingForList**](GroupV2Api.md#GroupV2ApprovePendingForList) | **Post** /GroupV2/{groupId}/Members/ApproveList/ | 
[**GroupV2BanMember**](GroupV2Api.md#GroupV2BanMember) | **Post** /GroupV2/{groupId}/Members/{membershipType}/{membershipId}/Ban/ | 
[**GroupV2CreateGroup**](GroupV2Api.md#GroupV2CreateGroup) | **Post** /GroupV2/Create/ | 
[**GroupV2DenyAllPending**](GroupV2Api.md#GroupV2DenyAllPending) | **Post** /GroupV2/{groupId}/Members/DenyAll/ | 
[**GroupV2DenyPendingForList**](GroupV2Api.md#GroupV2DenyPendingForList) | **Post** /GroupV2/{groupId}/Members/DenyList/ | 
[**GroupV2EditClanBanner**](GroupV2Api.md#GroupV2EditClanBanner) | **Post** /GroupV2/{groupId}/EditClanBanner/ | 
[**GroupV2EditFounderOptions**](GroupV2Api.md#GroupV2EditFounderOptions) | **Post** /GroupV2/{groupId}/EditFounderOptions/ | 
[**GroupV2EditGroup**](GroupV2Api.md#GroupV2EditGroup) | **Post** /GroupV2/{groupId}/Edit/ | 
[**GroupV2EditGroupMembership**](GroupV2Api.md#GroupV2EditGroupMembership) | **Post** /GroupV2/{groupId}/Members/{membershipType}/{membershipId}/SetMembershipType/{memberType}/ | 
[**GroupV2EditOptionalConversation**](GroupV2Api.md#GroupV2EditOptionalConversation) | **Post** /GroupV2/{groupId}/OptionalConversations/Edit/{conversationId}/ | 
[**GroupV2GetAdminsAndFounderOfGroup**](GroupV2Api.md#GroupV2GetAdminsAndFounderOfGroup) | **Get** /GroupV2/{groupId}/AdminsAndFounder/ | 
[**GroupV2GetAvailableAvatars**](GroupV2Api.md#GroupV2GetAvailableAvatars) | **Get** /GroupV2/GetAvailableAvatars/ | 
[**GroupV2GetAvailableThemes**](GroupV2Api.md#GroupV2GetAvailableThemes) | **Get** /GroupV2/GetAvailableThemes/ | 
[**GroupV2GetBannedMembersOfGroup**](GroupV2Api.md#GroupV2GetBannedMembersOfGroup) | **Get** /GroupV2/{groupId}/Banned/ | 
[**GroupV2GetGroup**](GroupV2Api.md#GroupV2GetGroup) | **Get** /GroupV2/{groupId}/ | 
[**GroupV2GetGroupByName**](GroupV2Api.md#GroupV2GetGroupByName) | **Get** /GroupV2/Name/{groupName}/{groupType}/ | 
[**GroupV2GetGroupByNameV2**](GroupV2Api.md#GroupV2GetGroupByNameV2) | **Post** /GroupV2/NameV2/ | 
[**GroupV2GetGroupOptionalConversations**](GroupV2Api.md#GroupV2GetGroupOptionalConversations) | **Get** /GroupV2/{groupId}/OptionalConversations/ | 
[**GroupV2GetGroupsForMember**](GroupV2Api.md#GroupV2GetGroupsForMember) | **Get** /GroupV2/User/{membershipType}/{membershipId}/{filter}/{groupType}/ | 
[**GroupV2GetInvitedIndividuals**](GroupV2Api.md#GroupV2GetInvitedIndividuals) | **Get** /GroupV2/{groupId}/Members/InvitedIndividuals/ | 
[**GroupV2GetMembersOfGroup**](GroupV2Api.md#GroupV2GetMembersOfGroup) | **Get** /GroupV2/{groupId}/Members/ | 
[**GroupV2GetPendingMemberships**](GroupV2Api.md#GroupV2GetPendingMemberships) | **Get** /GroupV2/{groupId}/Members/Pending/ | 
[**GroupV2GetPotentialGroupsForMember**](GroupV2Api.md#GroupV2GetPotentialGroupsForMember) | **Get** /GroupV2/User/Potential/{membershipType}/{membershipId}/{filter}/{groupType}/ | 
[**GroupV2GetRecommendedGroups**](GroupV2Api.md#GroupV2GetRecommendedGroups) | **Post** /GroupV2/Recommended/{groupType}/{createDateRange}/ | 
[**GroupV2GetUserClanInviteSetting**](GroupV2Api.md#GroupV2GetUserClanInviteSetting) | **Get** /GroupV2/GetUserClanInviteSetting/{mType}/ | 
[**GroupV2GroupSearch**](GroupV2Api.md#GroupV2GroupSearch) | **Post** /GroupV2/Search/ | 
[**GroupV2IndividualGroupInvite**](GroupV2Api.md#GroupV2IndividualGroupInvite) | **Post** /GroupV2/{groupId}/Members/IndividualInvite/{membershipType}/{membershipId}/ | 
[**GroupV2IndividualGroupInviteCancel**](GroupV2Api.md#GroupV2IndividualGroupInviteCancel) | **Post** /GroupV2/{groupId}/Members/IndividualInviteCancel/{membershipType}/{membershipId}/ | 
[**GroupV2KickMember**](GroupV2Api.md#GroupV2KickMember) | **Post** /GroupV2/{groupId}/Members/{membershipType}/{membershipId}/Kick/ | 
[**GroupV2RequestGroupMembership**](GroupV2Api.md#GroupV2RequestGroupMembership) | **Post** /GroupV2/{groupId}/Members/Apply/{membershipType}/ | 
[**GroupV2RescindGroupMembership**](GroupV2Api.md#GroupV2RescindGroupMembership) | **Post** /GroupV2/{groupId}/Members/Rescind/{membershipType}/ | 
[**GroupV2SetUserClanInviteSetting**](GroupV2Api.md#GroupV2SetUserClanInviteSetting) | **Post** /GroupV2/SetUserClanInviteSetting/{mType}/{allowInvites}/ | 
[**GroupV2UnbanMember**](GroupV2Api.md#GroupV2UnbanMember) | **Post** /GroupV2/{groupId}/Members/{membershipType}/{membershipId}/Unban/ | 


# **GroupV2AbdicateFoundership**
> InlineResponse20018 GroupV2AbdicateFoundership(ctx, founderIdNew, groupId, membershipType)


An administrative method to allow the founder of a group or clan to give up their position to another admin permanently.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **founderIdNew** | **int64**| The new founder for this group. Must already be a group admin. | 
  **groupId** | **int64**| The target group id. | 
  **membershipType** | [**BungieMembershipType**](.md)| Membership type of the provided founderIdNew. | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2AddOptionalConversation**
> InlineResponse20011 GroupV2AddOptionalConversation(ctx, groupId, groupsV2GroupOptionalConversationAddRequest)


Add a new optional conversation/chat channel. Requires admin permissions to the group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Group ID of the group to edit. | 
  **groupsV2GroupOptionalConversationAddRequest** | [**GroupsV2GroupOptionalConversationAddRequest**](GroupsV2GroupOptionalConversationAddRequest.md)|  | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2ApproveAllPending**
> InlineResponse20030 GroupV2ApproveAllPending(ctx, groupId, groupsV2GroupApplicationRequest)


Approve all of the pending users for the given group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| ID of the group. | 
  **groupsV2GroupApplicationRequest** | [**GroupsV2GroupApplicationRequest**](GroupsV2GroupApplicationRequest.md)|  | 

### Return type

[**InlineResponse20030**](inline_response_200_30.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2ApprovePending**
> InlineResponse20018 GroupV2ApprovePending(ctx, groupId, membershipId, membershipType, groupsV2GroupApplicationRequest)


Approve the given membershipId to join the group/clan as long as they have applied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| ID of the group. | 
  **membershipId** | **int64**| The membership id being approved. | 
  **membershipType** | [**BungieMembershipType**](.md)| Membership type of the supplied membership ID. | 
  **groupsV2GroupApplicationRequest** | [**GroupsV2GroupApplicationRequest**](GroupsV2GroupApplicationRequest.md)|  | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2ApprovePendingForList**
> InlineResponse20030 GroupV2ApprovePendingForList(ctx, groupId, groupsV2GroupApplicationListRequest)


Approve all of the pending users for the given group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| ID of the group. | 
  **groupsV2GroupApplicationListRequest** | [**GroupsV2GroupApplicationListRequest**](GroupsV2GroupApplicationListRequest.md)|  | 

### Return type

[**InlineResponse20030**](inline_response_200_30.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2BanMember**
> InlineResponse20019 GroupV2BanMember(ctx, groupId, membershipId, membershipType, groupsV2GroupBanRequest)


Bans the requested member from the requested group for the specified period of time.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Group ID that has the member to ban. | 
  **membershipId** | **int64**| Membership ID of the member to ban from the group. | 
  **membershipType** | [**BungieMembershipType**](.md)| Membership type of the provided membership ID. | 
  **groupsV2GroupBanRequest** | [**GroupsV2GroupBanRequest**](GroupsV2GroupBanRequest.md)|  | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2CreateGroup**
> InlineResponse20024 GroupV2CreateGroup(ctx, groupsV2GroupAction)


Create a new group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupsV2GroupAction** | [**GroupsV2GroupAction**](GroupsV2GroupAction.md)|  | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2DenyAllPending**
> InlineResponse20030 GroupV2DenyAllPending(ctx, groupId, groupsV2GroupApplicationRequest)


Deny all of the pending users for the given group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| ID of the group. | 
  **groupsV2GroupApplicationRequest** | [**GroupsV2GroupApplicationRequest**](GroupsV2GroupApplicationRequest.md)|  | 

### Return type

[**InlineResponse20030**](inline_response_200_30.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2DenyPendingForList**
> InlineResponse20030 GroupV2DenyPendingForList(ctx, groupId, groupsV2GroupApplicationListRequest)


Deny all of the pending users for the given group that match the passed-in .

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| ID of the group. | 
  **groupsV2GroupApplicationListRequest** | [**GroupsV2GroupApplicationListRequest**](GroupsV2GroupApplicationListRequest.md)|  | 

### Return type

[**InlineResponse20030**](inline_response_200_30.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2EditClanBanner**
> InlineResponse20019 GroupV2EditClanBanner(ctx, groupId, groupsV2ClanBanner)


Edit an existing group's clan banner. You must have suitable permissions in the group to perform this operation. All fields are required.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Group ID of the group to edit. | 
  **groupsV2ClanBanner** | [**GroupsV2ClanBanner**](GroupsV2ClanBanner.md)|  | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2EditFounderOptions**
> InlineResponse20019 GroupV2EditFounderOptions(ctx, groupId, groupsV2GroupOptionsEditAction)


Edit group options only available to a founder. You must have suitable permissions in the group to perform this operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Group ID of the group to edit. | 
  **groupsV2GroupOptionsEditAction** | [**GroupsV2GroupOptionsEditAction**](GroupsV2GroupOptionsEditAction.md)|  | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2EditGroup**
> InlineResponse20019 GroupV2EditGroup(ctx, groupId, groupsV2GroupEditAction)


Edit an existing group. You must have suitable permissions in the group to perform this operation. This latest revision will only edit the fields you pass in - pass null for properties you want to leave unaltered.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Group ID of the group to edit. | 
  **groupsV2GroupEditAction** | [**GroupsV2GroupEditAction**](GroupsV2GroupEditAction.md)|  | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2EditGroupMembership**
> InlineResponse20019 GroupV2EditGroupMembership(ctx, groupId, membershipId, membershipType, memberType)


Edit the membership type of a given member. You must have suitable permissions in the group to perform this operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| ID of the group to which the member belongs. | 
  **membershipId** | **int64**| Membership ID to modify. | 
  **membershipType** | [**BungieMembershipType**](.md)| Membership type of the provide membership ID. | 
  **memberType** | [**GroupsV2RuntimeGroupMemberType**](.md)| New membertype for the specified member. | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2EditOptionalConversation**
> InlineResponse20011 GroupV2EditOptionalConversation(ctx, conversationId, groupId, groupsV2GroupOptionalConversationEditRequest)


Edit the settings of an optional conversation/chat channel. Requires admin permissions to the group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **conversationId** | **int64**| Conversation Id of the channel being edited. | 
  **groupId** | **int64**| Group ID of the group to edit. | 
  **groupsV2GroupOptionalConversationEditRequest** | [**GroupsV2GroupOptionalConversationEditRequest**](GroupsV2GroupOptionalConversationEditRequest.md)|  | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetAdminsAndFounderOfGroup**
> InlineResponse20025 GroupV2GetAdminsAndFounderOfGroup(ctx, currentpage, groupId)


Get the list of members in a given group who are of admin level or higher.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currentpage** | **int32**| Page number (starting with 1). Each page has a fixed size of 50 items per page. | 
  **groupId** | **int64**| The ID of the group. | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetAvailableAvatars**
> InlineResponse20016 GroupV2GetAvailableAvatars(ctx, )


Returns a list of all available group avatars for the signed-in user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetAvailableThemes**
> InlineResponse20017 GroupV2GetAvailableThemes(ctx, )


Returns a list of all available group themes.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetBannedMembersOfGroup**
> InlineResponse20027 GroupV2GetBannedMembersOfGroup(ctx, currentpage, groupId)


Get the list of banned members in a given group. Only accessible to group Admins and above. Not applicable to all groups. Check group features.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currentpage** | **int32**| Page number (starting with 1). Each page has a fixed size of 50 entries. | 
  **groupId** | **int64**| Group ID whose banned members you are fetching | 

### Return type

[**InlineResponse20027**](inline_response_200_27.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetGroup**
> InlineResponse20022 GroupV2GetGroup(ctx, groupId)


Get information about a specific group of the given ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Requested group&#39;s id. | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetGroupByName**
> InlineResponse20022 GroupV2GetGroupByName(ctx, groupName, groupType)


Get information about a specific group with the given name and type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupName** | **string**| Exact name of the group to find. | 
  **groupType** | [**GroupsV2GroupType**](.md)| Type of group to find. | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetGroupByNameV2**
> InlineResponse20022 GroupV2GetGroupByNameV2(ctx, groupsV2GroupNameSearchRequest)


Get information about a specific group with the given name and type. The POST version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupsV2GroupNameSearchRequest** | [**GroupsV2GroupNameSearchRequest**](GroupsV2GroupNameSearchRequest.md)|  | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetGroupOptionalConversations**
> InlineResponse20023 GroupV2GetGroupOptionalConversations(ctx, groupId)


Gets a list of available optional conversation channels and their settings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Requested group&#39;s id. | 

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetGroupsForMember**
> InlineResponse20031 GroupV2GetGroupsForMember(ctx, filter, groupType, membershipId, membershipType)


Get information about the groups that a given member has joined.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **filter** | [**GroupsV2GroupsForMemberFilter**](.md)| Filter apply to list of joined groups. | 
  **groupType** | [**GroupsV2GroupType**](.md)| Type of group the supplied member founded. | 
  **membershipId** | **int64**| Membership ID to for which to find founded groups. | 
  **membershipType** | [**BungieMembershipType**](.md)| Membership type of the supplied membership ID. | 

### Return type

[**InlineResponse20031**](inline_response_200_31.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetInvitedIndividuals**
> InlineResponse20029 GroupV2GetInvitedIndividuals(ctx, currentpage, groupId)


Get the list of users who have been invited into the group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currentpage** | **int32**| Page number (starting with 1). Each page has a fixed size of 50 items per page. | 
  **groupId** | **int64**| ID of the group. | 

### Return type

[**InlineResponse20029**](inline_response_200_29.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetMembersOfGroup**
> InlineResponse20025 GroupV2GetMembersOfGroup(ctx, currentpage, groupId, optional)


Get the list of members in a given group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currentpage** | **int32**| Page number (starting with 1). Each page has a fixed size of 50 items per page. | 
  **groupId** | **int64**| The ID of the group. | 
 **optional** | ***GroupV2GetMembersOfGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupV2GetMembersOfGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **memberType** | [**optional.Interface of GroupsV2RuntimeGroupMemberType**](.md)| Filter out other member types. Use None for all members. | 
 **nameSearch** | **optional.String**| The name fragment upon which a search should be executed for members with matching display or unique names. | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetPendingMemberships**
> InlineResponse20029 GroupV2GetPendingMemberships(ctx, currentpage, groupId)


Get the list of users who are awaiting a decision on their application to join a given group. Modified to include application info.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **currentpage** | **int32**| Page number (starting with 1). Each page has a fixed size of 50 items per page. | 
  **groupId** | **int64**| ID of the group. | 

### Return type

[**InlineResponse20029**](inline_response_200_29.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetPotentialGroupsForMember**
> InlineResponse20032 GroupV2GetPotentialGroupsForMember(ctx, filter, groupType, membershipId, membershipType)


Get information about the groups that a given member has applied to or been invited to.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **filter** | [**GroupsV2GroupPotentialMemberStatus**](.md)| Filter apply to list of potential joined groups. | 
  **groupType** | [**GroupsV2GroupType**](.md)| Type of group the supplied member applied. | 
  **membershipId** | **int64**| Membership ID to for which to find applied groups. | 
  **membershipType** | [**BungieMembershipType**](.md)| Membership type of the supplied membership ID. | 

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetRecommendedGroups**
> InlineResponse20020 GroupV2GetRecommendedGroups(ctx, createDateRange, groupType)


Gets groups recommended for you based on the groups to whom those you follow belong.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createDateRange** | [**GroupsV2GroupDateRange**](.md)| Requested range in which to pull recommended groups | 
  **groupType** | [**GroupsV2GroupType**](.md)| Type of groups requested | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GetUserClanInviteSetting**
> InlineResponse20018 GroupV2GetUserClanInviteSetting(ctx, mType)


Gets the state of the user's clan invite preferences for a particular membership type - true if they wish to be invited to clans, false otherwise.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mType** | [**BungieMembershipType**](.md)| The Destiny membership type of the account we wish to access settings. | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2GroupSearch**
> InlineResponse20021 GroupV2GroupSearch(ctx, groupsV2GroupQuery)


Search for Groups.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupsV2GroupQuery** | [**GroupsV2GroupQuery**](GroupsV2GroupQuery.md)|  | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2IndividualGroupInvite**
> InlineResponse20028 GroupV2IndividualGroupInvite(ctx, groupId, membershipId, membershipType, groupsV2GroupApplicationRequest)


Invite a user to join this group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| ID of the group you would like to join. | 
  **membershipId** | **int64**| Membership id of the account being invited. | 
  **membershipType** | [**BungieMembershipType**](.md)| MembershipType of the account being invited. | 
  **groupsV2GroupApplicationRequest** | [**GroupsV2GroupApplicationRequest**](GroupsV2GroupApplicationRequest.md)|  | 

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2IndividualGroupInviteCancel**
> InlineResponse20028 GroupV2IndividualGroupInviteCancel(ctx, groupId, membershipId, membershipType)


Cancels a pending invitation to join a group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| ID of the group you would like to join. | 
  **membershipId** | **int64**| Membership id of the account being cancelled. | 
  **membershipType** | [**BungieMembershipType**](.md)| MembershipType of the account being cancelled. | 

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2KickMember**
> InlineResponse20026 GroupV2KickMember(ctx, groupId, membershipId, membershipType)


Kick a member from the given group, forcing them to reapply if they wish to re-join the group. You must have suitable permissions in the group to perform this operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Group ID to kick the user from. | 
  **membershipId** | **int64**| Membership ID to kick. | 
  **membershipType** | [**BungieMembershipType**](.md)| Membership type of the provided membership ID. | 

### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2RequestGroupMembership**
> InlineResponse20028 GroupV2RequestGroupMembership(ctx, groupId, membershipType, groupsV2GroupApplicationRequest)


Request permission to join the given group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| ID of the group you would like to join. | 
  **membershipType** | [**BungieMembershipType**](.md)| MembershipType of the account to use when joining. | 
  **groupsV2GroupApplicationRequest** | [**GroupsV2GroupApplicationRequest**](GroupsV2GroupApplicationRequest.md)|  | 

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2RescindGroupMembership**
> InlineResponse20026 GroupV2RescindGroupMembership(ctx, groupId, membershipType)


Rescind your application to join the given group or leave the group if you are already a member..

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| ID of the group. | 
  **membershipType** | [**BungieMembershipType**](.md)| MembershipType of the account to leave. | 

### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2SetUserClanInviteSetting**
> InlineResponse20019 GroupV2SetUserClanInviteSetting(ctx, allowInvites, mType)


Sets the state of the user's clan invite preferences - true if they wish to be invited to clans, false otherwise.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **allowInvites** | **bool**| True to allow invites of this user to clans, false otherwise. | 
  **mType** | [**BungieMembershipType**](.md)| The Destiny membership type of linked account we are manipulating. | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupV2UnbanMember**
> InlineResponse20019 GroupV2UnbanMember(ctx, groupId, membershipId, membershipType)


Unbans the requested member, allowing them to re-apply for membership.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**|  | 
  **membershipId** | **int64**| Membership ID of the member to unban from the group | 
  **membershipType** | [**BungieMembershipType**](.md)| Membership type of the provided membership ID. | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

