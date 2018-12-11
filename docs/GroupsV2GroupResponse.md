# GroupsV2GroupResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | [**GroupsV2GroupV2**](GroupsV2.GroupV2.md) |  | [optional] 
**Founder** | [**GroupsV2GroupMember**](GroupsV2.GroupMember.md) |  | [optional] 
**AlliedIds** | **[]int64** |  | [optional] 
**ParentGroup** | [**GroupsV2GroupV2**](GroupsV2.GroupV2.md) |  | [optional] 
**AllianceStatus** | [**GroupsV2GroupAllianceStatus**](GroupsV2.GroupAllianceStatus.md) |  | [optional] 
**GroupJoinInviteCount** | **int32** |  | [optional] 
**CurrentUserMemberMap** | [**map[string]GroupsV2GroupMember**](GroupsV2.GroupMember.md) | This property will be populated if the authenticated user is a member of the group. Note that because of account linking, a user can sometimes be part of a clan more than once. As such, this returns the highest member type available. | [optional] 
**CurrentUserPotentialMemberMap** | [**map[string]GroupsV2GroupPotentialMember**](GroupsV2.GroupPotentialMember.md) | This property will be populated if the authenticated user is an applicant or has an outstanding invitation to join. Note that because of account linking, a user can sometimes be part of a clan more than once. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


