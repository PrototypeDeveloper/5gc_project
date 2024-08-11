/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// N2InterfaceAmfInfo - AMF N2 interface information
type N2InterfaceAmfInfo struct {

	Ipv4EndpointAddress []string `json:"ipv4EndpointAddress,omitempty"`

	Ipv6EndpointAddress []Ipv6Addr `json:"ipv6EndpointAddress,omitempty"`

	// Fully Qualified Domain Name
	AmfName string `json:"amfName,omitempty"`
}
