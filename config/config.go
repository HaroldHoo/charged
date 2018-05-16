/**
 * Copyright 2018 harold. All rights reserved.
 * Filename: config/config.go
 * Author: harold
 * Mail: mail@yaolong.me
 * Date: 2018-04-17
 */

package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

type config struct {
	jsonMap	map[string]interface{}
	mu		sync.Mutex
}

var (
	conf *config
)

func Config() (*config, error){
	if conf != nil {
		return conf,nil
	}
	s, err := ioutil.ReadFile("/etc/charged/config.json")
	if err != nil {
		log.Printf("%s\n", err)
		return nil, err
	}

	conf = &config{}
	conf.mu.Lock()
	conf.jsonMap = make(map[string]interface{})
	err = json.Unmarshal([]byte(s), &conf.jsonMap)
	conf.mu.Unlock()

	if err != nil {
		log.Printf("%s\n", err)
		return nil, err
	}
	return conf,nil
}

func (c *config) Get(keys ...string) (ret interface{}, err error) {
	defer func(){
		if err := recover(); err != nil {
			// log.Printf("Notice: %s\n", err)
		}
	}()

	if len(keys) == 0 {
		return c.jsonMap, nil
	}

	if len(keys) == 1{
		return c.jsonMap[keys[0]], nil
	}

	tmp := make(map[string]interface{})
	tmp = c.jsonMap[keys[0]].(map[string]interface{})

	for k,v := range keys {
		if k == 0 {
			continue
		}
		if k == len(keys) - 1 {
			ret = tmp[v].(interface{})
		} else {
			tmp = tmp[v].(map[string]interface{})
		}
	}

	return ret, nil
}

