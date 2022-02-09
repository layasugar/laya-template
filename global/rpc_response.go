package global

import "github.com/layasugar/laya"

type RpcResp struct{}

func (res *RpcResp) SucRpc(ctx *laya.PbRPCContext, data interface{}, msg ...string) {

}

func (res *RpcResp) FailRpc(ctx *laya.PbRPCContext, err error) {

}
