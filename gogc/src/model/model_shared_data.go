/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// SharedData - Shared Data
type SharedData struct {

	SharedDataId string `json:"sharedDataId"`

	SharedProfileData *NfProfile `json:"sharedProfileData,omitempty"`

	SharedServiceData NfService `json:"sharedServiceData,omitempty"`

	AuthorizedWriteScope SharedScope `json:"authorizedWriteScope,omitempty"`
}
