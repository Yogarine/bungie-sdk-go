# Bungie\DefaultApi

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAvailableLocales**](DefaultApi.md#GetAvailableLocales) | **Get** /GetAvailableLocales/ | 
[**GetCommonSettings**](DefaultApi.md#GetCommonSettings) | **Get** /Settings/ | 
[**GetGlobalAlerts**](DefaultApi.md#GetGlobalAlerts) | **Get** /GlobalAlerts/ | 


# **GetAvailableLocales**
> InlineResponse20068 GetAvailableLocales(ctx, )


List of available localization cultures

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20068**](inline_response_200_68.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCommonSettings**
> InlineResponse20069 GetCommonSettings(ctx, )


Get the common settings used by the Bungie.Net environment.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20069**](inline_response_200_69.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGlobalAlerts**
> InlineResponse20070 GetGlobalAlerts(ctx, optional)


Gets any active global alert for display in the forum banners, help pages, etc. Usually used for DOC alerts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetGlobalAlertsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetGlobalAlertsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includestreaming** | **optional.Bool**| Determines whether Streaming Alerts are included in results | 

### Return type

[**InlineResponse20070**](inline_response_200_70.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

