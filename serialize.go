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
	"fmt"
	"strconv"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/gogo/protobuf/proto"
)

func IntEncode(inp int64) (r string) {
	r = base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(inp+33554432, 32)))
	return
}

func IntDecode(inp string) (r int64, err error) {
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
	r -= 33554432
	return
}

func GobSerialize(inp interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(inp)
	if err == nil {
		return buf.Bytes(), nil
	}
	return nil, err
}

func GobDeserialize(d []byte, inp interface{}) error {
	dec := gob.NewDecoder(bytes.NewBuffer(d))
	return dec.Decode(inp)
}

func PbSerialize(inp proto.Message) ([]byte, error) {
	return proto.Marshal(inp)
}

func PbDeserialize(b []byte, inp proto.Message) error {
	return proto.Unmarshal(b, inp)
}

func ThriftSerialize(inp interface{}) ([]byte, error) {
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
	return d.Read(inp.(thrift.TStruct), b)
}
