/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// SnssaiUpfInfoItem - Set of parameters supported by UPF for a given S-NSSAI
type SnssaiUpfInfoItem struct {

	SNssai ExtSnssai `json:"sNssai"`

	DnnUpfInfoList []DnnUpfInfoItem `json:"dnnUpfInfoList"`

	RedundantTransport bool `json:"redundantTransport,omitempty"`
}
