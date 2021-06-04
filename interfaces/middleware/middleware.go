package middleware

import (
	"bytes"
	"ddd-demo/infrastructure/auth"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc  {
	return func(context *gin.Context) {
		isValid := auth.TokenValid(context.Request)
		if !isValid {
			context.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error": "token invalid",
			})
		}
		context.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc  {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

//Avoid a large file from loading into memory
//If the file size is greater than 8MB dont allow it to even load into memory and waste our time.
func MaxSizeAllowed(n int64) gin.HandlerFunc  {
	return func(context *gin.Context) {
		context.Request.Body = http.MaxBytesReader(context.Writer, context.Request.Body, n)
		buff, errRead := context.GetRawData()
		if errRead != nil {
			//c.JSON(http.StatusRequestEntityTooLarge, "too large")
			context.JSON(http.StatusRequestEntityTooLarge, gin.H{
				"status": http.StatusRequestEntityTooLarge,
				"upload_err": "too large: upload an image less than 8mb",
			})
			context.Abort()
			return
		}
		buf := bytes.NewBuffer(buff)
		context.Request.Body = ioutil.NopCloser(buf)
	}
}