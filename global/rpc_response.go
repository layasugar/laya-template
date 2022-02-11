package global

import "github.com/layasugar/laya"

type RpcResp struct{}

func (res *RpcResp) SucRpc(ctx *laya.GrpcContext, data interface{}, msg ...string) {

}

func (res *RpcResp) FailRpc(ctx *laya.GrpcContext, err error) {

}
