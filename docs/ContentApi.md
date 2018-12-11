# Bungie\ContentApi

All URIs are relative to *https://www.bungie.net/Platform*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContentGetContentById**](ContentApi.md#ContentGetContentById) | **Get** /Content/GetContentById/{id}/{locale}/ | 
[**ContentGetContentByTagAndType**](ContentApi.md#ContentGetContentByTagAndType) | **Get** /Content/GetContentByTagAndType/{tag}/{type}/{locale}/ | 
[**ContentGetContentType**](ContentApi.md#ContentGetContentType) | **Get** /Content/GetContentType/{type}/ | 
[**ContentSearchContentByTagAndType**](ContentApi.md#ContentSearchContentByTagAndType) | **Get** /Content/SearchContentByTagAndType/{tag}/{type}/{locale}/ | 
[**ContentSearchContentWithText**](ContentApi.md#ContentSearchContentWithText) | **Get** /Content/Search/{locale}/ | 


# **ContentGetContentById**
> InlineResponse2008 ContentGetContentById(ctx, id, locale, optional)


Returns a content item referenced by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**|  | 
  **locale** | **string**|  | 
 **optional** | ***ContentGetContentByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentGetContentByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **head** | **optional.Bool**| false | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentGetContentByTagAndType**
> InlineResponse2008 ContentGetContentByTagAndType(ctx, locale, tag, type_, optional)


Returns the newest item that matches a given tag and Content Type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locale** | **string**|  | 
  **tag** | **string**|  | 
  **type_** | **string**|  | 
 **optional** | ***ContentGetContentByTagAndTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentGetContentByTagAndTypeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **head** | **optional.Bool**| Not used. | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentGetContentType**
> InlineResponse2007 ContentGetContentType(ctx, type_)


Gets an object describing a particular variant of content.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentSearchContentByTagAndType**
> InlineResponse2009 ContentSearchContentByTagAndType(ctx, locale, tag, type_, optional)


Searches for Content Items that match the given Tag and Content Type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locale** | **string**|  | 
  **tag** | **string**|  | 
  **type_** | **string**|  | 
 **optional** | ***ContentSearchContentByTagAndTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentSearchContentByTagAndTypeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **currentpage** | **optional.Int32**| Page number for the search results starting with page 1. | 
 **head** | **optional.Bool**| Not used. | 
 **itemsperpage** | **optional.Int32**| Not used. | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContentSearchContentWithText**
> InlineResponse2009 ContentSearchContentWithText(ctx, locale, optional)


Gets content based on querystring information passed in. Provides basic search and text search capabilities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locale** | **string**|  | 
 **optional** | ***ContentSearchContentWithTextOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContentSearchContentWithTextOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ctype** | **optional.String**| Content type tag: Help, News, etc. Supply multiple ctypes separated by space. | 
 **currentpage** | **optional.Int32**| Page number for the search results, starting with page 1. | 
 **head** | **optional.Bool**| Not used. | 
 **searchtext** | **optional.String**| Word or phrase for the search. | 
 **source** | **optional.String**| For analytics, hint at the part of the app that triggered the search. Optional. | 
 **tag** | **optional.String**| Tag used on the content to be searched. | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

