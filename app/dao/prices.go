/**
 * Copyright 2018 harold. All rights reserved.
 * Filename: rent_rate.go
 * Author: harold
 * Mail: mail@yaolong.me
 * Date: 2018-04-11
 */

package dao

import ()

var (
	RentRate           *Datas
	OrderRentPrice     *Datas
	FixedProductDefine *Datas
	OrderFixedPrice    *Datas
	YopRentPrice       *Datas
	YopFixedPrice      *Datas
)

// rent_rate
func NewRentRate() (ret *Datas, err error) {
	ret, err = NewDatas(yc_crm_common, "SELECT * from rent_rate;", "rent_rate_id")
	return
}
func GetRentRate() (ret *Datas, err error) {
	if RentRate == nil {
		RentRate, err = NewRentRate()
	}
	return RentRate, err
}
func UpdateRentRate() (err error) {
	RentRate, err = NewRentRate()
	return
}

// order_rent_price
func NewOrderRentPrice() (ret *Datas, err error) {
	ret, err = NewDatas(yc_crm_common, "SELECT * from order_rent_price;", "order_rent_price_id")
	return
}
func GetOrderRentPrice() (ret *Datas, err error) {
	if OrderRentPrice == nil {
		OrderRentPrice, err = NewOrderRentPrice()
	}
	return OrderRentPrice, err
}
func UpdateOrderRentPrice() (err error) {
	OrderRentPrice, err = NewOrderRentPrice()
	return
}

// fixed_product_define
func NewFixedProductDefine() (ret *Datas, err error) {
	ret, err = NewDatas(yc_crm_common, "SELECT * from fixed_product_define;", "fixed_product_define_id")
	return
}
func GetFixedProductDefine() (ret *Datas, err error) {
	if FixedProductDefine == nil {
		FixedProductDefine, err = NewFixedProductDefine()
	}
	return FixedProductDefine, err
}
func UpdateFixedProductDefine() (err error) {
	FixedProductDefine, err = NewFixedProductDefine()
	return
}

// order_fixed_price
func NewOrderFixedPrice() (ret *Datas, err error) {
	ret, err = NewDatas(yc_crm_common, "SELECT * from order_fixed_price;", "order_fixed_price_id")
	return
}
func GetOrderFixedPrice() (ret *Datas, err error) {
	if OrderFixedPrice == nil {
		OrderFixedPrice, err = NewOrderFixedPrice()
	}
	return OrderFixedPrice, err
}
func UpdateOrderFixedPrice() (err error) {
	OrderFixedPrice, err = NewOrderFixedPrice()
	return
}

// yop_rent_price
func NewYopRentPrice() (ret *Datas, err error) {
	ret, err = NewDatas(yc_crm_common, "SELECT * from yop_rent_price;", "yop_rent_price_id")
	return
}
func GetYopRentPrice() (ret *Datas, err error) {
	if YopRentPrice == nil {
		YopRentPrice, err = NewYopRentPrice()
	}
	return YopRentPrice, err
}
func UpdateYopRentPrice() (err error) {
	YopRentPrice, err = NewYopRentPrice()
	return
}

// yop_fixed_price
func NewYopFixedPrice() (ret *Datas, err error) {
	ret, err = NewDatas(yc_crm_common, "SELECT * from yop_fixed_price;", "yop_fixed_price_id")
	return
}
func GetYopFixedPrice() (ret *Datas, err error) {
	if YopFixedPrice == nil {
		YopFixedPrice, err = NewYopFixedPrice()
	}
	return YopFixedPrice, err
}
func UpdateYopFixedPrice() (err error) {
	YopFixedPrice, err = NewYopFixedPrice()
	return
}
