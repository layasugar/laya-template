package response

// 现做如下约定
// 1. 成功返回--{1."success",{},{}} code必须为1
// 2. 失败返回--{0,"系统发生错误！",{},{}} code 必须为0,code=0前端按照msg进行提示

type Response struct {
	Code     int
	Msg      string
	Data     interface{}
	WithData interface{}
	Page     PageRes
	Location string
}

type PageRes struct {
	CurPage int // 当前页
	Size    int // 每页条数
	Total   int // 总条数
}
