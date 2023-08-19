/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ProSeCapability - Indicate the supported ProSe Capability by the PCF.
type ProSeCapability struct {

	ProseDirectDiscovey bool `json:"proseDirectDiscovey,omitempty"`

	ProseDirectCommunication bool `json:"proseDirectCommunication,omitempty"`

	ProseL2UetoNetworkRelay bool `json:"proseL2UetoNetworkRelay,omitempty"`

	ProseL3UetoNetworkRelay bool `json:"proseL3UetoNetworkRelay,omitempty"`

	ProseL2RemoteUe bool `json:"proseL2RemoteUe,omitempty"`

	ProseL3RemoteUe bool `json:"proseL3RemoteUe,omitempty"`

	ProseL2UetoUeRelay bool `json:"proseL2UetoUeRelay,omitempty"`

	ProseL3UetoUeRelay bool `json:"proseL3UetoUeRelay,omitempty"`

	ProseL2EndUe bool `json:"proseL2EndUe,omitempty"`

	ProseL3EndUe bool `json:"proseL3EndUe,omitempty"`
}