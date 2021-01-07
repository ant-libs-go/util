/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-11-30 11:27:02
# File Name: version.go
# Description:
####################################################################### */

package util

import (
	"fmt"
	"strconv"
	"strings"
)

func VersionStringToInt(inp string) int32 {
	var s []string
	for _, v := range strings.Split(inp, ".") {
		if len(v) == 1 {
			v = fmt.Sprintf("0%s", v)
		}
		s = append(s, v)
	}
	r, _ := strconv.ParseInt(strings.Join(s, ""), 10, 64)
	return int32(r)
}

func VersionIntToString(inp int32) string {
	var s []string
	for _, v := range []int32{10000, 100, 1} {
		broken := inp / v
		inp -= broken * v
		s = append(s, strconv.Itoa(int(broken)))
	}
	return strings.Join(s, ".")
}

func CompareVersion(v1, v2 string) int {
	v1Nums := strings.Split(v1, ".")
	v2Nums := strings.Split(v2, ".")

	versionNumLen := len(v1Nums)
	if len(v2Nums) > len(v1Nums) {
		versionNumLen = len(v2Nums)
	}

	for len(v1Nums) < versionNumLen {
		v1Nums = append(v1Nums, "0")
	}

	for len(v2Nums) < versionNumLen {
		v2Nums = append(v2Nums, "0")
	}

	for i := 0; i < versionNumLen; i++ {
		vINum := StrToInt(v1Nums[i], 0)
		vJNum := StrToInt(v2Nums[i], 0)

		if vINum > vJNum {
			return 1

		} else if vINum < vJNum {
			return -1
		}
	}

	return 0
}
