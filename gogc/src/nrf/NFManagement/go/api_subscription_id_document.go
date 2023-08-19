/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RemoveSubscription - Deletes a subscription
func RemoveSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateSubscription - Updates a subscription
func UpdateSubscription(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
