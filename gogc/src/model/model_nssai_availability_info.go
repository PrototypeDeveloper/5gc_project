/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// NssaiAvailabilityInfo - This contains the Nssai availability information requested by the AMF
type NssaiAvailabilityInfo struct {

	SupportedNssaiAvailabilityData []SupportedNssaiAvailabilityData `json:"supportedNssaiAvailabilityData"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	AmfSetId string `json:"amfSetId,omitempty"`
}
