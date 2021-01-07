/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2021-01-07 13:09:42
# File Name: type_convert.go
# Description:
####################################################################### */

package util

import "strconv"

func StrToInt64(inp string, defaultValue int64) int64 {
	val, err := strconv.ParseInt(inp, 10, 64)
	if err != nil {
		return defaultValue
	}
	return val
}

func Int64ToStr(inp int64) string {
	return strconv.FormatInt(inp, 10)
}

func StrToInt32(inp string, defaultValue int32) int32 {
	return int32(StrToInt64(inp, int64(defaultValue)))
}

func Int32ToStr(inp int32) string {
	return Int64ToStr(int64(inp))
}

func StrToInt(inp string, defaultValue int) int {
	return int(StrToInt64(inp, int64(defaultValue)))
}

func IntToStr(inp int) string {
	return Int64ToStr(int64(inp))
}

func StrToFloat64(inp string, defaultValue float64) float64 {
	val, err := strconv.ParseFloat(inp, 64)
	if err != nil {
		return defaultValue
	}
	return val
}

func Float64ToStr(inp float64) string {
	return strconv.FormatFloat(inp, 'E', -1, 64)
}