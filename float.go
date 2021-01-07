/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2021-01-07 15:05:35
# File Name: float.go
# Description:
####################################################################### */

package util

import "math"

const MIN = 0.000001

func FloatIsEqual(f1, f2 float64) bool {
	if f1 > f2 {
		return math.Dim(f1, f2) < MIN
	} else {
		return math.Dim(f2, f1) < MIN
	}
}

// vim: set noexpandtab ts=4 sts=4 sw=4 :
