/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-06-13 08:08:55
# File Name: logs.go
# Description:
####################################################################### */

/**
 * import blogs "github.com/astaxie/beego/logs"
 * blogs.SetLogger(blogs.AdapterMultiFile, beego.AppConfig.String("log::format"))
 *
 * go func(){
 *    logs.New(&logs.Entry{Prefix: "[g:__GID__]"})
 *    defer logs.Close()
 *    logs.Error("test error: %s", "error")
 * }()
 */

package logs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ant-libs-go/util"
	"github.com/astaxie/beego/logs"
)

type Entry struct {
	Prefix string
}

var entries map[int]*Entry

func init() {
	entries = make(map[int]*Entry)
}

/**
 * 必须调用Close方法
 */
func New(entry *Entry) {
	gid := util.Goid()
	entry.Prefix = strings.Replace(entry.Prefix, "__GID__", strconv.Itoa(gid), -1)
	entries[gid] = entry
}

func Close() {
	gid := util.Goid()
	delete(entries, gid)
}

func getPrefix() string {
	if entry, ok := entries[util.Goid()]; ok {
		return fmt.Sprintf("%s ", entry.Prefix)
	}
	return ""
}

func Emergency(f string, v ...interface{}) {
	logs.Emergency(fmt.Sprintf("%s%s", getPrefix(), f), v...)
}

func Alert(f string, v ...interface{}) {
	logs.Alert(fmt.Sprintf("%s%s", getPrefix(), f), v...)
}

func Critical(f string, v ...interface{}) {
	logs.Critical(fmt.Sprintf("%s%s", getPrefix(), f), v...)
}

func Error(f string, v ...interface{}) {
	fmt.Println(len(entries))
	logs.Error(fmt.Sprintf("%s%s", getPrefix(), f), v...)
}

func Warn(f string, v ...interface{}) {
	logs.Warn(fmt.Sprintf("%s%s", getPrefix(), f), v...)
}

func Notice(f string, v ...interface{}) {
	logs.Notice(fmt.Sprintf("%s%s", getPrefix(), f), v...)
}

func Info(f string, v ...interface{}) {
	logs.Info(fmt.Sprintf("%s%s", getPrefix(), f), v...)
}

func Debug(f string, v ...interface{}) {
	logs.Debug(fmt.Sprintf("%s%s", getPrefix(), f), v...)
}

func Trace(f string, v ...interface{}) {
	logs.Trace(fmt.Sprintf("%s%s", getPrefix(), f), v...)
}
