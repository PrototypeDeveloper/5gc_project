/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// TsctsfInfo - Information of a TSCTSF NF Instance
type TsctsfInfo struct {

	// A map (list of key-value pairs) where a valid JSON string serves as key
	SNssaiInfoList map[string]SnssaiTsctsfInfoItem `json:"sNssaiInfoList,omitempty"`

	ExternalGroupIdentifiersRanges []IdentityRange `json:"externalGroupIdentifiersRanges,omitempty"`

	SupiRanges []SupiRange `json:"supiRanges,omitempty"`

	GpsiRanges []IdentityRange `json:"gpsiRanges,omitempty"`

	InternalGroupIdentifiersRanges []InternalGroupIdRange `json:"internalGroupIdentifiersRanges,omitempty"`
}
