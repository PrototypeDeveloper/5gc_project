/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// MbsServiceArea - MBS Service Area
type MbsServiceArea struct {

	// List of NR cell Ids
	NcgiList []NcgiTai `json:"ncgiList,omitempty"`

	// List of tracking area Ids
	TaiList []Tai `json:"taiList,omitempty"`
}
