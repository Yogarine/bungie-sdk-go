# Bungie\AppApi

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppGetApplicationApiUsage**](AppApi.md#AppGetApplicationApiUsage) | **Get** /App/ApiUsage/{applicationId}/ | 
[**AppGetBungieApplications**](AppApi.md#AppGetBungieApplications) | **Get** /App/FirstParty/ | 


# **AppGetApplicationApiUsage**
> InlineResponse200 AppGetApplicationApiUsage(ctx, applicationId, optional)


Get API usage by application for time frame specified. You can go as far back as 30 days ago, and can ask for up to a 48 hour window of time in a single request. You must be authenticated with at least the ReadUserData permission to access this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationId** | **int32**| ID of the application to get usage statistics. | 
 **optional** | ***AppGetApplicationApiUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppGetApplicationApiUsageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **end** | **optional.Time**| End time for query. Goes to now if not specified. | 
 **start** | **optional.Time**| Start time for query. Goes to 24 hours ago if not specified. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[apiKey](../README.md#apiKey), [oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppGetBungieApplications**
> InlineResponse2001 AppGetBungieApplications(ctx, )


Get list of applications created by Bungie.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

