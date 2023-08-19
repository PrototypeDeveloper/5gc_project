/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// DcsfInfo - Information of a DCSF NF Instance
type DcsfInfo struct {

	ImsDomianNameList []string `json:"imsDomianNameList,omitempty"`

	ImsiRanges []ImsiRange `json:"imsiRanges,omitempty"`

	ImsPrivateIdentityRanges []IdentityRange `json:"imsPrivateIdentityRanges,omitempty"`

	ImsPublicIdentityRanges []IdentityRange `json:"imsPublicIdentityRanges,omitempty"`

	MsisdnRanges []IdentityRange `json:"msisdnRanges,omitempty"`
}