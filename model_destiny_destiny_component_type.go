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
// DestinyDestinyComponentType : Represents the possible components that can be returned from Destiny \"Get\" calls such as GetProfile, GetCharacter, GetVendor etc...  When making one of these requests, you will pass one or more of these components as a comma separated list in the \"?components=\" querystring parameter. For instance, if you want baseline Profile data, Character Data, and character progressions, you would pass \"?components=Profiles,Characters,CharacterProgressions\" You may use either the numerical or string values.
type DestinyDestinyComponentType int32

// List of Destiny.DestinyComponentType
const (
	NONE DestinyDestinyComponentType = "0"
	PROFILES DestinyDestinyComponentType = "100"
	VENDORRECEIPTS DestinyDestinyComponentType = "101"
	PROFILEINVENTORIES DestinyDestinyComponentType = "102"
	PROFILECURRENCIES DestinyDestinyComponentType = "103"
	PROFILEPROGRESSION DestinyDestinyComponentType = "104"
	CHARACTERS DestinyDestinyComponentType = "200"
	CHARACTERINVENTORIES DestinyDestinyComponentType = "201"
	CHARACTERPROGRESSIONS DestinyDestinyComponentType = "202"
	CHARACTERRENDERDATA DestinyDestinyComponentType = "203"
	CHARACTERACTIVITIES DestinyDestinyComponentType = "204"
	CHARACTEREQUIPMENT DestinyDestinyComponentType = "205"
	ITEMINSTANCES DestinyDestinyComponentType = "300"
	ITEMOBJECTIVES DestinyDestinyComponentType = "301"
	ITEMPERKS DestinyDestinyComponentType = "302"
	ITEMRENDERDATA DestinyDestinyComponentType = "303"
	ITEMSTATS DestinyDestinyComponentType = "304"
	ITEMSOCKETS DestinyDestinyComponentType = "305"
	ITEMTALENTGRIDS DestinyDestinyComponentType = "306"
	ITEMCOMMONDATA DestinyDestinyComponentType = "307"
	ITEMPLUGSTATES DestinyDestinyComponentType = "308"
	VENDORS DestinyDestinyComponentType = "400"
	VENDORCATEGORIES DestinyDestinyComponentType = "401"
	VENDORSALES DestinyDestinyComponentType = "402"
	KIOSKS DestinyDestinyComponentType = "500"
	CURRENCYLOOKUPS DestinyDestinyComponentType = "600"
	PRESENTATIONNODES DestinyDestinyComponentType = "700"
	COLLECTIBLES DestinyDestinyComponentType = "800"
	RECORDS DestinyDestinyComponentType = "900"
)