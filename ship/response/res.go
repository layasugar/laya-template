package response

// 现做如下约定
// 1. 成功返回--{1."success",{},{}} code必须为1
// 2. 失败返回--{0,"系统发生错误！",{},{}} code 必须为0,code=0前端按照msg进行提示

type Response struct {
	Code     int32
	Msg      string
	Data     interface{}
	WithData interface{}
	Page     PageRes
}

type PageRes struct {
	CurPage int64 // 当前页
	Size    int64 // 每页条数
	Total   int64 // 总条数
}
