// Package router creates a controller layer to manage client requests
package router

import (
	"review-manager/api/src/customerror"
	"review-manager/api/src/environment"

	"github.com/gin-gonic/gin"
)

// InitializeRouter configuration with endpoints
func InitializeRouter(env environment.Environment) {

	router := gin.Default()

	group := router.Group("/api/v1")
	{
		group.POST("/reviews", createReview(env))

		group.DELETE("/reviews/:id", deleteReview(env))

		group.GET("/reviews/orders/:orderID", findReviewByOrderID(env))
	}

	router.Run()
}

func abortWithCustomError(c *gin.Context, defaultStatus int, err error) {
	if apiError, ok := err.(*customerror.Error); ok && apiError.StatusCode() != 0 {
		c.AbortWithStatusJSON(apiError.StatusCode(), gin.H{"error": apiError.Error()})
	} else {
		c.AbortWithStatusJSON(defaultStatus, gin.H{"error": err.Error()})
	}
}
