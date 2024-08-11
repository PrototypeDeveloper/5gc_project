/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// RestrictedSnssai - This contains the restricted SNssai information per PLMN
type RestrictedSnssai struct {

	HomePlmnId PlmnId `json:"homePlmnId"`

	SNssaiList []ExtSnssai `json:"sNssaiList"`

	HomePlmnIdList []PlmnId `json:"homePlmnIdList,omitempty"`

	RoamingRestriction bool `json:"roamingRestriction,omitempty"`
}
