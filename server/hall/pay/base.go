package pay
//
//import (
//	"encoding/json"
//	"github.com/jinzhu/gorm"
//	"io/ioutil"
//	"laya-go/ship"
//	"laya-go/ship/model"
//	"laya-go/ship/validate"
//	"net/http"
//	"strings"
//	"time"
//)
//
////支付接口
//type Pay interface {
//	GetConfig() Config
//	sign(params map[string]interface{}) (string, string)
//	CreateOrder(Params validate.CreateOrder, Extra CreateOrderParams) (CreateOrderReturn, interface{})
//}
//
////商户配置
//type Config struct {
//	Mid     string
//	ReqHost string
//	Secret  string
//}
//
////下单返回的公共内容
//type CreateOrderReturn struct {
//	OrderNo      string
//	ThirdOrderNo string
//	ThirdData    interface{}
//}
//
////下单前额外的参数
//type CreateOrderParams struct {
//	Prefix string
//	Uid    int64
//}
//
//type NotifyParams struct {
//}
//
////完成订单
//func NotifyFishOrder(cash model.Cash) error {
//
//	//千位分割符(金额)
//	ac := ship.Ac
//
//	//开启事务
//	tx := ship.DB.Begin()
//
//	//管理员ID
//	adminUid := 0
//
//	// 计算赠送金额
//	cashConfig := model.CashConfig{}
//	//给自己返利
//	if ship.DB.Where("amount <= ?", cash.Amount).Order("amount DESC").First(&cashConfig).RecordNotFound() {
//		cashConfig.SelfRebate = 0
//	}
//	user := model.User{ID: cash.Uid}
//	if ship.DB.Where(user).Where("is_rebate = ?", 0).Find(&user).RecordNotFound() {
//		cashConfig.SelfRebate = 0
//	}
//
//	// giveAmount := cashConfig.Proportion * amount
//	cash.GiveAmount = cashConfig.SelfRebate
//
//	//修改充值表信息
//	data := model.Cash{Status: 2, Remark: "", ApplyTime: time.Now(), IsIgnore: 2, AdminUid: int64(adminUid), GiveAmount: cash.GiveAmount}
//	if err := tx.Model(&cash).Update(&data).Error; err != nil {
//		tx.Rollback()
//		return err
//	}
//
//	//写入钱包流水
//	userWallet := model.UW{Uid: cash.Uid}
//	tx.Where(userWallet).Find(&userWallet)
//
//	//1、充值金额
//	walletLog := model.WalletLog{Uid: cash.Uid, Wtype: 1, Ttype: 1, Amount: cash.Amount, CreateTime: time.Now(), BeforeAmount: userWallet.Balance}
//
//	if err := tx.Create(&walletLog).Error; err != nil {
//		tx.Rollback()
//		return err
//	}
//
//	//2、赠送金额
//	if cash.GiveAmount > 0 {
//		//首冲给自己返利记录
//		walletLog2 := model.WalletLog{Uid: cash.Uid, Wtype: 14, Ttype: 1, Amount: cash.GiveAmount, CreateTime: time.Now(), BeforeAmount: userWallet.Balance}
//		if err := tx.Create(&walletLog2).Error; err != nil {
//			tx.Rollback()
//			return err
//		}
//
//		//首冲给上级送金
//		var userInviteInfo model.UserInviteInfo
//		if !tx.Model(&model.UserInviteInfo{}).Preload("UserInfo").Where("uid = ?", cash.Uid).Find(&userInviteInfo).RecordNotFound() {
//
//			//上级钱包
//			inviteWallet := model.UW{Uid: userInviteInfo.UserInfo.ID}
//			tx.Where(inviteWallet).Find(&inviteWallet)
//
//			//首冲给上级返利记录
//			walletLog3 := model.WalletLog{Uid: userInviteInfo.UserInfo.ID, Wtype: 30, Ttype: 1, Amount: cashConfig.ParentRebate, CreateTime: time.Now(), BeforeAmount: inviteWallet.Balance}
//			if err := tx.Create(&walletLog3).Error; err != nil {
//				tx.Rollback()
//				return err
//			}
//
//			//给上级加钱
//			if err := tx.Model(&model.UW{}).
//				Where("uid = ?", walletLog3.Uid).
//				UpdateColumn("balance", gorm.Expr("balance + ?", cashConfig.ParentRebate)).Error; err != nil {
//				tx.Rollback()
//				return err
//			}
//
//			//给上级发送系统消息
//			ParentNotice := model.Notice{Uid: userInviteInfo.UserInfo.ID, Title: "RechargedSuccess", Content: "PayParentRebate", CreateTime: time.Now(), Status: 2}
//			tx.Create(&ParentNotice)
//		}
//
//		//标记已参与首冲送金
//		user.IsRebate = 1
//		if err := tx.Model(&user).Updates(&user).Error; err != nil {
//			tx.Rollback()
//			return err
//		}
//	}
//
//	//修改钱包金额
//	wallet := model.UW{Uid: cash.Uid}
//	Balance := cash.Amount + cash.GiveAmount
//
//	if err := tx.Model(&wallet).Where(wallet).UpdateColumn("balance", gorm.Expr("balance + ?", Balance)).Error; err != nil {
//		tx.Rollback()
//		return err
//	}
//
//	// 发送系统消息
//	notice := model.Notice{Uid: cash.Uid, Title: "RechargedSuccess", Content: "CongratulationsSuccessfulRecharge|" + ac.FormatMoney(Balance), CreateTime: time.Now(), Status: 2}
//	tx.Create(&notice)
//
//	//提交事务
//	tx.Commit()
//
//	return nil
//}
//
////响应
//type Response struct {
//	Code      string              `json:"code"`
//	Message   string              `json:"message"`
//	List      []map[string]string `json:"list,omitempty"`
//	Data      []map[string]string `json:"data,omitempty"`
//	OrderNo   string              `json:"order_no,omitempty"`
//	POrderNo  string              `json:"p_order_no,omitempty"`
//	PayUrl    string              `json:"pay_url,omitempty"`
//	Amount    int                 `json:"amount,omitempty"`
//	PayAmount int                 `json:"pay_amount,omitempty"`
//	Nonce     string              `json:"nonce,omitempty"`
//	Sign      string              `json:"sign,omitempty"`
//}
//
////发起请求
//func HttpReq(url string, data string) (Response, interface{}) {
//
//	var resp Response
//	var err interface{}
//
//	//创建请求
//	client := &http.Client{}
//	request, er := http.NewRequest("POST", TTPay{}.GetConfig().ReqHost+url, strings.NewReader(data))
//	if er != nil {
//		return resp, er
//	}
//
//	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//
//	response, _ := client.Do(request)
//
//	defer response.Body.Close()
//
//	res, _ := ioutil.ReadAll(response.Body)
//
//	json.Unmarshal(res, &resp)
//
//	if resp.Code != "0000" {
//		err = resp.Message
//	}
//	return resp, err
//}
