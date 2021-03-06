/*
 * Bungie.Net API
 *
 * These endpoints constitute the functionality exposed by Bungie.net, both for more traditional website functionality and for connectivity to Bungie video games and their related functionality.
 *
 * API version: 2.3.2
 * Contact: support@bungie.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package bungie
import (
	"time"
)

type GroupsV2GroupPotentialMember struct {
	PotentialStatus GroupsV2GroupPotentialMemberStatus `json:"potentialStatus,omitempty"`
	GroupId int64 `json:"groupId,omitempty"`
	DestinyUserInfo UserUserInfoCard `json:"destinyUserInfo,omitempty"`
	BungieNetUserInfo UserUserInfoCard `json:"bungieNetUserInfo,omitempty"`
	JoinDate time.Time `json:"joinDate,omitempty"`
}
