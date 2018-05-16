/**
 * Copyright 2018 harold. All rights reserved.
 * Filename: app/utils/config/config.go
 * Author: harold
 * Mail: mail@yaolong.me
 * Date: 2018-04-23
 */

package config

import (
	"charged/config"
	"charged/app/utils/ycconfig"
	"errors"
)

var (
	ErrCfgTypeErr = errors.New("cfg type error.")
)

func Get(keys ...string) (ret interface{}, err error) {
	conf, e := config.Config()
	if e != nil {
		return nil, e
	}

	ret, err = conf.Get(keys...)
	switch ret.(type){
	case map[string]interface {} :
		if (ret.(map[string]interface {}))["cfg"] == nil {
			if err != nil {
				return
			}
		}
		break
	default:
		if err != nil {
			return
		}
	}

	// from cfg
	var bufkey []string
	for k,_ := range keys {
		bufkey = make([]string, 0, len(keys) + 1)
		for i:=0;i<k+1;i++ {
			bufkey = append(bufkey, keys[i])
		}
		bufkey = append(bufkey, "cfg")
		val, _ := conf.Get(bufkey...)

		if val != nil {
			// fmt.Printf("\n%#v\n", val)
			switch val.(type){
			case []interface {} :
				switch val.([]interface {})[0] {
				case "map":
					maps := make(map[string]string)
					maps, err = ycconfig.GetMapConfig(val.([]interface {})[1].(string))
					if err != nil {
						return nil, err
					}
					if k < len(keys)-1 {
						return maps[keys[k+1]], nil
					} else {
						return maps, nil
					}
					break
				case "list":
					ret, err = ycconfig.GetListConfig(val.([]interface {})[1].(string))
					return
				case "string":
					ret, err = ycconfig.GetStringConfig(val.([]interface {})[1].(string))
					return
				default:
					return nil, ErrCfgTypeErr
				}
				break
			default:
				return nil, ErrCfgTypeErr
			}
		}
	}

	return
}

