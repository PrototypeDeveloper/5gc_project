/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// NetworkNodeDiameterAddress - This data type is a part of smsfDiameterAddress and it should be present whenever smsf supports Diameter protocol. 
type NetworkNodeDiameterAddress struct {

	// Fully Qualified Domain Name
	Name string `json:"name"`

	// Fully Qualified Domain Name
	Realm string `json:"realm"`
}