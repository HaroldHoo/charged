/**
 * Copyright 2018 harold. All rights reserved.
 * Filename: controller.go
 * Author: harold
 * Mail: mail@yaolong.me
 * Date: 2018-04-19
 */

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CODE_SUCCESS      = 200
	CODE_NOT_FOUND    = 404
	CODE_PARAM_EMPTY  = 431
	CODE_PARAM_ERROR  = 432
	CODE_SERVER_ERROR = 500
)

var MSG = map[int]string{
	CODE_SUCCESS:      "success",
	CODE_NOT_FOUND:    "not found",
	CODE_PARAM_EMPTY:  "param empty",
	CODE_PARAM_ERROR:  "param error",
	CODE_SERVER_ERROR: "Internal Server Error",
}

func json(c *gin.Context, objs ...interface{}) {
	out := make(map[string]interface{})

	code := CODE_SUCCESS
	msg := MSG[code]
	if len(objs) == 2 {
		code = objs[1].(int)
		msg = MSG[code]
	}
	if len(objs) == 3 {
		code = objs[1].(int)

		switch objs[2].(type) {
		case string:
			msg = objs[2].(string)
			break
		default:
			msg = fmt.Sprintf("%s", objs[2])
			break
		}
	}

	out["result"] = objs[0]
	out["ret_code"] = code
	out["ret_msg"] = msg
	c.JSON(http.StatusOK, out)
}

func getUrlValuesList(c *gin.Context) (ret map[string][]string) {
	ret = make(map[string][]string)

	var (
		k string
		v []string
	)
	for k, v = range c.Request.URL.Query() {
		ret[k] = v
	}

	//@todo maybe too large
	req := c.Request
	req.ParseForm()
	for k, v = range req.PostForm {
		ret[k] = v
	}

	return
}

func getUrlValues(c *gin.Context) (ret map[string]string) {
	ret = make(map[string]string)

	var (
		k string
		v []string
	)
	for k, v = range c.Request.URL.Query() {
		ret[k] = v[0]
	}

	//@todo maybe too large
	req := c.Request
	req.ParseForm()
	for k, v = range req.PostForm {
		ret[k] = v[0]
	}

	return
}
