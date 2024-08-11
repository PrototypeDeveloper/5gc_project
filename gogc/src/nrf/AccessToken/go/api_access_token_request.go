/*
 * NRF OAuth2
 *
 * NRF OAuth2 Authorization.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"github.com/gin-gonic/gin"
)

type AccessTokenRequestAPI struct {
}

// Post /oauth2/token
// Access Token Request 
func (api *AccessTokenRequestAPI) AccessTokenRequest(c *gin.Context) {
	// Your handler implementation
	c.JSON(200, gin.H{"status": "OK"})
}

