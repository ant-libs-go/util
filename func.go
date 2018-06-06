/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-06-06 07:20:06
# File Name: func.go
# Description:
####################################################################### */

package util

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"errors"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/**
 * assign one struct to other struct
 * @param: origin
 * @param: target
 * @params: excludes ... the attribute name exclude assign
 */
func Assign(origin, target interface{}, excludes ...string) {
	val_origin := reflect.ValueOf(origin).Elem()
	val_target := reflect.ValueOf(target).Elem()

	for i := 0; i < val_origin.NumField(); i++ {
		if !val_target.FieldByName(val_origin.Type().Field(i).Name).IsValid() {
			continue
		}
		is_exclude := false
		for _, col := range excludes {
			if val_origin.Type().Field(i).Name == col {
				is_exclude = true
				break
			}
		}
		if is_exclude {
			continue
		}
		switch val_origin.Field(i).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val_target.FieldByName(val_origin.Type().Field(i).Name).SetInt(val_origin.Field(i).Int())
		case reflect.String:
			val_target.FieldByName(val_origin.Type().Field(i).Name).SetString(val_origin.Field(i).String())
		}
	}
}

/**
 * assign one struct to other struct
 * @param: origin
 * @param: target
 * @params: excludes ... the attribute name exclude check
 */
func IsChanged(origin, target interface{}, excludes ...string) bool {
	val_origin := reflect.ValueOf(origin).Elem()
	val_target := reflect.ValueOf(target).Elem()

	for i := 0; i < val_origin.NumField(); i++ {
		if !val_target.FieldByName(val_origin.Type().Field(i).Name).IsValid() {
			continue
		}
		is_exclude := false
		for _, col := range excludes {
			if val_origin.Type().Field(i).Name == col {
				is_exclude = true
				break
			}
		}
		if is_exclude {
			continue
		}
		switch val_origin.Field(i).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if val_target.Field(i).Int() != val_origin.Field(i).Int() {
				return true
			}
		case reflect.String:
			if val_target.Field(i).String() != val_origin.Field(i).String() {
				return true
			}
		case reflect.Slice:
			if !(reflect.ValueOf(val_origin.Field(i).Interface()).Len() == 0 && reflect.ValueOf(val_origin.Field(i).Interface()).Len() == 0) {
				if !reflect.DeepEqual(val_target.Field(i).Interface(), val_origin.Field(i).Interface()) {
					return true
				}
			}
		case reflect.Map:
			if !reflect.DeepEqual(val_target.Field(i).Interface(), val_origin.Field(i).Interface()) {
				return true
			}

		}
	}
	return false
}

// camel string, xx_yy to XxYy
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// snake string, XxYy to xx_yy, XxYY to xx_yy
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
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

func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

// regexp.Compile(`\[(?P<node>[\d_]+)\]$`)
// return {"node":val}
func FindStringSubmatch(re *regexp.Regexp, s string) (r map[string]string, err error) {
	r = make(map[string]string)
	match := re.FindStringSubmatch(s)
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

func GenGoroutineId(salt string) string {
	salt = salt + strconv.FormatInt(time.Now().UnixNano(), 10)

	h := md5.New()
	h.Write([]byte(salt))
	return hex.EncodeToString(h.Sum(nil))
}

func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	r := []byte{}
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		r = append(r, bytes[rd.Intn(len(bytes))])
	}
	return string(r)
}
