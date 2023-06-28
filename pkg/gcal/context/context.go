// Package contextx 提供每次 RAL 请求的上下文对象，主要用来输出日志。
package context

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/layasugar/laya/core/constants"
	"github.com/layasugar/laya/core/metautils"
	"github.com/layasugar/laya/core/util"
	"go.uber.org/zap"
)

// Request Web请求的上下文
type Request interface {
	LogId() string
	Inject(c context.Context, md metautils.NiceMD)
	Info(template string, args ...interface{})
	Warn(template string, args ...interface{})
	Error(template string, args ...interface{})
}

// Context 用作日志记录
type Context struct {
	ReqContext Request

	Caller      string
	ServiceName string
	ReqLen      int64
	RspLen      int64
	Method      string
	LogId       interface{}
	Protocol    string
	BalanceName string

	PackStatis *StatisItem

	MaxTry int

	curTryIndex   int
	invokeRecords []*InvokeRecord
	lock          *sync.RWMutex
}

// NewContext 创建一个context
func NewContext() (ctx *Context) {
	return &Context{
		PackStatis: &StatisItem{},
		LogId:      util.GenerateLogId(),
		lock:       new(sync.RWMutex),
	}
}

// CurRecord 当前的访问记录
func (ctx *Context) CurRecord() *InvokeRecord {
	for len(ctx.invokeRecords) < ctx.curTryIndex+1 {
		ctx.invokeRecords = append(ctx.invokeRecords, &InvokeRecord{
			timeStatis: map[string]*StatisItem{},
			index:      ctx.curTryIndex,
			timePoints: map[string]time.Time{},
			lock:       new(sync.RWMutex),
		})
	}

	return ctx.invokeRecords[ctx.curTryIndex]
}

func (ctx *Context) Log() {
	for _, invokeRecord := range ctx.invokeRecords {
		var datetime string
		var rspBody string
		if reqStartTime, ok := invokeRecord.timePoints["req_start_time"]; ok {
			datetime = reqStartTime.Format(constants.TIMEFORMAT)
		}
		begin := invokeRecord.timeStatis["cost"].StartPoint
		end := invokeRecord.timeStatis["cost"].StopPoint
		runtime := invokeRecord.timeStatis["cost"].GetSpan()
		if rspLog, ok := invokeRecord.RspLog.([]byte); ok {
			rspBody = string(rspLog)
		}
		ctx.ReqContext.Info("%s", "sdk_log",
			zap.String("datetime", datetime),
			zap.Any("message_type", "sdk"),
			zap.Any("request", invokeRecord.ReqLog),
			zap.Any("respon", rspBody),
			zap.Any("start_time", float64(begin.UnixMicro())/1e6),
			zap.Any("end_time", float64(end.UnixMicro())/1e6),
			zap.String("runtime", runtime),
		)
	}
}

// NextRecord 将访问记录往后移一位
func (ctx *Context) NextRecord() {
	ctx.curTryIndex++
}

// StatisItem 时间统计项
type StatisItem struct {
	StartPoint time.Time
	StopPoint  time.Time
}

// GetSpan 得到耗时
func (si *StatisItem) GetSpan() string {
	if si == nil || si.StartPoint.IsZero() || si.StopPoint.IsZero() {
		return "0"
	}

	span := si.StopPoint.Sub(si.StartPoint)
	return fmt.Sprintf("%.3f", float64(span/time.Nanosecond)/1000000)
}

// TimeStatisStart 开始一个统计项
func (ctx *Context) TimeStatisStart(topic string) {
	ctx.lock.RLock()
	if ctx.CurRecord().timeStatis[topic] != nil { // 被设置过了
		ctx.lock.RUnlock()
		return
	}
	ctx.lock.RUnlock()
	ctx.lock.Lock()
	defer ctx.lock.Unlock()
	if _, ok := ctx.CurRecord().timeStatis[topic]; !ok {
		ctx.CurRecord().timeStatis[topic] = &StatisItem{
			StartPoint: time.Now(),
		}
	}

}

// TimeStatisStop 停止一个统计项
func (ctx *Context) TimeStatisStop(topic string) {
	ctx.lock.RLock()
	defer ctx.lock.RUnlock()
	tmp := ctx.CurRecord().timeStatis[topic]
	if tmp == nil {
		return
	}
	tmp.StopPoint = time.Now()
}

// InvokeRecord 访问日志，因为重试可能有多条
type InvokeRecord struct {
	LogId  string
	RspLog interface{}
	ReqLog constants.RequestLog

	// RspCode 请求的响应码
	// http 代表 http status code，200 为正常，700+是自定义的错误码，表示发送请求时发生了error
	// nshead 等有自己的规则，不统一描述
	RspCode int

	// Path 请求的路径
	// http 相对path， 形如： /foo/bar
	Path string

	// IPPort ip和端口号
	IPPort string

	// Host 域名，可能和IPPort 一致
	Host string

	// 一次请求最多一条错误日志
	Error error

	timeStatis map[string]*StatisItem
	timePoints map[string]time.Time
	index      int
	lock       *sync.RWMutex
}

// GetTimeStatis 获取一个统计项
func (invokeRecord *InvokeRecord) GetTimeStatis(topic string) string {
	invokeRecord.lock.RLock()
	defer invokeRecord.lock.RUnlock()
	tmp := invokeRecord.timeStatis[topic]
	if tmp == nil {
		return "0"
	}
	return tmp.GetSpan()
}

// RecordTimePoint 打下一个时间点
func (invokeRecord *InvokeRecord) RecordTimePoint(topic string) {
	if _, ok := invokeRecord.timePoints[topic]; ok {
		return
	}
	invokeRecord.timePoints[topic] = time.Now()
}

// GetTimePoint 得到一个时间点 毫秒
func (invokeRecord *InvokeRecord) GetTimePoint(topic string) string {
	t := invokeRecord.timePoints[topic]
	if t.IsZero() {
		return "0"
	}

	return strconv.FormatInt(t.UnixNano()/1000000, 10)
}
