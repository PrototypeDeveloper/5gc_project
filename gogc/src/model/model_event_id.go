/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// EventId - Represents the analytics type.   Possible values are: - LOAD_LEVEL_INFORMATION: Represent the analytics of load level information of corresponding   network slice. - NETWORK_PERFORMANCE: Represent the analytics of network performance information. - NF_LOAD: Indicates that the event subscribed is NF Load. - SERVICE_EXPERIENCE: Represent the analytics of service experience information of the   specific applications. - UE_MOBILITY: Represent the analytics of UE mobility. - UE_COMMUNICATION: Represent the analytics of UE communication. - QOS_SUSTAINABILITY: Represent the analytics of QoS sustainability information in the   certain area. - ABNORMAL_BEHAVIOUR: Indicates that the event subscribed is abnormal behaviour information. - USER_DATA_CONGESTION: Represent the analytics of the user data congestion in the certain   area. - NSI_LOAD_LEVEL: Represent the analytics of Network Slice and the optionally associated   Network Slice Instance. - SM_CONGESTION: Represent the analytics of Session Management congestion control experience   information for specific DNN and/or S-NSSAI. - DISPERSION: Represents the analytics of dispersion. - RED_TRANS_EXP: Represents the analytics of Redundant Transmission Experience. - WLAN_PERFORMANCE: Represents the analytics of WLAN performance. - DN_PERFORMANCE: Represents the analytics of DN performance. - PDU_SESSION_TRAFFIC: Represents the analytics of PDU Session traffic. - E2E_DATA_VOL_TRANS_TIME: Represents the analytics of E2E data volume transfer time. - MOVEMENT_BEHAVIOUR: Represents the analytics of the Movement Behaviour information. - LOC_ACCURACY: Represents the analytics of location accuracy. - RELATIVE_PROXIMITY: Represents the analytics of Relative Proximity information. 
type EventId struct {
}
