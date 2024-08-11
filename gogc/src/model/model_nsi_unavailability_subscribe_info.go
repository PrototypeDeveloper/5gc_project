/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// NsiUnavailabilitySubscribeInfo - Present if the NF service consumer subscribes to events related to Network Slice Instance Replacement. 
type NsiUnavailabilitySubscribeInfo struct {

	NsiToSubscribe []string `json:"nsiToSubscribe,omitempty"`

	SnssaiToSubscribe []Snssai `json:"snssaiToSubscribe,omitempty"`
}