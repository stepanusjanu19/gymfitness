package routes

import (
	"api/lib/utils/response"
	"api/src/controllers"

	// "api/src/middleware"
	"net/http"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	apiGroup := route.Group("/api")
	{
		apiV1 := apiGroup.Group("/v1")
		{
			apiV1.POST("/login", controllers.LoginUser)
			apiV1.POST("/signup", controllers.RegisterUser)
			apiV1.POST("/verify-user", controllers.ValidateOTPAndLogin)
		}
	}
}
