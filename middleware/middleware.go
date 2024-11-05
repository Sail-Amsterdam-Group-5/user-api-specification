package middleware

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// MockAuthMiddleware injects a mock user role into the context for testing purposes
func MockAuthMiddleware(role string) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Inject the mock role into the context
        c.Set("userRole", role)
        c.Next()
    }
}

func CheckScope(requiredScope string) gin.HandlerFunc {
    return func(c *gin.Context) {
        scopesHeader := c.GetHeader("X-User-Scopes")
        if scopesHeader != requiredScope {
            c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
            c.Abort()
            return
        }
        c.Next()
    }
}