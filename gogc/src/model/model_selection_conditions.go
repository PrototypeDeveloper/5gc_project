/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"gopkg.in/validator.v2"
)

// SelectionConditions - It contains the set of conditions that shall be evaluated to determine whether a consumer shall select a given producer. The producer shall only be selected if the evaluation of the conditions is <true>. The set of conditions can be represented by a single  ConditionItem or by a ConditionGroup, where the latter contains a (recursive) list of conditions joined by the \"and\" or \"or\" logical relationships. 
type SelectionConditions struct {

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

	And []SelectionConditions `json:"and,omitempty"`

	Or []SelectionConditions `json:"or,omitempty"`
}