package ycconfig

import (
	"testing"
	"fmt"
)

func TestGetCurrentConfigEnv(t *testing.T) {
	fmt.Println("\nTestGetCurrentConfigEnv---------------------- ")
	if GetCurrentConfigEnv() != "test" {
		t.Error("Test ycconfig:GetCurrentConfigEnv() failed")
	} else {
		t.Log("Test ycconfig:GetCurrentConfigEnv() pass")
	}
}

func TestGetStringConfig(t *testing.T) {
	fmt.Println("\nTestGetStringConfig---------------------- ")
	res, err := GetStringConfig("dispatch.merchant.api_host")
	if err != nil || res != "http://newtest.merchant.yongche.org" {
		t.Error("Test ycconfig:GetStringConfig() failed")
	} else {
		t.Log("Test ycconfig:GetStringConfig() pass")
	}
}

func TestGetListConfig(t *testing.T) {
	fmt.Println("\nTestGetListConfig---------------------- ")
	res, err := GetListConfig("order.redis.servers")
	if err != nil || res[0] != "10.0.11.150" || res[1] != "10.0.11.151" {
		t.Error("Test ycconfig:GetListConfig() failed")
	} else {
		t.Log("Test ycconfig:GetListConfig() pass")
	}
}

func TestGetMapConfig(t *testing.T) {
	fmt.Println("\nTestGetMapConfig---------------------- ")
	res, err := GetMapConfig("risk-control.face_recog_global_config")
	if err != nil || res["confidence_threshold"] != "1e-4" || res["retry_limit_per_day"] != "5" {
		t.Error("Test ycconfig:GetMapConfig() failed")
	} else {
		t.Log("Test ycconfig:GetMapConfig() pass")
	}
}
