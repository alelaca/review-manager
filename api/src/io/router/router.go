package router

import (
	"github.com/gin-gonic/gin"
)

func InitializeRouter() {

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("ping", ping)
	}

	router.Run()
}
