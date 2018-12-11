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

type FireteamFireteamMember struct {
	DestinyUserInfo UserUserInfoCard `json:"destinyUserInfo,omitempty"`
	BungieNetUserInfo UserUserInfoCard `json:"bungieNetUserInfo,omitempty"`
	CharacterId int64 `json:"characterId,omitempty"`
	DateJoined time.Time `json:"dateJoined,omitempty"`
	HasMicrophone bool `json:"hasMicrophone,omitempty"`
	LastPlatformInviteAttemptDate time.Time `json:"lastPlatformInviteAttemptDate,omitempty"`
	LastPlatformInviteAttemptResult FireteamFireteamPlatformInviteResult `json:"lastPlatformInviteAttemptResult,omitempty"`
}