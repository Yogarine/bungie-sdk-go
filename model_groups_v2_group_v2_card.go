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

// A small infocard of group information, usually used for when a list of groups are returned
type GroupsV2GroupV2Card struct {
	GroupId int64 `json:"groupId,omitempty"`
	Name string `json:"name,omitempty"`
	GroupType GroupsV2GroupType `json:"groupType,omitempty"`
	CreationDate time.Time `json:"creationDate,omitempty"`
	About string `json:"about,omitempty"`
	Motto string `json:"motto,omitempty"`
	MemberCount int32 `json:"memberCount,omitempty"`
	Locale string `json:"locale,omitempty"`
	MembershipOption GroupsV2MembershipOption `json:"membershipOption,omitempty"`
	Capabilities GroupsV2Capabilities `json:"capabilities,omitempty"`
	ClanInfo GroupsV2GroupV2ClanInfo `json:"clanInfo,omitempty"`
	AvatarPath string `json:"avatarPath,omitempty"`
	Theme string `json:"theme,omitempty"`
}
