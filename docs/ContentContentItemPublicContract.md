# ContentContentItemPublicContract

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentId** | **int64** |  | [optional] 
**CType** | **string** |  | [optional] 
**CmsPath** | **string** |  | [optional] 
**CreationDate** | [**time.Time**](time.Time.md) |  | [optional] 
**ModifyDate** | [**time.Time**](time.Time.md) |  | [optional] 
**AllowComments** | **bool** |  | [optional] 
**HasAgeGate** | **bool** |  | [optional] 
**MinimumAge** | **int32** |  | [optional] 
**RatingImagePath** | **string** |  | [optional] 
**Author** | [**UserGeneralUser**](User.GeneralUser.md) |  | [optional] 
**AutoEnglishPropertyFallback** | **bool** |  | [optional] 
**Properties** | [**map[string]map[string]interface{}**](map[string]interface{}.md) | Firehose content is really a collection of metadata and \&quot;properties\&quot;, which are the potentially-but-not-strictly localizable data that comprises the meat of whatever content is being shown.  As Cole Porter would have crooned, \&quot;Anything Goes\&quot; with Firehose properties. They are most often strings, but they can theoretically be anything. They are JSON encoded, and could be JSON structures, simple strings, numbers etc... The Content Type of the item (cType) will describe the properties, and thus how they ought to be deserialized. | [optional] 
**Representations** | [**[]ContentContentRepresentation**](Content.ContentRepresentation.md) |  | [optional] 
**Tags** | **[]string** |  | [optional] 
**CommentSummary** | [**ContentCommentSummary**](Content.CommentSummary.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


