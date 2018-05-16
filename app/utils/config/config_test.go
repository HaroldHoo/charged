/**
 * Copyright 2018 harold. All rights reserved.
 * Filename: app/utils/config/config_test.go
 * Author: harold
 * Mail: mail@yaolong.me
 * Date: 2018-04-23
 */

package config
import (
	"testing"
	"fmt"
)

func TestConfig(t *testing.T) {
	r, e := Get("mysql", "yc_crm_common", "master", "port")
	if e != nil {
		t.Fatalf("%#v\n", e)
	}
	fmt.Printf("%#v\n", r)

	fmt.Printf("\n-------------\n")
	r, e = Get("hosts", "lbs_host")
	if e != nil {
		t.Fatalf("%#v\n", e)
	}
	fmt.Printf("%#v\n", r)
}

