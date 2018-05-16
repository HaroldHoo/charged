/**
 * Copyright 2018 harold. All rights reserved.
 * Filename: route/map.go
 * Author: harold
 * Mail: mail@yaolong.me
 * Date: 2018-04-10
 */

package route

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"charged/app/controller"
)

type (
	route struct {
		method string
		path string
		handlerFunc gin.HandlerFunc
	}
)

func RegisterRoutes(router *gin.Engine) {
	for _,route := range routes {
		router.Handle(route.method, route.path, route.handlerFunc)
	}
}

var (
	routes = []route{
		{"GET", "/", func(c *gin.Context) {
			c.String(http.StatusOK, "Welcome to charged!")
		}},
		// {"GET", "/V2/price/get/all", func(c *gin.Context) {
		// 	controller.PriceAll(c)
		// }},

		// price.get
		{"GET", "/V2/price/get", func(c *gin.Context) {
			controller.PriceGet(c)
		}},
		{"POST", "/V2/price/get", func(c *gin.Context) {
			controller.PriceGet(c)
		}},

		{"POST", "/V2/Data/Price/getPassengerPrice", func(c *gin.Context) {
			controller.GetPassengerPrice(c)
		}},

		// publish & notify
		{"GET", "/V2/price/publish", func(c *gin.Context) {
			controller.PricePublish(c)
		}},
		{"GET", "/V2/price/publish/notify", func(c *gin.Context) {
			controller.PricePublishNotify(c)
		}},
	}
)

