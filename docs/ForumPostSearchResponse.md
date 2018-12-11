# ForumPostSearchResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelatedPosts** | [**[]ForumPostResponse**](Forum.PostResponse.md) |  | [optional] 
**Authors** | [**[]UserGeneralUser**](User.GeneralUser.md) |  | [optional] 
**Groups** | [**[]GroupsV2GroupResponse**](GroupsV2.GroupResponse.md) |  | [optional] 
**SearchedTags** | [**[]TagsModelsContractsTagResponse**](Tags.Models.Contracts.TagResponse.md) |  | [optional] 
**Polls** | [**[]ForumPollResponse**](Forum.PollResponse.md) |  | [optional] 
**RecruitmentDetails** | [**[]ForumForumRecruitmentDetail**](Forum.ForumRecruitmentDetail.md) |  | [optional] 
**AvailablePages** | **int32** |  | [optional] 
**Results** | [**[]ForumPostResponse**](Forum.PostResponse.md) |  | [optional] 
**TotalResults** | **int32** |  | [optional] 
**HasMore** | **bool** |  | [optional] 
**Query** | [**QueriesPagedQuery**](Queries.PagedQuery.md) |  | [optional] 
**ReplacementContinuationToken** | **string** |  | [optional] 
**UseTotalResults** | **bool** | If useTotalResults is true, then totalResults represents an accurate count.  If False, it does not, and may be estimated/only the size of the current page.  Either way, you should probably always only trust hasMore.  This is a long-held historical throwback to when we used to do paging with known total results. Those queries toasted our database, and we were left to hastily alter our endpoints and create backward- compatible shims, of which useTotalResults is one. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


