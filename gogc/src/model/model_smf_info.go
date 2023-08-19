/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// SmfInfo - Information of an SMF NF Instance
type SmfInfo struct {

	SNssaiSmfInfoList []SnssaiSmfInfoItem `json:"sNssaiSmfInfoList"`

	TaiList []Tai `json:"taiList,omitempty"`

	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`

	// Fully Qualified Domain Name
	PgwFqdn string `json:"pgwFqdn,omitempty"`

	PgwIpAddrList []IpAddr `json:"pgwIpAddrList,omitempty"`

	AccessType []AccessType `json:"accessType,omitempty"`

	Priority int32 `json:"priority,omitempty"`

	VsmfSupportInd bool `json:"vsmfSupportInd,omitempty"`

	PgwFqdnList []string `json:"pgwFqdnList,omitempty"`

	// Deprecated
	SmfOnboardingCapability bool `json:"smfOnboardingCapability,omitempty"`

	IsmfSupportInd bool `json:"ismfSupportInd,omitempty"`

	SmfUPRPCapability bool `json:"smfUPRPCapability,omitempty"`
}