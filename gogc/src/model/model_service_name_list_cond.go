/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ServiceNameListCond - Subscription to a set of NFs based on their support for a Service Name in the Servic Name list 
type ServiceNameListCond struct {

	ConditionType string `json:"conditionType"`

	ServiceNameList []ServiceName `json:"serviceNameList"`
}
