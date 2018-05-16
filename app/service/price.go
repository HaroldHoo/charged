/**
 * Copyright 2018 harold. All rights reserved.
 * Filename: app/service/price.go
 * Author: harold
 * Mail: mail@yaolong.me
 * Date: 2018-04-17
 */

package service

import (
	"charged/app/dao"
	"errors"
	"sort"
)

var (
	ErrParamsEmpty = errors.New("params can not empty.")
	ErrTableWrong  = errors.New("param '_table' was wrong.")
)

func PriceCacheClean() {
	go func() {
		dao.UpdateRentRate()
		dao.UpdateOrderRentPrice()
		dao.UpdateFixedProductDefine()
		dao.UpdateOrderFixedPrice()
		dao.UpdateYopRentPrice()
		dao.UpdateYopFixedPrice()
	}()
}

func PriceGet(params map[string]string) (ret interface{}, err error) {
	var (
		d      *dao.Datas
		e      error
		maybe1 = false
	)

	switch "0" {
	case params["car_type_id"]:
		delete(params, "car_type_id")
	case params["product_type_id"]:
		delete(params, "product_type_id")
	case params["fixed_product_id"]:
		delete(params, "fixed_product_id")
	}

	switch params["_table"] {
	case "rent_rate":
		if params["_maybe1"] != "" && params["status"] == "1" && params["car_type_id"] != "" && params["city"] != "" && params["product_type_id"] != "" && params["is_segment"] == "0" {
			maybe1 = true
		}
		d, e = dao.GetRentRate()
	case "order_rent_price":
		if params["_maybe1"] != "" && params["status"] == "1" && params["car_type_id"] != "" && params["city"] != "" && params["product_type_id"] != "" && params["is_segment"] == "0" {
			maybe1 = true
		}
		d, e = dao.GetOrderRentPrice()
	case "fixed_product_define":
		if params["_maybe1"] != "" && params["status"] == "1" && params["car_type_id"] != "" && params["fixed_product_id"] != "" {
			maybe1 = true
		}
		d, e = dao.GetFixedProductDefine()
	case "order_fixed_price":
		if params["_maybe1"] != "" && params["status"] == "1" && params["car_type_id"] != "" && params["fixed_product_id"] != "" {
			maybe1 = true
		}
		d, e = dao.GetOrderFixedPrice()
	case "yop_rent_price":
		if params["_maybe1"] != "" && params["status"] == "1" && params["car_type_id"] != "" && params["city"] != "" && params["product_type_id"] != "" {
			maybe1 = true
		}
		d, e = dao.GetYopRentPrice()
	case "yop_fixed_price":
		if params["_maybe1"] != "" && params["status"] == "1" && params["car_type_id"] != "" && params["fixed_product_id"] != "" {
			maybe1 = true
		}
		d, e = dao.GetYopFixedPrice()
	default:
		return nil, ErrTableWrong
	}

	if e != nil {
		return nil, e
	}

	delete(params, "_table")
	delete(params, "_maybe1")

	if len(params) == 0 {
		return d.GetDatasSlice()
	}

	// sort
	sortbuf := make([]string, 0, len(params))
	for k, v := range params {
		if v == "" {
			delete(params, k)
			continue
		}
		sortbuf = append(sortbuf, k)
	}
	sort.Strings(sortbuf)

	var ps []string
	for _, v := range sortbuf {
		ps = append(ps, v, params[v])
	}
	// sort

	data, er := d.GetDatasSlice(ps...)
	if er != nil {
		return nil, er
	}

	if len(data) == 0 {
		return []string{}, nil
	}

	if maybe1 == true {
		if len(data) != 0 {
			return data[0], nil
		}
	}

	return data, nil
}
