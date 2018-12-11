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

type ContentModelsContentTypeDescription struct {
	CType string `json:"cType,omitempty"`
	Name string `json:"name,omitempty"`
	ContentDescription string `json:"contentDescription,omitempty"`
	PreviewImage string `json:"previewImage,omitempty"`
	Priority int32 `json:"priority,omitempty"`
	Reminder string `json:"reminder,omitempty"`
	Properties []ContentModelsContentTypeProperty `json:"properties,omitempty"`
	TagMetadata []ContentModelsTagMetadataDefinition `json:"tagMetadata,omitempty"`
	TagMetadataItems map[string]ContentModelsTagMetadataItem `json:"tagMetadataItems,omitempty"`
	UsageExamples []string `json:"usageExamples,omitempty"`
	ShowInContentEditor bool `json:"showInContentEditor,omitempty"`
	TypeOf string `json:"typeOf,omitempty"`
	BindIdentifierToProperty string `json:"bindIdentifierToProperty,omitempty"`
	BoundRegex string `json:"boundRegex,omitempty"`
	ForceIdentifierBinding bool `json:"forceIdentifierBinding,omitempty"`
	AllowComments bool `json:"allowComments,omitempty"`
	AutoEnglishPropertyFallback bool `json:"autoEnglishPropertyFallback,omitempty"`
	BulkUploadable bool `json:"bulkUploadable,omitempty"`
	Previews []ContentModelsContentPreview `json:"previews,omitempty"`
	SuppressCmsPath bool `json:"suppressCmsPath,omitempty"`
	PropertySections []ContentModelsContentTypePropertySection `json:"propertySections,omitempty"`
}
