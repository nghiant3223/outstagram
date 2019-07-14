package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Verify that a user is getting his own resource
func RedirectToDuplicateRoute(r *gin.Engine, to string) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		if authUserID, ok := c.Get("userID"); ok {
			c.Request.URL.Path = fmt.Sprintf(to, authUserID)
			r.HandleContext(c)
		}
	}
}
