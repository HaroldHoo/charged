/**
 * Copyright 2018 harold. All rights reserved.
 * Filename: dao_test.go
 * Author: harold
 * Mail: mail@yaolong.me
 * Date: 2018-04-11
 */

package dao

import(
	"testing"
	"fmt"
)

func TestGetSlicesFromQuery(t *testing.T){
	fmt.Println("\nTestGetSlicesFromQuery ---------------------- ")
	rows, err := yc_crm_common.getSlicesFromQuery(`
	SELECT
	rent_rate_id,
	city,
	car_type_id,
	granularity
	from rent_rate limit 2;
	`)

	if err != nil {
		t.Fatal(err)
	}

	for k,v := range *rows {
		fmt.Printf("%#v\n", k)
		for k2,v2 := range v {
			fmt.Printf("%#v, %#v\n", k2,v2)
		}
		fmt.Printf("\n")
	}
}

func TestGetPkeyMapsFromQuery(t *testing.T){
	fmt.Println("\nTestGetPkeyMapsFromQuery ---------------------- ")
	rows, err := yc_crm_common.getPkeyMapsFromQuery(`
	SELECT
	rent_rate_id,
	city,
	car_type_id,
	granularity
	from rent_rate limit 2;
	`, "rent_rate_id")

	if err != nil {
		t.Fatal(err)
	}

	for k,v := range *rows {
		fmt.Printf("%#v\n", k)
		for k2,v2 := range v {
			fmt.Printf("%#v, %#v\n", k2,v2)
		}
		fmt.Printf("\n")
	}
}

func TestGetPkeysByColName(t *testing.T){
	fmt.Println("\nTestGetPkeysByColName ---------------------- ")
	rows, err := yc_crm_common.getPkeyMapsFromQuery(`
	SELECT
	rent_rate_id,
	city,
	car_type_id,
	granularity
	from rent_rate
	`, "rent_rate_id")

	rows2, _ := yc_crm_common.getPkeyMapsFromQuery(`
	SELECT
	rent_rate_id,
	city,
	car_type_id,
	granularity
	from rent_rate where city="bj"
	`, "rent_rate_id")

	if err != nil {
		t.Fatal(err)
	}

	pks,_ := getPkeysByColName(rows, "city", "bj")
	if len(*rows2) != len(*pks) {
		t.Fatalf("len not eq")
	}
	for _,v := range *pks{
		if (*rows2)[v] == nil {
			t.Fatalf("faild")
		}
	}

	fmt.Printf("count:%d\n\n", len(*pks))
}

func TestNewAndPkeys(t *testing.T){
	fmt.Println("\nTestNewAndPkeys ---------------------- ")
	d,_ := GetRentRate()
	ds,_ := d.GetPkeysByColName("city", "hrb")
	fmt.Printf("%#v\n", ds)
	ds2,_ := d.GetPkeysByColName("city", "su")
	fmt.Printf("%#v\n", ds2)
}

func TestGetDatasMap(t *testing.T){
	fmt.Println("\nTestGetDatasMap ---------------------- ")
	d,_ := GetRentRate()
	_,err := d.GetDatasMap("product_type_id","17","city","sh")
	if err != nil {
		t.Fatalf("%#v\n", err)
	}
}

func TestGetDatasSlice(t *testing.T){
	fmt.Println("\nTestGetDatasSlice ---------------------- ")
	d,_ := GetRentRate()
	_,err := d.GetDatasSlice("product_type_id","17","city","sh")
	if err != nil {
		t.Fatalf("%#v\n", err)
	}
}

func TestGetDatasMapCorrectness(t *testing.T){
	fmt.Println("\nTestGetDatasMapCorrectness ---------------------- ")
	d,_ := GetRentRate()

	citys := []string{"gz","hrb","sh","tj","bj","sh","su","hz","tj"}

	var data1,data2,data3 map[int64]map[string]string

	for _,v := range citys {
		data1,_ = d.GetDatasMap("status","1","product_type_id","1","city",v)
		data2,_ = d.GetDatasMap("status","1","product_type_id","1","city",v)
		data3,_ = d.GetDatasMap("status","1","product_type_id","1","city",v)
		for id,v := range data1 {
			for col,val := range v{
				v2 := data2[id]
				v3 := data3[id]
				if v2[col] != val {
					t.Fatalf("TestGetDatasMapCacheHit Faild.\n")
				}
				if v3[col] != val {
					t.Fatalf("TestGetDatasMapCacheHit Faild.\n")
				}
			}
		}
	}
}

func TestGetDatasSliceCorrectness(t *testing.T){
	fmt.Println("\nTestGetDatasSliceCorrectness ---------------------- ")
	d,_ := GetRentRate()

	citys := []string{"gz","hrb","sh","tj","bj","sh","su","hz","tj"}

	var data1,data2,data3 []map[string]string

	for _,v := range citys {
		data1,_ = d.GetDatasSlice("status","1","product_type_id","1","city",v)
		data2,_ = d.GetDatasSlice("status","1","product_type_id","1","city",v)
		data3,_ = d.GetDatasSlice("status","1","product_type_id","1","city",v)
		for id,v := range data1 {
			for col,val := range v{
				v2 := data2[id]
				v3 := data3[id]
				if v2[col] != val {
					t.Fatalf("TestGetDatasMapCacheHit Faild.\n")
				}
				if v3[col] != val {
					t.Fatalf("TestGetDatasMapCacheHit Faild.\n")
				}
			}
		}
	}
}

func BenchmarkGetPkeysByColName(b *testing.B) {
	b.StopTimer()
	d,_ := GetRentRate()

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d.GetPkeysByColName("city", "sh")
	}
}

func BenchmarkTestGetDatasMap(b *testing.B) {
	b.StopTimer()
	d,_ := GetRentRate()

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d.GetDatasMap("product_type_id","17","city","sh")
	}
}

func BenchmarkTestGetDatasSlice(b *testing.B) {
	b.StopTimer()
	d,_ := GetRentRate()

	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		d.GetDatasSlice("product_type_id","17","city","sh")
	}
}

