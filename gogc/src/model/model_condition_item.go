/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ConditionItem - A ConditionItem consists of a number of attributes representing individual conditions (e.g. a SUPI range, or a TAI list). If several attributes/conditions are present, the evaluation of the ConditionItem is <true> if all attributes/conditions are evaluated as <true> (i.e., it follows the AND logical relationship). 
type ConditionItem struct {

	ConsumerNfTypes []NfType `json:"consumerNfTypes,omitempty"`

	ServiceFeature int32 `json:"serviceFeature,omitempty"`

	VsServiceFeature int32 `json:"vsServiceFeature,omitempty"`

	SupiRangeList []SupiRange `json:"supiRangeList,omitempty"`

	GpsiRangeList []IdentityRange `json:"gpsiRangeList,omitempty"`

	ImpuRangeList []IdentityRange `json:"impuRangeList,omitempty"`

	ImpiRangeList []IdentityRange `json:"impiRangeList,omitempty"`

	PeiList []string `json:"peiList,omitempty"`

	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`

	DnnList []string `json:"dnnList,omitempty"`
}
