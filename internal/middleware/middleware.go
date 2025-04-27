package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"scaf-gin/config"
	"scaf-gin/internal/core"
	"scaf-gin/internal/helper"
)

// BasicAuth is a middleware that checks for Basic Authentication credentials.
// If the credentials are incorrect, it returns an Unauthorized status.
func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, pass, ok := c.Request.BasicAuth()
		if !ok || user != config.BasicAuthUser || pass != config.BasicAuthPass {
			c.Header("WWW-Authenticate", "Basic realm=Authorization Required")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}

// ApiAuth is a middleware that validates the JWT token for API access.
// If the token is invalid, it returns an Unauthorized error in JSON format.
func ApiAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := helper.GetAccessToken(c)
		payload, err := core.Auth.VerifyAccessToken(token)
		if err != nil {
			c.Error(core.ErrUnauthorized)
			c.Abort()
			return
		}

		c.Set(helper.CONTEXT_KEY_PAYLOAD, payload)
		c.Next()
	}
}

// ApiErrorHandler is a middleware that handles API errors.
// It checks for specific error types and returns the appropriate HTTP status and message.
func ApiErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		var status int
		var resp gin.H

		switch {
		case errors.Is(err, core.ErrBadRequest):
			status = http.StatusBadRequest
			if appErr, ok := err.(*core.AppError); ok {
				resp = gin.H{
					"message": appErr.Error(),
					"details": appErr.Details(),
				}
			} else {
				resp = gin.H{
					"message": err.Error(),
					"details": []map[string]any{},
				}
			}
		case errors.Is(err, core.ErrUnauthorized):
			status = http.StatusUnauthorized
			resp = gin.H{"message": err.Error()}
		case errors.Is(err, core.ErrForbidden):
			status = http.StatusForbidden
			resp = gin.H{"message": err.Error()}
		case errors.Is(err, core.ErrNotFound):
			status = http.StatusNotFound
			resp = gin.H{"message": err.Error()}
		case errors.Is(err, core.ErrConflict):
			status = http.StatusConflict
			resp = gin.H{"message": err.Error()}
		case errors.Is(err, core.ErrUnexpected):
			core.Logger.Error(err.Error())
			status = http.StatusInternalServerError
			resp = gin.H{"message": err.Error()}
		default:
			core.Logger.Error(err.Error())
			status = http.StatusInternalServerError
			resp = gin.H{"message": err.Error()}
		}

		c.JSON(status, resp)
	}
}
