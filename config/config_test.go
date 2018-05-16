/**
 * Copyright 2018 harold. All rights reserved.
 * Filename: config_test.go
 * Author: harold
 * Mail: mail@yaolong.me
 * Date: 2018-04-17
 */

package config

import(
	"testing"
	"fmt"
)

func TestConfig(t *testing.T) {
	fmt.Println("\nTestConfig ---------------------- ")
	c,e := Config()
	if e != nil {
		t.Fatalf("%s\n", e)
	}

	r,err := c.Get()
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	fmt.Printf("%#v\n", r)
}

func TestConfigByKeys(t *testing.T) {
	fmt.Println("\nTestConfigByKeys ---------------------- ")
	c,e := Config()
	if e != nil {
		t.Fatalf("%s\n", e)
	}

	r, err := c.Get("mysql")
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	fmt.Printf("%#v\n", r)

	r, err = c.Get("mysql", "yc_crm_common", "slave", "port")
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	fmt.Printf("---------%#v\n", r)

	r, err = c.Get("xxxxxxxxxxxx")
	fmt.Printf("---------%#v\n", r)

	if err != nil {
		t.Fatalf("%s\n", err)
	}
	r, err = c.Get("mysql", "fffffffff")
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	fmt.Printf("---------%#v\n", r)
}

