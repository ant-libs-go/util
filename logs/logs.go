/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2018-06-13 08:08:55
# File Name: logs.go
# Description:
####################################################################### */

/**
 * import (
 * 	"os"
 * 	"github.com/cihub/seelog"
 * )
 * logger, err := seelog.LoggerFromConfigAsFile("log.xml")
 * if  err != nil {
 *     os.Exit(-1)
 * }
 * seelog.ReplaceLogger(logger)
 * defer seelog.Flush()
 *
 * log := logs.New(uuid)
 * log.Infof("this is a %s", "log")
 */

package logs

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/ant-libs-go/util"
	"github.com/cihub/seelog"
)

var (
	lock    sync.RWMutex
	entries map[string]*SessLog
)

type SessLog struct {
	sessid string
	last   int64
	logger seelog.LoggerInterface
}

func init() {
	entries = make(map[string]*SessLog)
	registerCleaner()
}

func New(sessid string) (r *SessLog) {
	if len(sessid) == 0 {
		sessid = strconv.Itoa(util.Goid())
	}
	lock.RLock()
	r, ok := entries[sessid]
	lock.RUnlock()
	if !ok {
		r = build(sessid)
	}
	return
}

func build(sessid string) *SessLog {
	o := &SessLog{sessid: sessid, logger: seelog.Current}
	lock.Lock()
	defer lock.Unlock()
	entries[sessid] = o
	return o.use()
}

func (this *SessLog) Close() {
	lock.Lock()
	delete(entries, this.sessid)
	lock.Unlock()
}

func registerCleaner() {
	go func() {
		for {
			fmt.Println("len:", len(entries))
			ts := time.Now().Unix()
			for _, entry := range entries {
				if ts-entry.last < 120 { // timeout for 2 minute
					continue
				}
				entry.Close()
			}
			time.Sleep(10 * time.Second) // interval 10 second
		}
	}()
}

func (this *SessLog) use() *SessLog {
	this.last = time.Now().Unix()
	return this
}

func (this *SessLog) Tracef(f string, v ...interface{}) {
	this.use().logger.Tracef(fmt.Sprintf("[sid:%s] %s", this.sessid, f), v...)
}

func (this *SessLog) Debugf(f string, v ...interface{}) {
	this.use().logger.Debugf(fmt.Sprintf("[sid:%s] %s", this.sessid, f), v...)
}

func (this *SessLog) Infof(f string, v ...interface{}) {
	this.use().logger.Infof(fmt.Sprintf("[sid:%s] %s", this.sessid, f), v...)
}

func (this *SessLog) Warnf(f string, v ...interface{}) {
	this.use().logger.Warnf(fmt.Sprintf("[sid:%s] %s", this.sessid, f), v...)
}

func (this *SessLog) Errorf(f string, v ...interface{}) {
	this.use().logger.Errorf(fmt.Sprintf("[sid:%s] %s", this.sessid, f), v...)
}

func (this *SessLog) Criticalf(f string, v ...interface{}) {
	this.use().logger.Criticalf(fmt.Sprintf("[sid:%s] %s", this.sessid, f), v...)
}
