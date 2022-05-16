/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-06-06 09:23:03
# File Name: map.go
# Description:
####################################################################### */

package util

import (
	"reflect"
	"strings"
)

func JoinMap(arr map[string]string, glue string, glue2 string) (r string) {
	t := make([]string, 0, len(arr))
	for n, v := range arr {
		t = append(t, n+glue+v)
	}
	r = strings.Join(t, glue2)
	return
}

func MapToQueryStr(inp map[string]string) (r string) {
	t := make(map[string]string, len(inp))
	for k, v := range inp {
		t[k] = UrlEncode(v)
	}
	r = JoinMap(t, "=", "&")
	return
}

func QueryStrToMap(inp string) (r map[string]string) {
	r = map[string]string{}
	for _, v := range strings.Split(inp, "&") {
		p := strings.Split(v, "=")
		if len(p) != 2 {
			continue
		}
		r[p[0]] = UrlDecode(p[1])
	}
	return
}

func MapKeys(inp map[string]interface{}) (r []string) {
	for k, _ := range inp {
		r = append(r, k)
	}
	return
}

func MapColumn(ms interface{}, col string) (r interface{}) {
	if reflect.TypeOf(ms).Kind() != reflect.Map {
		return
	}
	s := reflect.ValueOf(ms)
	for _, i := range s.MapKeys() {
		f := s.MapIndex(i).Elem().FieldByName(col)
		if f.IsValid() == false {
			continue
		}
		switch f.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
			if r == nil {
				r = []int32{}
			}
			r = append(r.([]int32), int32(f.Int()))
		case reflect.Int64:
			if r == nil {
				r = []int64{}
			}
			r = append(r.([]int64), f.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
			if r == nil {
				r = []uint32{}
			}
			r = append(r.([]uint32), uint32(f.Uint()))
		case reflect.Uint64:
			if r == nil {
				r = []uint64{}
			}
			r = append(r.([]uint64), f.Uint())
		case reflect.String:
			if r == nil {
				r = []string{}
			}
			r = append(r.([]string), f.String())
		case reflect.Bool:
			if r == nil {
				r = []bool{}
			}
			r = append(r.([]bool), f.Bool())
		case reflect.Float32:
			if r == nil {
				r = []float32{}
			}
			r = append(r.([]float32), float32(f.Float()))
		case reflect.Float64:
			if r == nil {
				r = []float64{}
			}
			r = append(r.([]float64), f.Float())
		}
	}
	return
}
