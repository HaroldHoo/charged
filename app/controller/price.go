/**
 * Copyright 2018 harold. All rights reserved.
 * Filename: app/controller/price.go
 * Author: harold
 * Mail: mail@yaolong.me
 * Date: 2018-04-17
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"charged/app/service"
	"charged/config"
	"log"
	"fmt"
	"time"
	"strings"
)

func PricePublish(c *gin.Context) {
	log.Printf("received price.publish.")

	conf, e1 := config.Config()
	if e1 != nil {
		json(c, e1)
		return
	}

	buf, e2 := conf.Get("me")
	urls := buf.([]interface{})
	if e2 != nil {
		json(c, e2)
		return
	}

	for _,v := range urls {
		go func(v interface{}){
			client := &http.Client{}
			client.Timeout = 3000 * time.Millisecond
			url := fmt.Sprintf("http://%s%s", v, "/V2/price/publish/notify")
			res, err := client.Get(url)

			if err != nil {
				for i:=0; i<5; i++ {
					time.Sleep(1 * time.Second)
					log.Printf("(%s) Try notifing... \"%s\"", err, url)
					_, errTry := client.Get(url)
					if errTry == nil {
						log.Printf("notify success \"%s\"", url)
						break
					}
				}
				return
			}

			if res.StatusCode != 200 {
				for i:=0; i<5; i++ {
					time.Sleep(1 * time.Second)
					log.Printf("(%3d) Try notifing... \"%s\"", res.StatusCode, url)
					resTry,_ := client.Get(url)
					if resTry.StatusCode == 200 {
						log.Printf("notify success \"%s\"", url)
						break
					}
				}
				return
			}

			log.Printf("notify success \"%s\"", url)
		}(v)
	}

	json(c, "ok")
}

func PricePublishNotify(c *gin.Context) {
	log.Printf("received price.publish.notify.")
	service.PriceCacheClean()
	json(c, "ok")
}

func PriceAll(c *gin.Context) {
	ret, err := service.PriceGet(map[string]string{})
	if err != nil {
		json(c, err)
		return
	}
	json(c, ret)
}

func PriceGet(c *gin.Context) {
	params := getUrlValues(c)

	if len(params) == 0 {
		json(c, []string{}, CODE_PARAM_EMPTY)
		return
	}

	ret, err := service.PriceGet(params)

	if err != nil {
		json(c, []string{}, CODE_SERVER_ERROR, err)
		return
	}

	json(c, ret)
}

//@todo
func GetPassengerPrice(c *gin.Context) {
	var (
		product_type_class, city, car_type_id, fixed_product_id string
	)
	product_type_class, _ = c.GetPostForm("product_type_class")
	if product_type_class == "" {
		json(c, []string{}, CODE_PARAM_ERROR, "product_type_class can not be empty")
		return
	}

	var citys []string
	city, _ = c.GetPostForm("city")
	if city != "" {
		citys = strings.Split(city, ",")
	}

	var car_type_ids []string
	car_type_id, _ = c.GetPostForm("car_type_id")
	if car_type_id != "" {
		car_type_ids = strings.Split(car_type_id, ",")
	}

	var fixed_product_ids []string
	fixed_product_id, _ = c.GetPostForm("fixed_product_id")
	if fixed_product_id != "" {
		fixed_product_ids = strings.Split(fixed_product_id, ",")
	}
	_,_ = car_type_ids, fixed_product_ids

	json(c, citys)
}

