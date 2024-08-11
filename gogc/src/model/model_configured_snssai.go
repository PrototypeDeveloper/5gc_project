/*
 * NSSF NS Selection
 *
 * NSSF Network Slice Selection Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ConfiguredSnssai - Contains the configured S-NSSAI authorized by the NSSF in the serving PLMN and optional mapped home S-NSSAI 
type ConfiguredSnssai struct {

	ConfiguredSnssai Snssai `json:"configuredSnssai"`

	MappedHomeSnssai Snssai `json:"mappedHomeSnssai,omitempty"`
}
