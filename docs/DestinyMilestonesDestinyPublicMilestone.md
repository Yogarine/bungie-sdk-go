# DestinyMilestonesDestinyPublicMilestone

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MilestoneHash** | **int32** | The hash identifier for the milestone. Use it to look up the DestinyMilestoneDefinition for static data about the Milestone. | [optional] 
**AvailableQuests** | [**[]DestinyMilestonesDestinyPublicMilestoneQuest**](Destiny.Milestones.DestinyPublicMilestoneQuest.md) | A milestone not need have even a single quest, but if there are active quests they will be returned here. | [optional] 
**Activities** | [**[]DestinyMilestonesDestinyPublicMilestoneChallengeActivity**](Destiny.Milestones.DestinyPublicMilestoneChallengeActivity.md) |  | [optional] 
**VendorHashes** | **[]int32** | Sometimes milestones - or activities active in milestones - will have relevant vendors. These are the vendors that are currently relevant.  Deprecated, already, for the sake of the new \&quot;vendors\&quot; property that has more data. What was I thinking. | [optional] 
**Vendors** | [**[]DestinyMilestonesDestinyPublicMilestoneVendor**](Destiny.Milestones.DestinyPublicMilestoneVendor.md) | This is why we can&#39;t have nice things. This is the ordered list of vendors to be shown that relate to this milestone, potentially along with other interesting data. | [optional] 
**StartDate** | [**time.Time**](time.Time.md) | If known, this is the date when the Milestone started/became active. | [optional] 
**EndDate** | [**time.Time**](time.Time.md) | If known, this is the date when the Milestone will expire/recycle/end. | [optional] 
**Order** | **int32** | Used for ordering milestones in a display to match how we order them in BNet. May pull from static data, or possibly in the future from dynamic information. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


