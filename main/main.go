package main

import (
	"fmt"
	"github.com/HaroldHoo/srvmanager"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"time"
	"charged/route"
)

func main() {
	m := srvmanager.New()

	gin.DefaultWriter = m.GetAccessLogWriter()
	gin.DefaultErrorWriter = m.GetErrorLogWriter()

	router := gin.New()
	router.Use(logger(gin.DefaultWriter), gin.Recovery())

	route.RegisterRoutes(router)

	srv := &http.Server{
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	m.Run(srv)
}

func logger(out io.Writer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		referer := c.Request.Header.Get("referer")
		user_agent := c.Request.Header.Get("user-agent")
		x_clientip := c.Request.Header.Get("x-clientip")
		// comment := c.Errors.ByType(gin.ErrorTypePrivate).String()

		if raw != "" {
			path = path + "?" + raw
		}
		fmt.Fprintf(out, "%s [%s] \"%s %s\" %3d \"%s\" \"%s\" \"%s\" %.6f\n",
			clientIP,
			end.Format(time.RFC3339),
			method,
			path,
			statusCode,
			referer,
			user_agent,
			x_clientip,
			latency.Seconds(),
		)
	}
}
