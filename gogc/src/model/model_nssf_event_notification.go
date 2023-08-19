/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// NssfEventNotification - This contains the notification for created event subscription
type NssfEventNotification struct {

	SubscriptionId string `json:"subscriptionId"`

	AuthorizedNssaiAvailabilityData []AuthorizedNssaiAvailabilityData `json:"authorizedNssaiAvailabilityData"`

	// Indicate the impacted S-NSSAIs, the current status for each reported S-NSSAI, and  if available the alternative S-NSSAI per impacted S-NSSAI for the S-NSSAIs that are  reported as being not available. 
	AltNssai []SnssaiReplaceInfo `json:"altNssai,omitempty"`
}