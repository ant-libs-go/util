/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-11-30 11:25:26
# File Name: string.go
# Description:
####################################################################### */

package util

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// camel string, xx_yy to XxYy
func CamelString(inp string) string {
	data := make([]byte, 0, len(inp))
	j := false
	k := false
	num := len(inp) - 1
	for i := 0; i <= num; i++ {
		d := inp[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && inp[i+1] >= 'a' && inp[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// snake string, XxYy to xx_yy, XxYY to xx_yy
func SnakeString(inp string) string {
	data := make([]byte, 0, len(inp)*2)
	j := false
	num := len(inp)
	for i := 0; i < num; i++ {
		d := inp[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

// regexp.Compile(`\[(?P<node>[\d_]+)\]$`)
// return {"node":val}
func FindStringSubmatch(re *regexp.Regexp, inp string) (r map[string]string, err error) {
	r = make(map[string]string)
	match := re.FindStringSubmatch(inp)
	if match == nil {
		return nil, errors.New("no match")
	}
	for i, name := range re.SubexpNames() {
		if i == 0 || name == "" {
			continue
		}
		r[name] = match[i]
	}
	return
}

func Md5String(inp string) string {
	h := md5.New()
	h.Write([]byte(inp))
	return hex.EncodeToString(h.Sum(nil))
}

func Sha1String(inp string) string {
	h := sha1.New()
	h.Write([]byte(inp))
	return hex.EncodeToString(h.Sum(nil))
}

func GzipString(inp string) (r []byte, err error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	defer w.Close()
	if _, err = w.Write([]byte(inp)); err == nil {
		err = w.Flush()
	}
	if err != nil {
		return
	}
	r = buf.Bytes()
	return
}

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

func UrlEncode(inp string) string {
	return url.QueryEscape(inp)
}

func UrlDecode(inp string) string {
	r, err := url.QueryUnescape(inp)
	if err != nil {
		return ""
	}
	return r
}

func Max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func MinInt(x, y int) int {
	if x < y {
		return x
	}

	return y
}
