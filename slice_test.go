/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2021-01-07 14:21:10
# File Name: slice_test.go
# Description:
####################################################################### */

package util

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

type s struct {
	A [2]string

	B string
	C bool
	D float64
	G int32
	H uint32

	E map[string]string
	F *d
}

type d struct {
	AA string
	BB int32
}

func TestBasic(t *testing.T) {
	s1 := &s{A: [2]string{"aa", "bb"}, B: "b1", C: true, D: 1.5, E: map[string]string{"aa": "bb"}, F: &d{"dd", 12}, G: 10, H: 20}
	s2 := &s{F: &d{}}
	fmt.Println(Assign(s1, s2))
	fmt.Printf("%+v\n", s1)
	fmt.Printf("%+v\n", s1.F)
	fmt.Printf("%+v\n", s2)
	fmt.Printf("%+v\n", s2.F)
	s1.F.AA = "dddd"
	fmt.Printf("%+v\n", s1)
	fmt.Printf("%+v\n", s1.F)
	fmt.Printf("%+v\n", s2)
	fmt.Printf("%+v\n", s2.F)
}

// vim: set noexpandtab ts=4 sts=4 sw=4 :
