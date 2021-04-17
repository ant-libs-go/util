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
	sessid        string
	prefix        string
	behaviorLevel seelog.LogLevel
	last          int64
	logger        seelog.LoggerInterface
}

func init() {
	entries = make(map[string]*SessLog)
	registerReleaser()
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
	o := &SessLog{sessid: sessid, prefix: "[sid:" + sessid + "]", logger: seelog.Current}
	lock.Lock()
	entries[sessid] = o.use()
	lock.Unlock()
	return o.use()
}

func (this *SessLog) SetBehaviorLevel(bl seelog.LogLevel) {
	this.behaviorLevel = bl
}

func (this *SessLog) Release() {
	lock.Lock()
	this.release()
	lock.Unlock()
}

// unsafe
func (this *SessLog) release() {
	delete(entries, this.sessid)
}

func registerReleaser() {
	go func() {
		for {
			ts := time.Now().Unix()
			lock.Lock()
			for _, entry := range entries {
				if ts-entry.last < 120 { // timeout for 2 minute
					continue
				}
				entry.release()
			}
			lock.Unlock()
			time.Sleep(10 * time.Second) // interval 10 second
		}
	}()
}

func (this *SessLog) use() *SessLog {
	this.last = time.Now().Unix()
	return this
}

func (this *SessLog) Tracef(f string, v ...interface{}) {
	if this.behaviorLevel > seelog.TraceLvl {
		return
	}
	this.use().logger.Tracef(this.prefix+" "+f, v...)
}

func (this *SessLog) Debugf(f string, v ...interface{}) {
	if this.behaviorLevel > seelog.DebugLvl {
		return
	}
	this.use().logger.Debugf(this.prefix+" "+f, v...)
}

func (this *SessLog) Infof(f string, v ...interface{}) {
	if this.behaviorLevel > seelog.InfoLvl {
		return
	}
	this.use().logger.Infof(this.prefix+" "+f, v...)
}

func (this *SessLog) Warnf(f string, v ...interface{}) {
	if this.behaviorLevel > seelog.WarnLvl {
		return
	}
	this.use().logger.Warnf(this.prefix+" "+f, v...)
}

func (this *SessLog) Errorf(f string, v ...interface{}) {
	if this.behaviorLevel > seelog.ErrorLvl {
		return
	}
	this.use().logger.Errorf(this.prefix+" "+f, v...)
}

func (this *SessLog) Criticalf(f string, v ...interface{}) {
	if this.behaviorLevel > seelog.CriticalLvl {
		return
	}
	this.use().logger.Criticalf(this.prefix+" "+f, v...)
}
