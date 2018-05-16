package ycconfig

/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -lycclient -L/lib64
#include <stdlib.h>
#include "ycconfig/yc_config_client.h"
*/
import "C"

import (
	"errors"
	"log"
)

const (
	ycconfig_cache_filename = "/home/y/var/ycconfig/mmap_cache.conf"
)

func GetCurrentConfigEnv() string {
	return C.GoString(C.yc_config_env())
}

func GetStringConfig(key string) (string, error) {
	var cs C.StringConfig
	if C.yc_config_get_string(C.CString(key), &cs) != 0 {
		log.Println("Get config failed, key: ", key)
		return "", errors.New("Get config failed, key: " + key)
	}

	return C.GoString(cs.config), nil
}

func GetListConfig(key string) ([]string, error) {
	var cs C.ListConfig
	if C.yc_config_get_list(C.CString(key), &cs) != 0 {
		log.Println("Get config failed, key: ", key)
		return nil, errors.New("Get config failed, key: " + key)
	}

	res := make([]string, cs.size)
	for i := 0; i < int(cs.size); i++ {
		res[i] = C.GoString(cs.config[i])
	}
	return res, nil
}

func GetMapConfig(key string) (map[string]string, error) {
	var cs C.MapConfig
	if C.yc_config_get_map(C.CString(key), &cs) != 0 {
		log.Println("Get config failed, key: ", key)
		return nil, errors.New("Get config failed, key: " + key)
	}

	res := make(map[string]string, cs.size)
	for i := 0; i < int(cs.size); i++ {
		res[C.GoString(cs.config[i].key)] = C.GoString(cs.config[i].value)
	}
	return res, nil
}

func init() {
	if int(C.yc_config_init(C.CString(ycconfig_cache_filename))) != 0 {
		log.Println("Init ycconfig cache file failed!")
	}
}

