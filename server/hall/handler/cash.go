package handler
//
//import (
//	"github.com/gin-gonic/gin"
//	"laya-go/server/hall/pay"
//	"laya-go/ship"
//	"laya-go/ship/model"
//	"laya-go/ship/validate"
//	"log"
//	"net/http"
//	"strconv"
//	"time"
//)
//
//func CreateOrder(c *gin.Context) {
//	lang := c.GetHeader("Accept-Language")
//	language := ship.GetLang(lang)
//
//	var params validate.CreateOrder
//	if err := c.ShouldBind(&params); err != nil {
//		c.JSON(http.StatusOK, ship.GetMessage("ParamErr", 40004, language, err))
//		return
//	}
//	uid := c.GetInt64("uid")
//
//	//获取支付的对象
//	PayObj, ok := GetPay(params.PayType)
//	if !ok {
//		c.JSON(http.StatusOK, ship.GetMessage("PayChannelClose", 0, language, map[string]interface{}{}))
//		return
//	}
//
//	//调用渠道下单
//	resp, err := PayObj.Instance.CreateOrder(params, pay.CreateOrderParams{Prefix: PayObj.Prefix, Uid: uid})
//	if err != nil {
//		c.JSON(http.StatusOK, ship.GetMessage("PayChannelErr", 0, language, err))
//		return
//	}
//
//	//查询玩家信息
//	user := model.User{ID: uid}
//	ship.DB.Where(user).Find(&user)
//	userWallet := model.UW{Uid: uid}
//	ship.DB.Where(userWallet).Find(&userWallet)
//
//	//写入数据库
//	cashLog := model.Cash{
//		Uid:          uid,
//		Name:         user.RealName,
//		ImageUrl:     PayObj.Prefix,
//		Mtype:        1, //充值
//		OrderNo:      resp.OrderNo,
//		ThirdOrderNo: resp.ThirdOrderNo,
//		Amount:       int64(params.Amount),
//		GiveAmount:   0,
//		CreateTime:   time.Now(),
//		URemark:      strconv.Itoa(params.BankType) + ":" + params.Remark,
//		Status:       1,
//		ApplyTime:    time.Now(),
//		PayType:      int64(params.PayType),
//		Balance:      userWallet.Balance,
//	}
//	if err := ship.DB.Create(&cashLog).Error; err != nil {
//		c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, err))
//		return
//	}
//
//	c.JSON(http.StatusOK, ship.GetMessage("Success", 0, language, resp))
//	return
//}
//
////回调
//func Notify(c *gin.Context) {
//	channel := c.Param("channel")
//	log.Println(channel)
//}
//
////支付对象
//type PayObj struct {
//	Prefix   string
//	Instance pay.Pay
//}
//
////获取支付对象
//func GetPay(PayType int) (PayObj, bool) {
//
//	cashType := model.PayType{Status: 1, ID: int64(PayType)}
//
//	ship.DB.Where(cashType).Find(&cashType)
//
//	//支付渠道列表
//	Instances := map[string]pay.Pay{
//		"TT": new(pay.TTPay),
//	}
//
//	PayInstance, ok := Instances[cashType.AccountName]
//	return PayObj{Prefix: cashType.AccountName, Instance: PayInstance}, ok
//}
