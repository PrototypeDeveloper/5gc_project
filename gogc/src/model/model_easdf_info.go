/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// EasdfInfo - Information of an EASDF NF Instance
type EasdfInfo struct {

	SNssaiEasdfInfoList []SnssaiEasdfInfoItem `json:"sNssaiEasdfInfoList,omitempty"`

	EasdfN6IpAddressList []IpAddr `json:"easdfN6IpAddressList,omitempty"`

	UpfN6IpAddressList []IpAddr `json:"upfN6IpAddressList,omitempty"`

	// A map of InterfaceUpfInfoItems containing the N6 tunnelling information for establishing  a N6 tunnel between the V-UPF and the V-EASDF, where a valid JSON string serves as key. 
	N6TunnelInfoList map[string]InterfaceUpfInfoItem `json:"n6TunnelInfoList,omitempty"`

	DnsSecurityProtocols []DnsSecurityProtocol `json:"dnsSecurityProtocols,omitempty"`
}
