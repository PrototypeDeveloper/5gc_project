/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// EasdfInfo - Information of an EASDF NF Instance
type EasdfInfo struct {

	SNssaiEasdfInfoList []SnssaiEasdfInfoItem `json:"sNssaiEasdfInfoList,omitempty"`

	EasdfN6IpAddressList []IpAddr `json:"easdfN6IpAddressList,omitempty"`

	UpfN6IpAddressList []IpAddr `json:"upfN6IpAddressList,omitempty"`
}
