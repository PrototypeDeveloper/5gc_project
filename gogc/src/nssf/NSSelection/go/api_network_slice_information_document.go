/*
 * NSSF NS Selection
 *
 * NSSF Network Slice Selection Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NSSelectionGet - Retrieve the Network Slice Selection Information
func NSSelectionGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
