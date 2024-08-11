/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// MlAnalyticsInfo - ML Analytics Filter information supported by the Nnwdaf_MLModelProvision service
type MlAnalyticsInfo struct {

	MlAnalyticsIds []NwdafEvent `json:"mlAnalyticsIds,omitempty"`

	SnssaiList []Snssai `json:"snssaiList,omitempty"`

	TrackingAreaList []Tai `json:"trackingAreaList,omitempty"`

	MlModelInterInfo MlModelInterInfo `json:"mlModelInterInfo,omitempty"`

	FlCapabilityType FlCapabilityType `json:"flCapabilityType,omitempty"`

	// indicating a time in seconds.
	FlTimeInterval int32 `json:"flTimeInterval,omitempty"`

	NfTypeList []NfType `json:"nfTypeList,omitempty"`

	NfSetIdList []string `json:"nfSetIdList,omitempty"`
}
