/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// NcgiTai - List of NR cell ids, with their pertaining TAIs
type NcgiTai struct {

	Tai Tai `json:"tai"`

	// List of List of NR cell ids
	CellList []Ncgi `json:"cellList"`
}
