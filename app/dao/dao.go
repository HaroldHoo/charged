/**
 * Copyright 2018 harold. All rights reserved.
 * Filename: rent_rate.go
 * Author: harold
 * Mail: mail@yaolong.me
 * Date: 2018-04-11
 */

package dao

import (
	"errors"
	"strconv"
	"database/sql"
	"fmt"
	"sync"
	"charged/app/utils/config"
	_ "github.com/go-sql-driver/mysql"
)

type Dao struct {
	mu     sync.Mutex
	db     *sql.DB

	dbname   string
	host     string
	port     string
	username string
	password string
	charset  string
}

var yc_crm_common *Dao

func init() {
	var err error
	yc_crm_common = &Dao{}
	yc_crm_common.mu.Lock()
	defer yc_crm_common.mu.Unlock()

	var buf interface{}
	buf, _ = config.Get("mysql", "yc_crm_common", "master", "dbname")
	yc_crm_common.dbname = buf.(string)
	buf, _ = config.Get("mysql", "yc_crm_common", "master", "host")
	yc_crm_common.host= buf.(string)
	buf, _ = config.Get("mysql", "yc_crm_common", "master", "port")
	yc_crm_common.port= buf.(string)
	buf, _ = config.Get("mysql", "yc_crm_common", "master", "username")
	yc_crm_common.username= buf.(string)
	buf, _ = config.Get("mysql", "yc_crm_common", "master", "password")
	yc_crm_common.password= buf.(string)
	buf, _ = config.Get("mysql", "yc_crm_common", "master", "charset")
	yc_crm_common.charset= buf.(string)

	yc_crm_common.db, err = sql.Open("mysql", fmt.Sprintf(
		"%s:@tcp(%s:%s)/%s?charset=%s&timeout=%s",
		yc_crm_common.username,
		yc_crm_common.host,
		yc_crm_common.port,
		yc_crm_common.dbname,
		yc_crm_common.charset,
		"3s",
	))
	if err != nil {
		panic(err)
	}
}

var (
	ErrNilPtr = errors.New("destination pointer is nil")
)

func getPkeysByColName(data *map[int64]map[string]string, colName string, match string) (ret *[]int64, err error){
	if data == nil {
		return nil, ErrNilPtr
	}

	var slice []int64
	ret = &slice
	for pid,mp := range *data {
		if mp[colName] == match {
			slice = append(slice, pid)
		}
	}

	return
}

func (dao *Dao) getPkeyMapsFromQuery(query string, pkey string) (ret *map[int64]map[string]string, err error) {
	dao.mu.Lock()
	defer dao.mu.Unlock()

	rows := &sql.Rows{}
	if rows, err = dao.db.Query(query); err != nil {
		return nil, err
	}

	var cn []string
	if cn, err = rows.Columns(); err != nil {
		return nil, err
	}

	buf := make([]string, len(cn))
	dest := make([]interface{}, len(cn))
	for k, _ := range buf {
		dest[k] = &buf[k]
	}

	tmp := make(map[int64]map[string]string)
	ret = &tmp
	for rows.Next() {
		rows.Scan(dest...)

		mapdata := make(map[string]string)
		for k, v := range cn {
			mapdata[v] = buf[k]
		}

		p,err := strconv.ParseInt(mapdata[pkey], 10, 64)
		if err != nil {
			return nil, err
		}

		(*ret)[p] = mapdata
	}
	return ret, nil
}

func (dao *Dao) getSlicesFromQuery(query string) (ret *[]map[string]string, err error) {
	dao.mu.Lock()
	defer dao.mu.Unlock()

	rows := &sql.Rows{}
	if rows, err = dao.db.Query(query); err != nil {
		return nil, err
	}

	var cn []string
	if cn, err = rows.Columns(); err != nil {
		return nil, err
	}

	buf := make([]string, len(cn))
	dest := make([]interface{}, len(cn))
	for k, _ := range buf {
		dest[k] = &buf[k]
	}

	ret = new([]map[string]string)
	for rows.Next() {
		rows.Scan(dest...)

		mapdata := make(map[string]string)
		for k, v := range cn {
			mapdata[v] = buf[k]
		}
		*ret = append(*ret, mapdata)
	}
	return ret, nil
}
