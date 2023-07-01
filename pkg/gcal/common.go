package gcal

import (
	"context"
	"github.com/layasugar/laya-template/pkg/core"
	"strings"

	"github.com/layasugar/laya-template/pkg/core/constants"
	"github.com/layasugar/laya-template/pkg/core/metautils"
	"github.com/layasugar/laya-template/pkg/gcal/converter"
	"github.com/layasugar/laya-template/pkg/gcal/pool"
	"github.com/layasugar/laya-template/pkg/gcal/protocol"
	"github.com/layasugar/laya-template/pkg/gcal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var pbTc = &pool.Pool{}

// HTTPRequest 别名
type HTTPRequest = protocol.HTTPRequest

// HTTPHead 别名
type HTTPHead = protocol.HTTPHead

// ConverterType 别名
type ConverterType = converter.ConverterType

// JSONConverter 别名
var JSONConverter = converter.JSON

// FORMConverter 别名
var FORMConverter = converter.FORM

// RAWConverter 别名
var RAWConverter = converter.RAW

// LoadService load one service from struct
func LoadService(configs []map[string]interface{}) error {
	return service.LoadService(configs)
}

func GetRpcConn(serverName string) *grpc.ClientConn {
	srv, ok := service.GetService(serverName)
	if !ok {
		return nil
	}

	curConnKey := pool.Key{
		Schema: "tcp",
		Addr:   srv.GetAddr(),
	}
	tcConn, _ := pbTc.Get(curConnKey)
	if tcConn == nil {
		conn, errDial := grpc.Dial(srv.GetAddr(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithUnaryInterceptor(clientInterceptor),
		)
		if errDial != nil {
			return nil
		}
		//c := pool.Func{
		//	Factory: func() (interface{}, error) {
		//		return grpc.Dial(srv.GetAddr(),
		//			grpc.WithTransportCredentials(insecure.NewCredentials()),
		//			grpc.WithUnaryInterceptor(clientInterceptor),
		//		)
		//	},
		//}
		//pbTc.SetFunc(curConnKey, c)
		defer func() {
			_ = pbTc.Put(curConnKey, conn)
		}()
		return conn
	}
	conn := tcConn.(*grpc.ClientConn)
	return conn
}

// clientInterceptor 提供客户端的拦截器, 注入trace, 注入logId
func clientInterceptor(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	var x = make(map[string][]string)
	// 反射ctx, 判断是webContext, 还是grpcContext
	if oldCtx, ok := ctx.(*core.Context); ok {
		x[constants.X_REQUESTID] = []string{oldCtx.LogId()}
		oldCtx.Inject(context.TODO(), x)
	}

	// 转换key为小写不然rst
	var md = make(metautils.NiceMD)
	for k, v := range x {
		key := strings.ToLower(k)
		if len(v) > 0 {
			md.Set(key, v[0])
		}
	}

	newCtx := md.ToOutgoing(context.Background())
	return invoker(newCtx, method, req, reply, cc, opts...)
}
