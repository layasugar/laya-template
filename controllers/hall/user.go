package main

import (
	"github.com/gin-gonic/gin"
	"laya-go/server/hall/model"
	"laya-go/server/hall/validate"
	"laya-go/ship"
	r "laya-go/ship/response"
	"laya-go/ship/utils"
	"strconv"
)

func Login(c *gin.Context) {
	var LoginRequestData validate.LoginData
	if err := c.ShouldBind(&LoginRequestData); err != nil {
		c.Set("$.Login.Params.code", r.ParamsValidateErr)
		return
	}
	user := model.User{}

	// 1.判断用户是否存在并取出用户信息
	if ship.DB.Model(&model.User{}).Where("phone = ?", LoginRequestData.Name).Where("zone = ?", LoginRequestData.Zone).Find(&user).RecordNotFound() {
		c.Set("$.Login.User.code", r.NotFoundUser)
		return
	}

	// 2.验证用户是否被冻结
	if user.Status == -1 {
		c.Set("$.Login.User.Freeze.code", r.FreezeUser)
		return
	}

	// 3.验证密码正确性
	if LoginRequestData.Password != user.Password {
		c.Set("$.Login.User.Password.code", r.PasswordErr)
		return
	}

	// 4.验证通过生成token,并写入redis
	strUid := strconv.FormatInt(user.ID, 10)
	token := utils.GetToken()
	oldToken, err := ship.Redis.HGet("user:token", strUid).Result()
	if err == nil {
		if err := ship.Redis.HDel("user:uid", oldToken).Err(); err != nil {
			c.Set("$.Login.Response.code", r.Err)
			return
		}
	}
	if err := ship.Redis.HSet("user:token", strUid, token).Err(); err != nil {
		c.Set("$.Login.Response.code", r.Err)
		return
	}
	if err := ship.Redis.HSet("user:uid", token, strUid).Err(); err != nil {
		c.Set("$.Login.Response.code", r.Err)
		return
	}

	data := map[string]interface{}{"ID": user.ID, "UserName": user.UserName, "Phone": user.Phone, "Zone": user.Zone, "Token": token}

	c.Set("$.Login.success.response", r.Response{Code: r.Success, Data: data})
}

//func TokenLogin(c *gin.Context) {
//	//获取语言类型
//	lang := c.GetHeader("Accept-Language")
//	language := ship.GetLang(lang)
//	var tld validate.TLD
//	if err := c.ShouldBind(&tld); err != nil {
//		c.JSON(http.StatusOK, ship.GetMessage("ParamErr", 40004, language, err))
//		return
//	}
//	token := tld.Token
//
//	uid, err := ship.Redis.HGet("user:uid", token).Result()
//	if err != nil {
//		c.JSON(http.StatusOK, ship.GetMessage("TokenErr", 40003, language, err))
//		return
//	}
//
//	ID, _ := strconv.ParseInt(uid, 10, 64)
//	user := model.User{ID: ID}
//
//	// 1.判断用户是否存在并取出用户信息
//	if ship.DB.Where(user).Find(&user).RecordNotFound() {
//		c.JSON(http.StatusOK, ship.GetMessage("NoUser", 40001, language, err))
//		return
//	}
//
//	// 1.1 验证用户是否被冻结
//	if user.Status == 1 {
//		c.JSON(http.StatusOK, ship.GetMessage("AccountFrozen", 40054, language, err))
//		return
//	}
//
//	// 2.刷新用户的token信息
//	newToken := ship.GetToken()
//	if err := ship.Redis.HDel("user:uid", token).Err(); err != nil {
//		c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, err))
//		return
//	}
//	if err := ship.Redis.HSet("user:token", strconv.FormatInt(user.ID, 10), newToken).Err(); err != nil {
//		c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, err))
//		return
//	}
//	if err := ship.Redis.HSet("user:uid", newToken, strconv.FormatInt(user.ID, 10)).Err(); err != nil {
//		c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, err))
//		return
//	}
//
//	// 3. 更新登录时间以及ip
//	ip := ship.RemoteIp(c.Request)
//	ship.DB.Model(&user).Where("id = ?", user.ID).Updates(model.User{LastIp: ip, LastTime: time.Now()})
//
//	data := map[string]interface{}{"ID": user.ID, "Nickname": user.Nickname, "Avatar": user.Avatar, "QQ": user.QQ, "Phone": user.Phone, "InviteCode": user.InviteCode, "Token": newToken}
//
//	c.JSON(http.StatusOK, ship.GetMessage("Success", 1, language, data))
//}
//
//func Register(c *gin.Context) {
//	//获取语言类型
//	lang := c.GetHeader("Accept-Language")
//	language := ship.GetLang(lang)
//	//千位分割符(金额)
//	ac := ship.Ac
//	var rd validate.RD
//	if err := c.ShouldBind(&rd); err != nil {
//		log.Info(err)
//		c.JSON(http.StatusOK, ship.GetMessage("ParamErr", 40004, language, err))
//		return
//	}
//	// 1. 用户名重复判断
//	// user := model.User{Account: rd.Username}
//	// if !ship.DB.Where(user).Find(&user).RecordNotFound() {
//	//	c.JSON(http.StatusOK, ship.UserExist)
//	//	return
//	// }
//
//	// 2.手机号重复判断
//	//user := model.User{Phone: rd.Phone, RegisterIp: ship.RemoteIp(c.Request)}
//	user := model.User{Phone: rd.Phone, Zone: rd.Zone}
//	if !ship.DB.Where(user).Find(&user).RecordNotFound() {
//		c.JSON(http.StatusOK, ship.GetMessage("PhoneExist", 40006, language, map[string]interface{}{}))
//		return
//	}
//
//	// 3.密码与重复密码校验
//	if rd.Password != rd.RPassword {
//		c.JSON(http.StatusOK, ship.GetMessage("PWDFail", 40007, language, map[string]interface{}{}))
//		return
//	}
//
//	// 5.短信验证码验证
//	//code, err := ship.Redis.Get("phone:code:" + rd.Zone + rd.Phone).Result()
//	//
//	//if err != nil {
//	//	c.JSON(http.StatusOK, ship.GetMessage("CodeErr", 40008, language, map[string]interface{}{}))
//	//	return
//	//}
//	//if code != rd.PhoneCode {
//	//	c.JSON(http.StatusOK, ship.GetMessage("PhoneCodeErr", 40009, language, map[string]interface{}{}))
//	//	return
//	//}
//
//	// 6.写入数据库并初始化钱包等级和余额宝
//	tx := ship.DB.Begin()
//	nickname := rd.Phone
//
//	var insertID []int64
//	ship.DB.Raw("select max(id) as id from tb_user").Pluck("id", &insertID)
//	var ID int64
//	ID = insertID[0] + 137
//	user = model.User{ID: ID, Account: rd.Phone, Nickname: nickname, Password: rd.Password, Phone: rd.Phone, CreateTime: time.Now(), LastTime: time.Now(), Status: 2, RegisterIp: ship.RemoteIp(c.Request), Zone: rd.Zone}
//
//	if rd.RemoteIp != "" {
//		user.RegisterIp = rd.RemoteIp
//	}
//
//	// 查询渠道信息
//	if rd.ChannelNumber != "" {
//		// 查询渠道详情
//		var channelInfo model.Channel
//		tx.Model(channelInfo).Where("number = ?", rd.ChannelNumber).First(&channelInfo)
//		if channelInfo.ID > 0 {
//			user.ChannelId = channelInfo.ID
//		}
//	}
//
//	if err := tx.Create(&user).Error; err != nil {
//		tx.Rollback()
//		c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, map[string]interface{}{}))
//		return
//	}
//	// 6.1 邀请码生成
//	inviteCode := ship.GetRandomString6(uint64(user.ID))
//	if err := tx.Model(&user).Update("invite_code", inviteCode).Error; err != nil {
//		tx.Rollback()
//		c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, map[string]interface{}{}))
//		return
//	}
//
//	// 6.2 余额宝初始化
//	yueBao := model.YEB{Uid: user.ID, Amount: 0}
//	if err := tx.Create(&yueBao).Error; err != nil {
//		tx.Rollback()
//		c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, map[string]interface{}{}))
//		return
//	}
//
//	// 6.3 用户会员等级初始化
//	userLevel := model.UL{Uid: user.ID, Level: 0}
//	if err := tx.Create(&userLevel).Error; err != nil {
//		tx.Rollback()
//		c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, map[string]interface{}{}))
//		return
//	}
//
//	// 6.7 获得注册奖励计算
//	var regConfig model.Config
//	tx.Where("id = ?", 1).Find(&regConfig)
//	regConfigVal, _ := strconv.ParseInt(regConfig.Val, 10, 64)
//
//	// 6.4 用户钱包初始化
//	userWallet := model.UW{Uid: user.ID, Balance: regConfigVal, FreezeBalance: 0}
//	if err := tx.Create(&userWallet).Error; err != nil {
//		tx.Rollback()
//		c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, map[string]interface{}{}))
//		return
//	}
//
//	// 注册奖励记录
//	walletLog := model.WalletLog{}
//	walletLog.Uid = user.ID
//	walletLog.Amount = regConfigVal
//	walletLog.CreateTime = time.Now()
//	walletLog.Ttype = 1
//	walletLog.Wtype = 26 // 注册奖励
//	if err := tx.Create(&walletLog).Error; err != nil {
//		tx.Rollback()
//		c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, map[string]interface{}{}))
//		return
//	}
//
//	// 4. 邀请码判断
//	if rd.InviteCode != "" {
//		agentUser := model.User{InviteCode: rd.InviteCode}
//		if ship.DB.Where(agentUser).Find(&agentUser).RecordNotFound() {
//			tx.Rollback()
//			c.JSON(http.StatusOK, ship.GetMessage("InviteCodeErr", 40010, language, map[string]interface{}{}))
//			return
//		} else {
//			// 6.5 邀请关系绑定
//			userInvite := model.UI{Uid: user.ID, AgentUid: agentUser.ID, CreateTime: time.Now()}
//			if err := tx.Create(&userInvite).Error; err != nil {
//				tx.Rollback()
//				c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, map[string]interface{}{}))
//				return
//			}
//			// 6.5.1 邀请人获得邀请奖励
//			var inviteConfig model.Config
//			tx.Where("id = ?", 2).Find(&inviteConfig)
//			inviteConfigVal, _ := strconv.ParseInt(inviteConfig.Val, 10, 64)
//			if inviteConfigVal > 0 {
//				//获取余额
//				var agentWallet model.UW
//				tx.Model(model.UW{}).Where("uid = ?", agentUser.ID).Find(&agentWallet)
//
//				if err := tx.Model(model.UW{}).Where("uid = ?", agentUser.ID).UpdateColumn("balance", gorm.Expr("balance + ?", inviteConfigVal)).Error; err != nil {
//					tx.Rollback()
//					c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, map[string]interface{}{}))
//					return
//				}
//
//				// 新增钱包明细
//				walletLog := model.WalletLog{}
//				walletLog.Uid = agentUser.ID
//				walletLog.Amount = inviteConfigVal
//				walletLog.CreateTime = time.Now()
//				walletLog.Ttype = 1
//				walletLog.Wtype = 23 // 邀请奖励
//				walletLog.BeforeAmount = agentWallet.Balance
//				if err := tx.Create(&walletLog).Error; err != nil {
//					tx.Rollback()
//					c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, map[string]interface{}{}))
//					return
//				}
//
//				notice := model.Notice{Uid: agentUser.ID, Title: "InvitePrize", Content: "CongratulationsInvitation|" + strconv.FormatInt(user.ID, 10) + "|" + ac.FormatMoney(inviteConfigVal), CreateTime: time.Now(), Status: 2}
//				if err := tx.Create(&notice).Error; err != nil {
//					tx.Rollback()
//					c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, map[string]interface{}{}))
//					return
//				}
//			}
//		}
//	}
//
//	// 6.6 发送系统消息
//	notice := model.Notice{Uid: user.ID, Title: "RegisteredSuccess", Content: "CongratulationsBecoming", CreateTime: time.Now(), Status: 2}
//	if err := tx.Create(&notice).Error; err != nil {
//		tx.Rollback()
//		c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, map[string]interface{}{}))
//		return
//	}
//
//	// 发送注册有奖消息
//	if regConfigVal > 0 {
//		notice := model.Notice{Uid: user.ID, Title: "AwardRegistration", Content: "CongratulationsWinning|" + ac.FormatMoney(regConfigVal), CreateTime: time.Now(), Status: 2}
//		if err := tx.Create(&notice).Error; err != nil {
//			tx.Rollback()
//			c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, map[string]interface{}{}))
//			return
//		}
//	}
//
//	tx.Commit()
//	// 7.返回数据
//	c.JSON(http.StatusOK, ship.GetMessage("Success", 1, language, map[string]interface{}{}))
//}
//
//var manager *captchas.Manager
//
//func getManager() *captchas.Manager {
//	if manager == nil {
//		store := redisstore.New(ship.Redis) // redis store
//		driver := drivers.NewString([]drivers.StringOption{
//			drivers.StringHeight(40),
//			drivers.StringWidth(120),
//			drivers.StringLength(4),
//			drivers.StringNoiseCount(0),
//			drivers.StringFonts([]string{}),
//			drivers.StringSource("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
//			drivers.StringBGColor(&color.RGBA{191, 211, 244, 255}),
//		}...)                                 // string driver
//		manager = captchas.New(store, driver) // manager
//	}
//	return manager
//}
//
//// 图形验证码
//func Captcha(c *gin.Context) {
//	//获取语言类型
//	lang := c.GetHeader("Accept-Language")
//	language := ship.GetLang(lang)
//	captcha, err := getManager().Generate()
//	if err != nil {
//		c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, map[string]interface{}{}))
//		return
//	}
//	data := map[string]interface{}{"ID": captcha.ID(), "Img": captcha.EncodeToString()}
//	c.JSON(http.StatusOK, ship.GetMessage("Success", 1, language, data))
//}
//
//// 短信验证码
//func Phone(c *gin.Context) {
//	//获取语言类型
//	lang := c.GetHeader("Accept-Language")
//	language := ship.GetLang(lang)
//	var p validate.P
//	if err := c.ShouldBind(&p); err != nil {
//		c.JSON(http.StatusOK, ship.GetMessage("ParamErr", 40004, language, map[string]interface{}{}))
//		return
//	}
//
//	if p.Type == 1 { // 注册
//
//		user := model.User{Phone: p.Phone, Zone: p.Zone}
//		if !ship.DB.Where(user).Find(&user).RecordNotFound() {
//			c.JSON(http.StatusOK, ship.GetMessage("PhoneExist", 40006, language, map[string]interface{}{}))
//			return
//		}
//	} else if p.Type == 2 { // 忘记密码
//		user := model.User{Phone: p.Phone, Zone: p.Zone}
//		if ship.DB.Where(user).Find(&user).RecordNotFound() {
//			c.JSON(http.StatusOK, ship.GetMessage("NoUser", 40001, language, map[string]interface{}{}))
//			return
//		}
//	}
//
//	//图形验证码
//	//p.Value = strings.ToUpper(p.Value)
//	//if err := getManager().Verify(p.ID, p.Value, true); err != nil {
//	//	c.JSON(http.StatusOK, ship.GetMessage("ImgCodeErr", 40055, language, map[string]interface{}{}))
//	//	return
//	//}
//
//	/* // 1. 策略手机号正确性判断
//	   iorgo, _ := regexp.MatchString(`^(1[3|5|7|8|9][0-9]\d{4,8})$`, p.Phone)
//	   if !iorgo {
//	       c.JSON(http.StatusOK, ship.GetMessage("PhoneErr",40011,language,map[string]interface{}{}))
//	       return
//	   }*/
//
//	// 2. 策略同个手机号60秒后才能发送
//	num1, _ := ship.Redis.Exists("phone:time:" + p.Zone + p.Phone).Result()
//	if num1 != 0 {
//		c.JSON(http.StatusOK, ship.GetMessage("MsgErr", 40012, language, map[string]interface{}{}))
//		return
//	}
//	if err := ship.Redis.Set("phone:time:"+p.Zone+p.Phone, "1", 60*time.Second).Err(); err != nil {
//		c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, map[string]interface{}{}))
//		return
//	}
//
//	// // 3. 策略一个手机号24小时只能发送5条
//	// num2, _ := ship.Redis.Get("phone:times:" + p.Phone).Result()
//	// intNum2, _ := strconv.ParseInt(num2, 10, 64)
//	// if intNum2 >= 5 {
//	//	c.JSON(http.StatusOK, ship.GetMessage("MsgFail",40013,language,map[string]interface{}{}))
//	//	return
//	// }
//	// intNum2 += 1
//	// if err := ship.Redis.Set("phone:times:"+p.Phone, intNum2, 24*time.Hour).Err(); err != nil {
//	//	 c.JSON(http.StatusOK, ship.GetMessage("Err",0,language,map[string]interface{}{}))
//	//	return
//	// }
//	//
//	// // 4. 策略同一个ip24小时只能发送10条
//	// num3, _ := ship.Redis.Get("phone:ip:" + ship.RemoteIp(c.Request)).Result()
//	// intNum3, _ := strconv.ParseInt(num3, 10, 64)
//	// if intNum3 >= 10 {
//	//	c.JSON(http.StatusOK, ship.GetMessage("MsgIpFail",40014,language,map[string]interface{}{}))
//	//	return
//	// }
//	// intNum3 += 1
//	// if err := ship.Redis.Set("phone:ip:"+ship.RemoteIp(c.Request), intNum3, 24*time.Hour).Err(); err != nil {
//	//	 c.JSON(http.StatusOK, ship.GetMessage("Err",0,language,map[string]interface{}{}))
//	//	return
//	// }
//
//	// 6. 写入redis，300秒过期
//	code := ship.GenValidateCode(6)
//
//	if err := ship.Redis.Set("phone:code:"+p.Zone+p.Phone, code, 5*time.Minute).Err(); err != nil {
//		c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, map[string]interface{}{}))
//		return
//	}
//
//	// 7. 执行发送信息
//	sendResult := SendPhoneCode1(code, p.Phone, p.Zone, p.Type)
//	if sendResult.Code != 0 {
//		c.JSON(http.StatusOK, ship.GetMessage("PhoneCodeFail", 40025, language, sendResult))
//		return
//	}
//
//	// 8. 成功返回
//	c.JSON(http.StatusOK, ship.GetMessage("Success", 1, language, sendResult))
//}
//
//// 修改用户密码
//func EditUserPwd(c *gin.Context) {
//	//获取语言类型
//	lang := c.GetHeader("Accept-Language")
//	language := ship.GetLang(lang)
//	var phone = c.PostForm("Phone")
//	var Code = c.PostForm("Code")
//	var Zone = c.PostForm("Zone")
//	var NewPassword = c.PostForm("NewPassword")
//	var ReNewPassword = c.PostForm("ReNewPassword")
//
//	// 5.短信验证码验证
//	code, err := ship.Redis.Get("phone:code:" + Zone + phone).Result()
//	if err != nil {
//		c.JSON(http.StatusOK, ship.GetMessage("CodeErr", 40008, language, map[string]interface{}{}))
//		return
//	}
//	if code != Code {
//		c.JSON(http.StatusOK, ship.GetMessage("PhoneCodeErr", 40009, language, map[string]interface{}{}))
//		return
//	}
//
//	// 判断重复密码是否一致
//	if NewPassword != ReNewPassword {
//		c.JSON(http.StatusOK, ship.GetMessage("PWDFail", 40007, language, map[string]interface{}{}))
//		return
//	}
//
//	// 修改密码
//	ship.DB.Model(&model.User{}).Where("phone = ?", phone).Where("zone = ?", Zone).Update("password", NewPassword)
//
//	// 返回信息=
//	c.JSON(http.StatusOK, ship.GetMessage("Success", 1, language, map[string]interface{}{}))
//}

// 获取用户信息
func GetUserInfo(c *gin.Context) {
	uid := c.GetInt64("uid")
	user := model.User{ID: uid}
	if ship.DB.Model(&model.User{}).Where(&user).Find(&user).RecordNotFound() {
		c.Set("$.err", r.NotFoundUser)
		return
	}
	c.Set("$.data", user)
}

//// 修改用户信息
//func EditUser(c *gin.Context) {
//	//获取语言类型
//	lang := c.GetHeader("Accept-Language")
//	language := ship.GetLang(lang)
//	uid := c.GetInt64("uid")
//	var etype = c.PostForm("etype") // 修改类型枚举
//	var nickname = c.PostForm("nickname")
//	var avatar = c.PostForm("avatar")
//	var qq = c.PostForm("qq")
//	var phone = c.PostForm("phone")
//	var zone = c.PostForm("zone")
//
//	user := model.User{}
//	user.ID = uid
//
//	if ship.DB.Where(user).Find(&user).RecordNotFound() {
//		c.JSON(http.StatusOK, ship.GetMessage("NoUser", 40001, language, map[string]interface{}{}))
//		return
//	}
//
//	switch etype {
//	// 修改昵称
//	case "1":
//		if nickname == "" {
//			c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, map[string]interface{}{}))
//			return
//		}
//		user.Nickname = nickname
//		break
//	// 修改头像
//	case "2":
//		if avatar == "" {
//			c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, map[string]interface{}{}))
//			return
//		}
//		user.Avatar = avatar
//		break
//	// 修改QQ
//	case "3":
//		if qq == "" {
//			c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, map[string]interface{}{}))
//			return
//		}
//		user.QQ = qq
//		break
//	// 修改手机号
//	case "4":
//		if phone == "" || zone == "" {
//			c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, map[string]interface{}{}))
//			return
//		}
//		user.Phone = phone
//		user.Zone = zone
//		break
//	}
//
//	if err := ship.DB.Save(&user).Error; err != nil {
//		c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, map[string]interface{}{}))
//		return
//	}
//
//	c.JSON(http.StatusOK, ship.GetMessage("Success", 1, language, map[string]interface{}{}))
//}
//
//// 修改密码
//func EditPwd(c *gin.Context) {
//	//获取语言类型
//	lang := c.GetHeader("Accept-Language")
//	language := ship.GetLang(lang)
//	uid := c.GetInt64("uid")
//	var oldpassword = c.PostForm("oldpassword")
//	var newpassword = c.PostForm("newpassword")
//	var renewpassword = c.PostForm("renewpassword")
//	var Phone = c.PostForm("phone")
//	var Zone = c.PostForm("zone")
//	var PhoneCode = c.PostForm("PhoneCode")
//
//	if PhoneCode == "" {
//		c.JSON(http.StatusOK, ship.GetMessage("PhoneCodeErr", 40009, language, map[string]interface{}{}))
//		return
//	}
//	// 5.短信验证码验证
//	code, err := ship.Redis.Get("phone:code:" + Zone + Phone).Result()
//	if err != nil {
//		c.JSON(http.StatusOK, ship.GetMessage("CodeErr", 40008, language, map[string]interface{}{}))
//		return
//	}
//	if code != PhoneCode {
//		c.JSON(http.StatusOK, ship.GetMessage("PhoneCodeErr", 40009, language, map[string]interface{}{}))
//		return
//	}
//
//	if len(oldpassword) != 32 || len(newpassword) != 32 || len(renewpassword) != 32 {
//		c.JSON(http.StatusOK, ship.GetMessage("ParamErr", 40004, language, map[string]interface{}{}))
//		return
//	}
//
//	// 判断重复密码是否一致
//	if newpassword != renewpassword {
//		c.JSON(http.StatusOK, ship.GetMessage("PWDFail", 40007, language, map[string]interface{}{}))
//		return
//	}
//
//	user := model.User{ID: uid, Password: oldpassword}
//
//	if ship.DB.Where(user).Find(&user).RecordNotFound() {
//		c.JSON(http.StatusOK, ship.GetMessage("PWDErr", 40007, language, map[string]interface{}{}))
//		return
//	}
//
//	// 修改密码
//	user.Password = newpassword
//	if err := ship.DB.Save(&user).Error; err != nil {
//		c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, map[string]interface{}{}))
//		return
//	}
//
//	// 返回信息
//	c.JSON(http.StatusOK, ship.GetMessage("Success", 1, language, map[string]interface{}{}))
//}
//
//// 签到
//func SignIn(c *gin.Context) {
//	//获取语言类型
//	lang := c.GetHeader("Accept-Language")
//	language := ship.GetLang(lang)
//	uid := c.GetInt64("uid")
//	ymd := time.Now().Format("2006-01-02")
//	//千位分割符(金额)
//	ac := ship.Ac
//
//	// 获取锁开启事务
//	lockKey := "SignIn:lock:" + strconv.FormatInt(uid, 10)
//	lock, err := ship.GetLock(ship.Redis, lockKey, 5*time.Second, 5*time.Second)
//	if err != nil {
//		c.JSON(http.StatusOK, ship.GetMessage("SysErr", 40024, language, map[string]interface{}{}))
//		return
//	}
//
//	// 1.判断是否已经签到了
//	usl := model.USL{Uid: uid, Ymd: ymd}
//	if !ship.DB.Model(model.USL{}).Where(usl).Find(&usl).RecordNotFound() {
//		ship.ReleaseLock(ship.Redis, lockKey, lock)
//		c.JSON(http.StatusOK, ship.GetMessage("Signed", 40059, language, map[string]interface{}{}))
//		return
//	}
//	//response := ship.Response{Code: 1, Msg: map[string]interface{}{"Chinese":"成功","English":"success"}}
//	var Data interface{}
//	tx := ship.DB.Begin()
//	// 2. 赠送奖励
//	config := model.Config{}
//	tx.Model(config).Where("id = ?", 3).Find(&config)
//	if config.Status == 1 {
//		userWallet := model.UW{Uid: uid}
//		tx.Where(userWallet).Find(&userWallet)
//		amount, _ := strconv.ParseInt(config.Val, 10, 64)
//		//userWallet.Balance += amount
//		usl.Amount = amount
//		walletLog := model.WalletLog{Uid: uid, Wtype: 19, Ttype: 1, Amount: amount, CreateTime: time.Now(), BeforeAmount: userWallet.Balance}
//		if err := tx.Create(&walletLog).Error; err != nil {
//			tx.Rollback()
//			ship.ReleaseLock(ship.Redis, lockKey, lock)
//			c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, err))
//			return
//		}
//		if err := tx.Model(model.UW{}).Where(userWallet).UpdateColumn("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
//			tx.Rollback()
//			ship.ReleaseLock(ship.Redis, lockKey, lock)
//			c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, err))
//			return
//		}
//
//		// 发送消息通知
//		notice := model.Notice{Uid: uid, Title: "SignedSuccess", Status: 2, CreateTime: time.Now()}
//		notice.Content = "SignedInSuccess|" + ac.FormatMoney(amount)
//		if err := tx.Create(&notice).Error; err != nil {
//			tx.Rollback()
//			ship.ReleaseLock(ship.Redis, lockKey, lock)
//			c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, err))
//			return
//		}
//
//		Data = amount
//
//	} else {
//		Data = 0
//	}
//
//	usl.CreateTime = time.Now()
//	if err := tx.Create(&usl).Error; err != nil {
//		tx.Rollback()
//		ship.ReleaseLock(ship.Redis, lockKey, lock)
//		c.JSON(http.StatusOK, ship.GetMessage("Err", 0, language, err))
//		return
//	}
//
//	tx.Commit()
//	ship.ReleaseLock(ship.Redis, lockKey, lock)
//
//	c.JSON(http.StatusOK, ship.GetMessage("Success", 1, language, Data))
//}
