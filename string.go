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
	"fmt"
	"net/url"
	"regexp"
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

func LeftUpper(s string) string {
	if len(s) > 0 {
		return strings.ToUpper(string(s[0])) + s[1:]
	}
	return s
}

func LeftLower(s string) string {
	if len(s) > 0 {
		return strings.ToLower(string(s[0])) + s[1:]
	}
	return s
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

func Tprintf(format string, params map[string]string) string {
	for k, v := range params {
		format = strings.Replace(format, "%("+k+")s", fmt.Sprintf("%s", v), -1)
	}
	return format
}

func BytesReplace(s, old, new []byte, n int) []byte {
	if n == 0 {
		return s
	}

	if len(old) < len(new) {
		return bytes.Replace(s, old, new, n)
	}

	if n < 0 {
		n = len(s)
	}

	var wid, i, j, w int
	for i, j = 0, 0; i < len(s) && j < n; j++ {
		wid = bytes.Index(s[i:], old)
		if wid < 0 {
			break
		}

		w += copy(s[w:], s[i:i+wid])
		w += copy(s[w:], new)
		i += wid + len(old)
	}

	w += copy(s[w:], s[i:])
	return s[0:w]
}
