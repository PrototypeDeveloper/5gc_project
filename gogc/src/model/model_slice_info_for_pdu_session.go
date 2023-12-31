/*
 * NSSF NS Selection
 *
 * NSSF Network Slice Selection Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// SliceInfoForPduSession - Contains the slice information requested during PDU Session establishment procedure
type SliceInfoForPduSession struct {

	SNssai Snssai `json:"sNssai"`

	RoamingIndication RoamingIndication `json:"roamingIndication"`

	HomeSnssai Snssai `json:"homeSnssai,omitempty"`
}
