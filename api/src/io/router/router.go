// Package router creates a controller layer to manage client requests
package router

import (
	"review-manager/api/src/customerror"
	"review-manager/api/src/environment"

	"github.com/gin-gonic/gin"
)

// Initializes router configuration with endpoints
func InitializeRouter(env environment.Environment) {

	router := gin.Default()

	group := router.Group("/api/v1")
	{
		group.GET("/ping", ping)

		group.GET("/reviews/:id", findReviewByID(env))

		group.POST("/reviews", createReview(env))
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
