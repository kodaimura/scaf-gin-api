package helper

import (
	"net/http"
	"strings"

	"scaf-gin/config"
	"scaf-gin/internal/core"

	"github.com/gin-gonic/gin"
)

// GetAccessToken retrieves the access token from cookie or Authorization header.
// If the cookie is not found, it attempts to extract a Bearer token from the header.
func GetAccessToken(c *gin.Context) string {
	token, err := c.Cookie(COOKIE_KEY_ACCESS_TOKEN)
	if err == nil {
		return token
	}

	bearer := c.GetHeader("Authorization")
	if strings.HasPrefix(bearer, "Bearer ") {
		return strings.TrimSpace(bearer[7:])
	}

	return ""
}

// GetRefreshToken retrieves the access token from cookie.
func GetRefreshToken(c *gin.Context) string {
	token, err := c.Cookie(COOKIE_KEY_REFRESH_TOKEN)
	if err == nil {
		return token
	}

	return ""
}

// SetAccessTokenCookie sets the access token cookie in the response
func SetAccessTokenCookie(c *gin.Context, accessToken string) {
	maxAge := config.AccessTokenExpiresSeconds
	if accessToken == "" {
		maxAge = -1
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		COOKIE_KEY_ACCESS_TOKEN, 
		accessToken, 
		maxAge, 
		"/", config.AppHost, 
		config.CookieAccessSecure, 
		config.CookieAccessHttpOnly,
	)
}

// SetRefreshTokenCookie sets the refresh token cookie in the response
func SetRefreshTokenCookie(c *gin.Context, refreshToken string) {
	maxAge := config.RefreshTokenExpiresSeconds
	if refreshToken == "" {
		maxAge = -1
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		COOKIE_KEY_REFRESH_TOKEN, 
		refreshToken, 
		maxAge,
		"/", config.AppHost, 
		config.CookieRefreshSecure, 
		config.CookieRefreshHttpOnly,
	)
}

// GetPayload retrieves the AuthPayload from the context.
// Returns an empty AuthPayload if the context value is not present or invalid.
func GetPayload(c *gin.Context) core.AuthPayload {
	pl, ok := c.Get(CONTEXT_KEY_PAYLOAD)
	if !ok {
		return core.AuthPayload{}
	}

	if payload, ok := pl.(core.AuthPayload); ok {
		return payload
	}
	return core.AuthPayload{}
}

// GetAccountId returns the account ID from the AuthPayload in the context.
func GetAccountId(c *gin.Context) int {
	return GetPayload(c).AccountId
}

// GetAccountName returns the account name from the AuthPayload in the context.
func GetAccountName(c *gin.Context) string {
	return GetPayload(c).AccountName
}
