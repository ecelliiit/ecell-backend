package routes

import "github.com/gin-gonic/gin"

const (
	POST   = "POST"
	GET    = "GET"
	PUT    = "PUT"
	DELETE = "DELETE"
	PATCH  = "PATCH"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Auth-Token, X-Amzn-Trace-Id, x-amzn-RequestId, x-amz-apigw-id, x-amzn-ErrorType, x-amzn-ErrorMessage, Date")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Auth-Token, X-Amzn-Trace-Id, x-amzn-RequestId, x-amz-apigw-id, x-amzn-ErrorType, x-amzn-ErrorMessage, Date")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		// Handle browser preflight requests, where it asks for allowed origin.
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func BuildRoutes() *gin.Engine {
	// Add routes here
	engine := gin.New()

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(CORS())

	//intantiate all routes
	InitSubscriberRoutes(engine)
	return engine
}
