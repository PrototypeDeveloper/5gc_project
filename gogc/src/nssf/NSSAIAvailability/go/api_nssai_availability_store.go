/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NSSAIAvailabilityOptions - Discover communication options supported by NSSF for NSSAI Availability
func NSSAIAvailabilityOptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
