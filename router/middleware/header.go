package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// NoCache is a middleware function that appends headers
// to prevent the client from caching the HTTP response.
func NoCache(c *gin.Context) {
	c.Header("Cache-Control","no-cache,no-stroe,max-age=0,must-revalidate,value")
	c.Header("Expires","Thu,01 Jan 1970 00:00:00 GHT")
	c.Header("Last-Modified",time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

// Options is a middleware function that appends headers
// for options requests and aborts then exits the middleware
// chain and ends the request.
func Option(c *gin.Context) {
	if c.Request.Method != "OPTIONS"{
		c.Next()
	}else {
		c.Header("Access-Control-Allow-Origin","*")
		c.Header("Access-Control-Allow-Methods","GET,POST,PUT,PATCH,DELETE,OPTION")
		c.Header("Access-control-Allow-Headers","authorization,origin,content-type,accept")
		c.Header("Content-type","application/json")
		c.AbortWithStatus(http.StatusOK)
	}
}


// Secure is a middleware function that appends security
// and resource access headers.
func Secure(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin","*")
	c.Header("X-Frame-Options","DENY")
	c.Header("X-Content-Type-Options","nosniff")
	c.Header("X-XSS-Protection","1;mode-block")
	c.Header("Content-type","*")
	c.Header("Authorization","*")
	if c.Request.TLS !=nil {
		c.Header("Strict-Transport-Security","max-age=31536000")
	}
}