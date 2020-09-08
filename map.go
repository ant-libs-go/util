/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-06-06 09:23:03
# File Name: map.go
# Description:
####################################################################### */

package util

import (
	"fmt"
	"strings"
)

func JoinMap(arr map[string]string, glue string, glue2 string) (r string) {
	var t []string
	for n, v := range arr {
		t = append(t, fmt.Sprintf("%s%s%s", n, glue, v))
	}
	r = strings.Join(t, glue2)
	return
}

func MapToQueryStr(inp map[string]string) (r string) {
	for k, v := range inp {
		inp[k] = UrlEncode(v)
	}
	r = JoinMap(inp, "=", "&")
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
