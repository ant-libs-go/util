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
	var cursor string
	defer func() {
		if err != nil {
			return
		}
		if e := recover(); e != nil {
			err = fmt.Errorf("assign %s fail, %s", cursor, e)
		}
	}()

	val_origin := reflect.ValueOf(origin).Elem()
	val_target := reflect.ValueOf(target).Elem()

	for i := 0; i < val_origin.NumField(); i++ {
		cursor = val_origin.Type().Field(i).Name
		if exist, _ := InSlice(val_origin.Type().Field(i).Name, excludes); exist {
			continue
		}

		is_valid := val_target.FieldByName(cursor).IsValid()
		switch val_origin.Field(i).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if is_valid {
				val_target.FieldByName(cursor).SetInt(val_origin.Field(i).Int())
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if is_valid {
				val_target.FieldByName(cursor).SetUint(val_origin.Field(i).Uint())
			}
		case reflect.String:
			if is_valid {
				val_target.FieldByName(cursor).SetString(val_origin.Field(i).String())
			}
		case reflect.Bool:
			if is_valid {
				val_target.FieldByName(cursor).SetBool(val_origin.Field(i).Bool())
			}
		case reflect.Float32, reflect.Float64:
			if is_valid {
				val_target.FieldByName(cursor).SetFloat(val_origin.Field(i).Float())
			}
		case reflect.Map, reflect.Array, reflect.Slice: // ptr type not deep copy
			if is_valid {
				val_target.FieldByName(cursor).Set(reflect.ValueOf(val_origin.Field(i).Interface()))
			}
		case reflect.Struct:
			err = Assign(val_origin.Field(i).Addr().Interface(), val_target.FieldByName(cursor).Addr().Interface(), excludes...)
		case reflect.Ptr:
			err = Assign(val_origin.Field(i).Interface(), val_target.FieldByName(cursor).Interface(), excludes...)
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
func StructIsEqual(origin, target interface{}, excludes ...string) bool {
	val_origin := reflect.ValueOf(origin).Elem()
	val_target := reflect.ValueOf(target).Elem()

	for i := 0; i < val_origin.NumField(); i++ {
		if !val_target.FieldByName(val_origin.Type().Field(i).Name).IsValid() {
			continue
		}
		if exist, _ := InSlice(val_origin.Type().Field(i).Name, excludes); exist {
			continue
		}

		switch val_origin.Field(i).Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if val_target.Field(i).Int() != val_origin.Field(i).Int() {
				return false
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if val_target.Field(i).Uint() != val_origin.Field(i).Uint() {
				return false
			}
		case reflect.String:
			if val_target.Field(i).String() != val_origin.Field(i).String() {
				return false
			}
		case reflect.Bool:
			if val_target.Field(i).Bool() != val_origin.Field(i).Bool() {
				return false
			}
		case reflect.Float32, reflect.Float64:
			if FloatIsEqual(val_target.Field(i).Float(), val_origin.Field(i).Float()) == false {
				return false
			}
		case reflect.Map, reflect.Array, reflect.Slice, reflect.Struct, reflect.Ptr:
			if reflect.DeepEqual(val_target.Field(i).Interface(), val_origin.Field(i).Interface()) == false {
				return false
			}
			/*
				old code, only slice
				if !(reflect.ValueOf(val_origin.Field(i).Interface()).Len() == 0 && reflect.ValueOf(val_origin.Field(i).Interface()).Len() == 0) {
					if reflect.DeepEqual(val_target.Field(i).Interface(), val_origin.Field(i).Interface()) == false {
						return false
					}
				}
			*/
		}
	}
	return true
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
			fmt.Printf("panic recover, %s\n", err)
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

// 获取指定日期周一零点时间
func FirstTimeOfWeek(t time.Time) (r time.Time) {
	offset := int(time.Monday - t.Weekday())
	if offset > 0 {
		offset = -6
	}
	r = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	return
}

// 获取指定日期零点时间
func FirstTimeOfDay(t time.Time) (r time.Time) {
	r = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	return
}
