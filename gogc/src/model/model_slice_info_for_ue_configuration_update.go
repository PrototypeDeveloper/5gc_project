/*
 * NSSF NS Selection
 *
 * NSSF Network Slice Selection Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// SliceInfoForUeConfigurationUpdate - Contains the slice information requested during UE configuration update procedure
type SliceInfoForUeConfigurationUpdate struct {

	SubscribedNssai []SubscribedSnssai `json:"subscribedNssai,omitempty"`

	AllowedNssaiCurrentAccess AllowedNssai `json:"allowedNssaiCurrentAccess,omitempty"`

	AllowedNssaiOtherAccess AllowedNssai `json:"allowedNssaiOtherAccess,omitempty"`

	DefaultConfiguredSnssaiInd bool `json:"defaultConfiguredSnssaiInd,omitempty"`

	RequestedNssai []Snssai `json:"requestedNssai,omitempty"`

	MappingOfNssai []MappingOfSnssai `json:"mappingOfNssai,omitempty"`

	UeSupNssrgInd bool `json:"ueSupNssrgInd,omitempty"`

	SuppressNssrgInd bool `json:"suppressNssrgInd,omitempty"`

	RejectedNssaiRa []Snssai `json:"rejectedNssaiRa,omitempty"`

	NsagSupported bool `json:"nsagSupported,omitempty"`
}
