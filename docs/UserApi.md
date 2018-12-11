# Bungie\UserApi

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserGetAvailableThemes**](UserApi.md#UserGetAvailableThemes) | **Get** /User/GetAvailableThemes/ | 
[**UserGetBungieNetUserById**](UserApi.md#UserGetBungieNetUserById) | **Get** /User/GetBungieNetUserById/{id}/ | 
[**UserGetMembershipDataById**](UserApi.md#UserGetMembershipDataById) | **Get** /User/GetMembershipsById/{membershipId}/{membershipType}/ | 
[**UserGetMembershipDataForCurrentUser**](UserApi.md#UserGetMembershipDataForCurrentUser) | **Get** /User/GetMembershipsForCurrentUser/ | 
[**UserGetPartnerships**](UserApi.md#UserGetPartnerships) | **Get** /User/{membershipId}/Partnerships/ | 
[**UserSearchUsers**](UserApi.md#UserSearchUsers) | **Get** /User/SearchUsers/ | 


# **UserGetAvailableThemes**
> InlineResponse2004 UserGetAvailableThemes(ctx, )


Returns a list of all available user themes.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGetBungieNetUserById**
> InlineResponse2002 UserGetBungieNetUserById(ctx, id)


Loads a bungienet user by membership id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The requested Bungie.net membership id. | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGetMembershipDataById**
> InlineResponse2005 UserGetMembershipDataById(ctx, membershipId, membershipType)


Returns a list of accounts associated with the supplied membership ID and membership type. This will include all linked accounts (even when hidden) if supplied credentials permit it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **membershipId** | **int64**| The membership ID of the target user. | 
  **membershipType** | [**BungieMembershipType**](.md)| Type of the supplied membership ID. | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGetMembershipDataForCurrentUser**
> InlineResponse2005 UserGetMembershipDataForCurrentUser(ctx, )


Returns a list of accounts associated with signed in user. This is useful for OAuth implementations that do not give you access to the token response.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGetPartnerships**
> InlineResponse2006 UserGetPartnerships(ctx, membershipId)


Returns a user's linked Partnerships.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **membershipId** | **int64**| The ID of the member for whom partnerships should be returned. | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserSearchUsers**
> InlineResponse2003 UserSearchUsers(ctx, optional)


Returns a list of possible users based on the search string

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UserSearchUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserSearchUsersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**| The search string. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

