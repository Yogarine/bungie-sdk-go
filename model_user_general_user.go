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

type UserGeneralUser struct {
	MembershipId int64 `json:"membershipId,omitempty"`
	UniqueName string `json:"uniqueName,omitempty"`
	NormalizedName string `json:"normalizedName,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	ProfilePicture int32 `json:"profilePicture,omitempty"`
	ProfileTheme int32 `json:"profileTheme,omitempty"`
	UserTitle int32 `json:"userTitle,omitempty"`
	SuccessMessageFlags int64 `json:"successMessageFlags,omitempty"`
	IsDeleted bool `json:"isDeleted,omitempty"`
	About string `json:"about,omitempty"`
	FirstAccess time.Time `json:"firstAccess,omitempty"`
	LastUpdate time.Time `json:"lastUpdate,omitempty"`
	LegacyPortalUID int64 `json:"legacyPortalUID,omitempty"`
	Context UserUserToUserContext `json:"context,omitempty"`
	PsnDisplayName string `json:"psnDisplayName,omitempty"`
	XboxDisplayName string `json:"xboxDisplayName,omitempty"`
	FbDisplayName string `json:"fbDisplayName,omitempty"`
	ShowActivity bool `json:"showActivity,omitempty"`
	Locale string `json:"locale,omitempty"`
	LocaleInheritDefault bool `json:"localeInheritDefault,omitempty"`
	LastBanReportId int64 `json:"lastBanReportId,omitempty"`
	ShowGroupMessaging bool `json:"showGroupMessaging,omitempty"`
	ProfilePicturePath string `json:"profilePicturePath,omitempty"`
	ProfilePictureWidePath string `json:"profilePictureWidePath,omitempty"`
	ProfileThemeName string `json:"profileThemeName,omitempty"`
	UserTitleDisplay string `json:"userTitleDisplay,omitempty"`
	StatusText string `json:"statusText,omitempty"`
	StatusDate time.Time `json:"statusDate,omitempty"`
	ProfileBanExpire time.Time `json:"profileBanExpire,omitempty"`
	BlizzardDisplayName string `json:"blizzardDisplayName,omitempty"`
}
