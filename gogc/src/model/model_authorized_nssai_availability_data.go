/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// AuthorizedNssaiAvailabilityData - This contains the Nssai availability data information per TA authorized by the NSSF
type AuthorizedNssaiAvailabilityData struct {

	Tai Tai `json:"tai"`

	SupportedSnssaiList []ExtSnssai `json:"supportedSnssaiList"`

	RestrictedSnssaiList []RestrictedSnssai `json:"restrictedSnssaiList,omitempty"`

	TaiList []Tai `json:"taiList,omitempty"`

	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`

	NsagInfos []NsagInfo `json:"nsagInfos,omitempty"`
}
