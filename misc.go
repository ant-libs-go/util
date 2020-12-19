/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-06-06 07:20:06
# File Name: misc.go
# Description:
####################################################################### */

package util

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"math/rand"
	"net"
	"reflect"
	"runtime"
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
func Assign(origin, target interface{}, excludes ...string) (err error) {
	var curFieldName string
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("assign %s: %s", curFieldName, e)
		}
	}()

	val_origin := reflect.ValueOf(origin).Elem()
	val_target := reflect.ValueOf(target).Elem()

	for i := 0; i < val_origin.NumField(); i++ {
		curFieldName = val_origin.Type().Field(i).Name
		is_exclude := false
		for _, col := range excludes {
			if curFieldName == col {
				is_exclude = true
				break
			}
		}
		if is_exclude {
			continue
		}
		is_valid := val_target.FieldByName(curFieldName).IsValid()
		switch val_origin.Field(i).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if !is_valid {
				continue
			}
			val_target.FieldByName(curFieldName).SetInt(val_origin.Field(i).Int())
		case reflect.String:
			if !is_valid {
				continue
			}
			val_target.FieldByName(curFieldName).SetString(val_origin.Field(i).String())
		case reflect.Struct:
			Assign(val_origin.Field(i).Addr().Interface(), target, excludes...)
		case reflect.Ptr:
			Assign(val_origin.Field(i).Interface(), target, excludes...)
		}
	}
	return
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

func DateRange(s, e string) (r []string) {
	st, _ := time.Parse("20060102", s)
	et, _ := strconv.Atoi(e)
	for {
		t, _ := strconv.Atoi(st.Format("20060102"))
		if t > et {
			break
		}
		r = append(r, strconv.Itoa(t))
		st = st.AddDate(0, 0, +1)
	}
	return
}

func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func GenRandomId(salt string) string {
	return Md5String(salt + strconv.FormatInt(time.Now().UnixNano(), 10))
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

/**
 * 获取协程id，该方法性能较差
 */
func Goid() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recover:panic info:%v", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

/**
 * 获取本机IP地址
 */
func GetLocalIP() (r string, err error) {
	var addrs []net.Addr
	if addrs, err = net.InterfaceAddrs(); err != nil {
		return
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				r = ipnet.IP.String()
				break
			}
		}
	}
	return
}

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func IfDo(condition bool, fn func()) {
	if condition {
		fn()
	}
}
