/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2022-05-16 10:37:40
# File Name: map_test.go
# Description:
####################################################################### */

package util

import (
	"fmt"
	"testing"
)

type da struct {
	AA string
	BB string
}

func TestMapColumn(t *testing.T) {
	ms := map[string]*da{
		"a1": &da{AA: "AA", BB: "BB"},
		"a2": &da{AA: "AA", BB: "BB"},
		"a3": &da{AA: "AA", BB: "BB"},
		"a4": &da{AA: "AA", BB: "BB"},
	}
	fmt.Println(MapColumn(ms, "BB"))
}

// vim: set noexpandtab ts=4 sts=4 sw=4 :
