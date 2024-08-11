/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ConditionGroup - List (array) of conditions (joined by the \"and\" or \"or\" logical relationship), under which an NF Instance with an NFStatus or NFServiceStatus value set to, \"CANARY_RELEASE\", or with a \"canaryRelease\" attribute set to true, shall be selected by an NF Service Consumer. 
type ConditionGroup struct {

	And []SelectionConditions `json:"and,omitempty"`

	Or []SelectionConditions `json:"or,omitempty"`
}