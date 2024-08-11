/*
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// StoredSearchResult - Contains a complete search result (i.e. a number of discovered NF Instances), stored by NRF as a consequence of a prior search result 
type StoredSearchResult struct {

	NfInstances []NfProfile `json:"nfInstances"`

	CompleteNfInstances []NfProfile `json:"completeNfInstances,omitempty"`
}
