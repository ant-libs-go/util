/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-06-06 09:23:03
# File Name: map.go
# Description:
####################################################################### */

package util

import (
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
