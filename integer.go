/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2020-12-08 17:43:01
# File Name: int.go
# Description:
####################################################################### */

package util

func MaxInt64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func MinInt64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func MaxInt32(x, y int32) int32 {
	if x > y {
		return x
	}
	return y
}

func MinInt32(x, y int32) int32 {
	if x < y {
		return x
	}
	return y
}

// vim: set noexpandtab ts=4 sts=4 sw=4 :
