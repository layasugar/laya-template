package pay
//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"github.com/micro/go-micro/v2/util/log"
//	"laya-go/ship"
//	"laya-go/ship/model"
//	"laya-go/ship/validate"
//	"net/http"
//	"sort"
//	"strconv"
//	"strings"
//	"time"
//)
//
//type TTPay struct {
//}
//
////获取配置
//func (p TTPay) GetConfig() Config {
//	return Config{
//		Mid:     "1009",
//		ReqHost: "https://vn.kkvox.com/api/v1",
//		Secret:  "9406dbb9-5662-4a5a-9ae9-ecc13419de88",
//	}
//}
//
//func (p TTPay) CreateOrder(Params validate.CreateOrder, Extra CreateOrderParams) (CreateOrderReturn, interface{}) {
//	uidStr := strconv.FormatInt(Extra.Uid, 10)
//
//	FormData := map[string]interface{}{
//		"mid":        p.GetConfig().Mid,
//		"order_no":   Extra.Prefix + uidStr + time.Now().Format("20060102150405") + ship.RandString(4),
//		"pay_type":   1,               //支付方式，复制银行转账:1
//		"bank_type":  Params.BankType, //银行类型，详见银行列表
//		"amount":     float64(Params.Amount),
//		"notify_url": ship.APIHOST + "/hall/cash/ttNotify",
//		"extra":      uidStr, //支付中心回调时会原样返回
//		"nonce":      ship.RandString(8),
//	}
//
//	_, reqData := p.sign(FormData)
//
//	resp, err := HttpReq("/open/CreateOrder", reqData)
//	if err != nil {
//		return CreateOrderReturn{}, err
//	}
//	res := CreateOrderReturn{
//		OrderNo:      FormData["order_no"].(string),
//		ThirdOrderNo: resp.POrderNo,
//		ThirdData:    resp,
//	}
//	return res, nil
//}
//
////签名
//func (p TTPay) sign(params map[string]interface{}) (string, string) {
//	var paramStr string
//	var keys []string
//	for key, _ := range params {
//		keys = append(keys, key)
//	}
//	//排序
//	sort.Strings(keys)
//	//字符串拼接
//	for _, k := range keys {
//		if k != "sign" {
//			if k == "amount" {
//				paramStr += k + "=" + strings.Replace(fmt.Sprintf("%12.0f", params[k]), " ", "", -1) + "&"
//			} else {
//				paramStr += k + "=" + fmt.Sprintf("%v", params[k]) + "&"
//			}
//		}
//	}
//	requestParamsStr := paramStr + "sign="
//
//	//去掉最后一个&
//	paramStr = paramStr[0 : len(paramStr)-1]
//
//	paramStr += p.GetConfig().Secret
//
//	log.Info(paramStr)
//
//	//MD5再转大写
//	signStr := strings.ToUpper(ship.MD5(paramStr))
//	requestParamsStr += signStr
//	return signStr, requestParamsStr
//}
//
////回调同步返回
//type res struct {
//	Code    string      `json:"code"`
//	Message interface{} `json:"message"`
//}
//
//type notify struct {
//	Code     string `form:"code" json:"code"`
//	Message  string `form:"message" json:"message"`
//	Mid      string `form:"mid" json:"mid"`
//	OrderNo  string `form:"order_no" json:"order_no"`
//	POrderNo string `form:"p_order_no" json:"p_order_no"`
//	Amount   int    `form:"amount" json:"amount"`
//	Extra    string `form:"extra" json:"extra"`
//	PayTime  string `form:"pay_time" json:"pay_time"`
//	Nonce    string `form:"nonce" json:"nonce"`
//	Sign     string `form:"sign" json:"sign"`
//}
//
////天天回调
//func TTNotify(c *gin.Context) {
//	uuid := ship.GetUUID()
//	log.Info("回调开始：", uuid)
//	var params notify
//	if err := c.ShouldBind(&params); err != nil {
//		c.JSON(http.StatusOK, res{Code: "1111", Message: "参数错误"})
//		return
//	}
//	paramsStr, _ := json.Marshal(params)
//	log.Info(string(paramsStr), uuid)
//
//	//验签
//	var paramsMap map[string]interface{}
//	json.Unmarshal(paramsStr, &paramsMap)
//	sign, _ := TTPay{}.sign(paramsMap)
//
//	if sign != params.Sign {
//		log.Info("签名：")
//		log.Info(sign, params, uuid)
//		c.JSON(http.StatusOK, res{Code: "1111", Message: "签名错误"})
//		return
//	}
//
//	log.Info(sign, params, uuid)
//
//	//查询订单
//	order := model.Cash{OrderNo: params.OrderNo}
//	if ship.DB.Where(order).Find(&order).RecordNotFound() {
//		log.Info("订单不存在")
//		c.JSON(http.StatusOK, res{Code: "1111", Message: "订单号不存在"})
//		return
//	}
//	log.Info(order, uuid)
//	if order.Status != 1 {
//		c.JSON(http.StatusOK, res{Code: "1111", Message: "订单重复处理"})
//		return
//	}
//
//	if params.Code == "0000" {
//		err := NotifyFishOrder(order)
//		if err != nil {
//			c.JSON(http.StatusOK, res{Code: "1111", Message: err})
//			return
//		}
//	} else {
//
//		// 发送系统消息
//		notice := model.Notice{Uid: order.Uid, Title: "PrepaidRefuse", Content: "PrepaidOrderRefuse|" + params.Message + "|" + order.OrderNo, CreateTime: time.Now(), Status: 2}
//		ship.DB.Create(&notice)
//
//		//日志
//		detail := "RechargeAuditRejected|" + order.OrderNo + "|" + params.Message
//		ship.ManagerLogs(c.GetInt64("uid"), 6, detail, ship.RemoteIp(c.Request))
//	}
//	log.Info("回调结束", uuid)
//	c.JSON(http.StatusOK, res{Code: "0000", Message: "Success"})
//	return
//}
//
////获取银行列表
//func GetTTBankList(c *gin.Context) {
//	lang := c.GetHeader("Accept-Language")
//	language := ship.GetLang(lang)
//
//	FormData := map[string]interface{}{
//		"mid":   TTPay{}.GetConfig().Mid,
//		"nonce": ship.RandString(8),
//	}
//
//	_, reqData := TTPay{}.sign(FormData)
//
//	resp, err := HttpReq("/open/BankList", reqData)
//	if err != nil {
//		c.JSON(http.StatusOK, ship.GetMessage("Err", 1, language, err))
//		return
//	}
//
//	c.JSON(http.StatusOK, ship.GetMessage("Success", 0, language, resp.List))
//	return
//}
