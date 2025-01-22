/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2020-08-08 13:12:44
# File Name: serialize.go
# Description:
####################################################################### */

package util

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/apache/thrift/lib/go/thrift"
	"google.golang.org/protobuf/proto"
	"github.com/vmihailenco/msgpack"
)

func IntEncode(inp int64, salt int64) (r string) {
	IfDo(salt == 0, func() { salt = 33554432 })
	r = base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(inp+salt, 32)))
	return
}

func IntDecode(inp string, salt int64) (r int64, err error) {
	IfDo(salt == 0, func() { salt = 33554432 })
	bs, err := base64.StdEncoding.DecodeString(inp)
	if err != nil {
		err = fmt.Errorf("%s base64Decode, %s", inp, err)
		return
	}
	r, err = strconv.ParseInt(string(bs), 32, 64)
	if err != nil {
		err = fmt.Errorf("%s parse, %s", inp, err)
		return
	}
	// 32 to 10, - 32768
	r -= salt
	return
}

func JsonEncode(inp interface{}) ([]byte, error) {
	return json.Marshal(inp)
}

func JsonDecode(d []byte, inp interface{}) error {
	err := json.Unmarshal(d, inp)
	return err
}

func GobEncode(inp interface{}) ([]byte, error) {
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(inp)
	return buf.Bytes(), err
}

func GobDecode(d []byte, inp interface{}) error {
	err := gob.NewDecoder(bytes.NewBuffer(d)).Decode(inp)
	return err
}

func PbEncode(inp proto.Message) ([]byte, error) {
	return proto.Marshal(inp)
}

func PbDecode(b []byte, inp proto.Message) error {
	return proto.Unmarshal(b, inp)
}

func MsgpackEncode(inp interface{}) ([]byte, error) {
	var buf bytes.Buffer
	err := msgpack.NewEncoder(&buf).Encode(inp)
	return buf.Bytes(), err
}

func MsgpackDecode(b []byte, inp interface{}) error {
	err := msgpack.NewDecoder(bytes.NewReader(b)).Decode(inp)
	return err
}

func ThriftEncode(inp interface{}) ([]byte, error) {
	b := thrift.NewTMemoryBufferLen(1024)
	p := thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(b)
	t := &thrift.TSerializer{
		Transport: b,
		Protocol:  p,
	}
	t.Transport.Close()
	return t.Write(context.Background(), inp.(thrift.TStruct))
}

func ThriftDecode(b []byte, inp interface{}) error {
	t := thrift.NewTMemoryBufferLen(1024)
	p := thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(t)
	d := &thrift.TDeserializer{
		Transport: t,
		Protocol:  p,
	}
	d.Transport.Close()
	return d.Read(context.Background(), inp.(thrift.TStruct), b)
}

/* serialize value */
type SerializeValue struct {
	Point []string
	Range []*struct {
		S int
		E int
	}
	Dict map[string]string
}

func NewSerializeValue(inp []byte) (r *SerializeValue, err error) {
	err = json.Unmarshal(inp, &r)
	return
}

func (this *SerializeValue) Map() (r map[string]string) {
	if this == nil {
		return
	}
	return this.Dict
}

func (this *SerializeValue) Slice() (r []string) {
	if this == nil {
		return
	}
	if len(this.Point) != 0 {
		r = this.Point
		return
	}
	for _, v := range this.Range {
		if v.S > v.E {
			continue
		}
		for ; v.S <= v.E; v.S++ {
			r = append(r, strconv.Itoa(v.S))
		}
	}
	return
}
