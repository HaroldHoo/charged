/**
 * Copyright 2018 harold. All rights reserved.
 * Filename: ./dao/dao_datas.go
 * Author: harold
 * Mail: mail@yaolong.me
 * Date: 2018-04-13
 */

package dao
import(
	"log"
	"sync"
	"time"
	"errors"
	"bytes"
	"crypto/md5"
	"fmt"
)

type Datas struct {
	pkeyMaps		*map[int64]map[string]string
	pkeys			*map[string]map[string]*[]int64
	mu				sync.Mutex
	time			time.Time
	cacheDatasMap	map[string]map[int64]map[string]string
	cacheDatasSlice	map[string][]map[string]string
}

var (
	ErrParamNotEven = errors.New("param not even number")
)

func NewDatas(dao *Dao, query string, pkey string) (ret *Datas, err error) {
	start := time.Now()
	d := &Datas{}
	d.mu.Lock()
	defer d.mu.Unlock()

	d.pkeyMaps, err = dao.getPkeyMapsFromQuery(query, pkey)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	end := time.Now()
	latency := end.Sub(start)
	log.Printf("%.6f dbname:%s,sql:\"%s\"",
		latency.Seconds(),
		dao.dbname,
		query,
	)

	pkeys := make(map[string]map[string]*[]int64)
	d.pkeys = &pkeys

	d.cacheDatasMap = make(map[string]map[int64]map[string]string)
	d.cacheDatasSlice = make(map[string][]map[string]string)
	d.time = time.Now()
	return d, nil
}

func (datas *Datas) GetDatasSlice(params ...string) (ret []map[string]string, err error){
	if datas == nil {
		return nil, ErrNilPtr
	}

	plen := len(params)
	if plen & 1 == 1 {
		return nil, ErrParamNotEven
	}

	var (
		kk string
		md5buf bytes.Buffer
		md5key string
	)
	md5buf = bytes.Buffer{}
	ps := make(map[string]string)
	for k,v := range params {
		if k&1 == 0 {
			kk = v
			ps[kk] = ""
		}else{
			ps[kk] = v
		}
		md5buf.WriteString(v)
	}
	md5key = fmt.Sprintf("%x", md5.Sum(md5buf.Bytes()))
	if datas.cacheDatasSlice[md5key] != nil {
		ret = datas.cacheDatasSlice[md5key]
		return ret, nil
	}

	datasMap, e := datas.GetDatasMap(params...)
	if e != nil {
		return nil, e
	}
	for _, dataval := range datasMap {
		ret = append(ret, dataval)
	}

	datas.mu.Lock()
	datas.cacheDatasSlice[md5key] = ret
	datas.mu.Unlock()

	return
}

func (datas *Datas) GetDatasMap(params ...string) (ret map[int64]map[string]string, err error){
	if datas == nil {
		return nil, ErrNilPtr
	}

	plen := len(params)
	if plen == 0 {
		ret = *datas.pkeyMaps
		return
	}

	if plen & 1 == 1 {
		return nil, ErrParamNotEven
	}

	var (
		kk string
		md5buf bytes.Buffer
		md5key string
	)
	md5buf = bytes.Buffer{}
	ps := make(map[string]string)
	for k,v := range params {
		if k&1 == 0 {
			kk = v
			ps[kk] = ""
		}else{
			ps[kk] = v
		}
		md5buf.WriteString(v)
	}
	md5key = fmt.Sprintf("%x", md5.Sum(md5buf.Bytes()))
	ret = make(map[int64]map[string]string)
	if datas.cacheDatasMap[md5key] != nil {
		ret = datas.cacheDatasMap[md5key]
		return ret, nil
	}

	var i int64
	pk := make(map[int64]int64)
	for k,v := range ps {
		var kp *[]int64
		kp, err = datas.GetPkeysByColName(k,v)
		if err != nil {
			return nil ,err
		}

		for _,vv := range *kp {
			pk[vv]++
		}
		i++
	}

	for k,v := range pk {
		if v == i {
			ret[k] = (*datas.pkeyMaps)[k]
		}
	}

	datas.mu.Lock()
	datas.cacheDatasMap[md5key] = ret
	datas.mu.Unlock()

	return
}

func (datas *Datas) GetPkeysByColName(colName string, match string) (ret *[]int64, err error){
	if datas == nil {
		return nil, ErrNilPtr
	}
	if (*datas.pkeys)[colName][match] == nil {
		datas.mu.Lock()
		defer datas.mu.Unlock()
		keys := make(map[string]*[]int64)
		keys[match], err = getPkeysByColName(datas.pkeyMaps, colName, match)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		(*datas.pkeys)[colName] = keys
	}

	return (*datas.pkeys)[colName][match], nil
}

