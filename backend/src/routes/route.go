package routes

import (
	"backend/lib/utils/response"
	"backend/src/controllers"
	// "backend/src/middleware"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Init(route *gin.Engine) {
	route.Use(cors.Default())

	limiter := tollbooth.NewLimiter(1, &limiter.ExpirableOptions{
		DefaultExpirationTTL: time.Hour,
	})

	route.GET("/ping", func(ctx *gin.Context) {
		response.JSON(ctx.Writer, http.StatusOK, gin.H{
			"message": "PONG",
		})
	})

	route.GET("/", func(ctx *gin.Context) {
		response.JSON(ctx.Writer, http.StatusOK, gin.H{
			"message": "...................WELCOME API...................",
		})
	})

	route.Use(func(ctx *gin.Context) {
		httpErr := tollbooth.LimitByRequest(limiter, ctx.Writer, ctx.Request)
		if httpErr != nil {
			response.JSON(ctx.Writer, httpErr.StatusCode, gin.H{
				"error": "Too many requests, please try again later",
			})
			ctx.Abort()
			return
		}
	})

	apiGroup := route.Group("/api/v1")
	{
		apiGroup.POST("/login", controllers.LoginUser)
		apiGroup.POST("/signup", controllers.RegisterUser)
	}
}
