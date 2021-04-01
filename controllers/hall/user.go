package hall

import (
	"github.com/gin-gonic/gin"
)

func (ctrl *controller) Login(c *gin.Context) {
	//var LoginRequestData LoginData
	//if err := c.ShouldBind(&LoginRequestData); err != nil {
	//	c.Set("$.Login.Params.code", response.ParamsValidateErr)
	//	return
	//}
	//user := db.User{}
	//
	//// 1.判断用户是否存在并取出用户信息
	//if result := db.Dao.Model(&user).Where("phone = ?", LoginRequestData.Name).Where("zone = ?", LoginRequestData.Zone).First(&user);
	//	errors.Is(result.Error, gorm.ErrRecordNotFound) {
	//	c.Set("$.Login.User.code", response.NotFoundUser)
	//	return
	//}
	//
	//// 2.验证用户是否被冻结
	//if user.Status == -1 {
	//	c.Set("$.Login.User.Freeze.code", response.FreezeUser)
	//	return
	//}
	//
	//// 3.验证密码正确性
	//if LoginRequestData.Password != user.Password {
	//	c.Set("$.Login.User.Password.code", response.PasswordErr)
	//	return
	//}
	//
	//// 4.验证通过生成token,并写入redis
	//strUid := strconv.FormatInt(user.ID, 10)
	//token := "sad"
	//oldToken, err := rdb.Dao.HGet(context.Background(), "user:token", strUid).Result()
	//if err == nil {
	//	if err := rdb.Dao.HDel(context.Background(), "user:uid", oldToken).Err(); err != nil {
	//		c.Set("$.Login.Response.code", response.Err)
	//		return
	//	}
	//}
	//if err := rdb.Dao.HSet(context.Background(), "user:token", strUid, token).Err(); err != nil {
	//	c.Set("$.Login.Response.code", response.Err)
	//	return
	//}
	//if err := rdb.Dao.HSet(context.Background(), "user:uid", token, strUid).Err(); err != nil {
	//	c.Set("$.Login.Response.code", response.Err)
	//	return
	//}
	//
	//data := map[string]interface{}{"ID": user.ID, "UserName": user.UserName, "Phone": user.Phone, "Zone": user.Zone, "Token": token}
	//fmt.Println(data)
	////c.Set("$.Login.success.response", response.Response{Code: response.Success, Data: data})
}

func (ctrl *controller) TokenLogin(c *gin.Context) {
	//fmt.Println(response.ErrSysErr)
	//
	//ctrl.Fail(c, response.ErrSysErr)
	//return
	//var tld TokenLoginData
	//if err := c.ShouldBind(&tld); err != nil {
	//	ctrl.Fail(c, response.ErrSysErr)
	//	return
	//}
	//token := tld.Token
	//
	//uid, err := rdb.Dao.HGet(context.Background(), "user:uid", token).Result()
	//if err != nil {
	//	c.Set("$.Login.Params.code", response.ParamsValidateErr)
	//	return
	//}
	//
	//ID, _ := strconv.ParseInt(uid, 10, 64)
	//user := db.User{ID: ID}
	//
	//// 1.判断用户是否存在并取出用户信息
	//if result := db.Dao.Model(&user).Where(user).First(&user);
	//	errors.Is(result.Error, gorm.ErrRecordNotFound) {
	//	c.Set("$.Login.Params.code", response.ParamsValidateErr)
	//	return
	//}
	//
	//// 1.1 验证用户是否被冻结
	//if user.Status == 1 {
	//	c.Set("$.Login.Params.code", response.ParamsValidateErr)
	//	return
	//}
	//
	//// 2.刷新用户的token信息
	//newToken := "safdas"
	//if err := rdb.Dao.HDel(context.Background(), "user:uid", token).Err(); err != nil {
	//	c.Set("$.Login.Params.code", response.ParamsValidateErr)
	//	return
	//}
	//if err := rdb.Dao.HSet(context.Background(), "user:token", strconv.FormatInt(user.ID, 10), newToken).Err(); err != nil {
	//	c.Set("$.Login.Params.code", response.ParamsValidateErr)
	//	return
	//}
	//if err := rdb.Dao.HSet(context.Background(), "user:uid", newToken, strconv.FormatInt(user.ID, 10)).Err(); err != nil {
	//	c.Set("$.Login.Params.code", response.ParamsValidateErr)
	//	return
	//}
	//
	//// 3. 更新登录时间以及ip
	//ip := utils.RemoteIp(c.Request)
	//db.Dao.Model(&user).Where("id = ?", user.ID).Updates(db.User{LastLoginIp: ip, LastLoginTime: time.Now()})
	//
	//data := map[string]interface{}{"ID": user.ID, "UserName": user.UserName, "Zone": user.Zone, "Phone": user.Phone, "InviteCode": user.InviteCode, "Token": newToken}
	//
	//c.Set("$.TokenLogin.response", data)
}

func (ctrl *controller) Register(c *gin.Context) {
	//var rd RegisterData
	//if err := c.ShouldBind(&rd); err != nil {
	//	c.Set("$.Login.TokenErr.code", response.ParamsValidateErr)
	//	return
	//}
	//
	//// 2.手机号重复判断
	//user := db.User{Phone: rd.Phone, Zone: rd.Zone}
	//if result := db.Dao.Where(user).Find(&user);
	//	!errors.Is(result.Error, gorm.ErrRecordNotFound) {
	//	c.Set("$.Login.Params.code", response.ParamsValidateErr)
	//	return
	//}
	//
	//// 3.密码与重复密码校验
	//if rd.Password != rd.RPassword {
	//	c.Set("$.Login.Params.code", response.ParamsValidateErr)
	//	return
	//}
	//
	//// 5.短信验证码验证
	////code, err := rdb.Dao.Get("phone:code:" + rd.Zone + rd.Phone).Result()
	////
	////if err != nil {
	////	c.JSON(http.StatusOK, laya.GetMessage("CodeErr", 40008, language, map[string]interface{}{}))
	////	return
	////}
	////if code != rd.PhoneCode {
	////	c.JSON(http.StatusOK, laya.GetMessage("PhoneCodeErr", 40009, language, map[string]interface{}{}))
	////	return
	////}
	//
	//// 6.写入数据库并初始化钱包等级和余额宝
	//tx := db.Dao.Begin()
	//nickname := rd.Phone
	//
	//var insertID []int64
	//db.Dao.Raw("select max(id) as id from tb_user").Pluck("id", &insertID)
	//var ID int64
	//ID = insertID[0] + 137
	//user = db.User{ID: ID, UserName: nickname, Password: rd.Password, Phone: rd.Phone, CreatedAt: time.Now(), LastLoginTime: time.Now(), Status: 2, LastLoginIp: utils.RemoteIp(c.Request), Zone: rd.Zone}
	//
	//if err := tx.Create(&user).Error; err != nil {
	//	tx.Rollback()
	//	c.Set("$.Login.Params.code", response.ParamsValidateErr)
	//	return
	//}
	//
	//// 6.1 邀请码生成
	//inviteCode := utils.GetRandomString6(uint64(user.ID))
	//if err := tx.Model(&user).Update("invite_code", inviteCode).Error; err != nil {
	//	tx.Rollback()
	//	c.Set("$.Login.Params.code", response.ParamsValidateErr)
	//	return
	//}

	//// 6.3 用户会员等级初始化
	//userLevel := db.UL{Uid: user.ID, Level: 0}
	//if err := tx.Create(&userLevel).Error; err != nil {
	//	tx.Rollback()
	//	c.JSON(http.StatusOK, laya.GetMessage("Err", 0, language, map[string]interface{}{}))
	//	return
	//}
	//
	//// 6.4 用户钱包初始化
	//userWallet := db.UW{Uid: user.ID, Balance: regConfigVal, FreezeBalance: 0}
	//if err := tx.Create(&userWallet).Error; err != nil {
	//	tx.Rollback()
	//	c.JSON(http.StatusOK, laya.GetMessage("Err", 0, language, map[string]interface{}{}))
	//	return
	//}
	//
	//// 4.邀请码判断
	//if rd.InviteCode != "" {
	//	agentUser := db.User{InviteCode: rd.InviteCode}
	//	if rdb.Dao.Where(agentUser).Find(&agentUser).RecordNotFound() {
	//		tx.Rollback()
	//		c.JSON(http.StatusOK, laya.GetMessage("InviteCodeErr", 40010, language, map[string]interface{}{}))
	//		return
	//	} else {
	//		// 5.邀请关系绑定
	//		userInvite := db.UI{Uid: user.ID, AgentUid: agentUser.ID, CreateTime: time.Now()}
	//		if err := tx.Create(&userInvite).Error; err != nil {
	//			tx.Rollback()
	//			c.JSON(http.StatusOK, laya.GetMessage("Err", 0, language, map[string]interface{}{}))
	//			return
	//		}
	//	}
	//}
	//
	//// 6.发送系统消息
	//notice := db.Notice{Uid: user.ID, Title: "RegisteredSuccess", Content: "CongratulationsBecoming", CreateTime: time.Now(), Status: 2}
	//if err := tx.Create(&notice).Error; err != nil {
	//	tx.Rollback()
	//	c.JSON(http.StatusOK, laya.GetMessage("Err", 0, language, map[string]interface{}{}))
	//	return
	//}
	//
	//// 7.发送注册有奖消息
	//if regConfigVal > 0 {
	//	notice := db.Notice{Uid: user.ID, Title: "AwardRegistration", Content: "CongratulationsWinning|" + ac.FormatMoney(regConfigVal), CreateTime: time.Now(), Status: 2}
	//	if err := tx.Create(&notice).Error; err != nil {
	//		tx.Rollback()
	//		c.JSON(http.StatusOK, laya.GetMessage("Err", 0, language, map[string]interface{}{}))
	//		return
	//	}
	//}

	//tx.Commit()
	//// 7.返回数据
	//c.Set("$.Register.response", nil)
}

//
//var manager *captchas.Manager
//
//func (ctrl *controller) getManager() *captchas.Manager {
//	if manager == nil {
//		store := redisstore.New(rdb.Dao) // rdb store
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
//func (ctrl *controller) Captcha(c *gin.Context) {
//	//获取语言类型
//	lang := c.GetHeader("Accept-Language")
//	language := laya.GetLang(lang)
//	captcha, err := getManager().Generate()
//	if err != nil {
//		c.JSON(http.StatusOK, laya.GetMessage("Err", 0, language, map[string]interface{}{}))
//		return
//	}
//	data := map[string]interface{}{"ID": captcha.ID(), "Img": captcha.EncodeToString()}
//	c.JSON(http.StatusOK, laya.GetMessage("Success", 1, language, data))
//}
//
//// 短信验证码
//func (ctrl *controller) Phone(c *gin.Context) {
//	//获取语言类型
//	lang := c.GetHeader("Accept-Language")
//	language := laya.GetLang(lang)
//	var p validate.P
//	if err := c.ShouldBind(&p); err != nil {
//		c.JSON(http.StatusOK, laya.GetMessage("ParamErr", 40004, language, map[string]interface{}{}))
//		return
//	}
//
//	if p.Type == 1 { // 注册
//
//		user := db.User{Phone: p.Phone, Zone: p.Zone}
//		if !rdb.Dao.Where(user).Find(&user).RecordNotFound() {
//			c.JSON(http.StatusOK, laya.GetMessage("PhoneExist", 40006, language, map[string]interface{}{}))
//			return
//		}
//	} else if p.Type == 2 { // 忘记密码
//		user := db.User{Phone: p.Phone, Zone: p.Zone}
//		if rdb.Dao.Where(user).Find(&user).RecordNotFound() {
//			c.JSON(http.StatusOK, laya.GetMessage("NoUser", 40001, language, map[string]interface{}{}))
//			return
//		}
//	}
//
//	//图形验证码
//	//p.Value = strings.ToUpper(p.Value)
//	//if err := getManager().Verify(p.ID, p.Value, true); err != nil {
//	//	c.JSON(http.StatusOK, laya.GetMessage("ImgCodeErr", 40055, language, map[string]interface{}{}))
//	//	return
//	//}
//
//	/* // 1. 策略手机号正确性判断
//	   iorgo, _ := regexp.MatchString(`^(1[3|5|7|8|9][0-9]\d{4,8})$`, p.Phone)
//	   if !iorgo {
//	       c.JSON(http.StatusOK, laya.GetMessage("PhoneErr",40011,language,map[string]interface{}{}))
//	       return
//	   }*/
//
//	// 2. 策略同个手机号60秒后才能发送
//	num1, _ := rdb.Dao.Exists("phone:time:" + p.Zone + p.Phone).Result()
//	if num1 != 0 {
//		c.JSON(http.StatusOK, laya.GetMessage("MsgErr", 40012, language, map[string]interface{}{}))
//		return
//	}
//	if err := rdb.Dao.Set("phone:time:"+p.Zone+p.Phone, "1", 60*time.Second).Err(); err != nil {
//		c.JSON(http.StatusOK, laya.GetMessage("Err", 0, language, map[string]interface{}{}))
//		return
//	}
//
//	// // 3. 策略一个手机号24小时只能发送5条
//	// num2, _ := rdb.Dao.Get("phone:times:" + p.Phone).Result()
//	// intNum2, _ := strconv.ParseInt(num2, 10, 64)
//	// if intNum2 >= 5 {
//	//	c.JSON(http.StatusOK, laya.GetMessage("MsgFail",40013,language,map[string]interface{}{}))
//	//	return
//	// }
//	// intNum2 += 1
//	// if err := rdb.Dao.Set("phone:times:"+p.Phone, intNum2, 24*time.Hour).Err(); err != nil {
//	//	 c.JSON(http.StatusOK, laya.GetMessage("Err",0,language,map[string]interface{}{}))
//	//	return
//	// }
//	//
//	// // 4. 策略同一个ip24小时只能发送10条
//	// num3, _ := rdb.Dao.Get("phone:ip:" + laya.RemoteIp(c.Request)).Result()
//	// intNum3, _ := strconv.ParseInt(num3, 10, 64)
//	// if intNum3 >= 10 {
//	//	c.JSON(http.StatusOK, laya.GetMessage("MsgIpFail",40014,language,map[string]interface{}{}))
//	//	return
//	// }
//	// intNum3 += 1
//	// if err := rdb.Dao.Set("phone:ip:"+laya.RemoteIp(c.Request), intNum3, 24*time.Hour).Err(); err != nil {
//	//	 c.JSON(http.StatusOK, laya.GetMessage("Err",0,language,map[string]interface{}{}))
//	//	return
//	// }
//
//	// 6. 写入redis，300秒过期
//	code := laya.GenValidateCode(6)
//
//	if err := rdb.Dao.Set("phone:code:"+p.Zone+p.Phone, code, 5*time.Minute).Err(); err != nil {
//		c.JSON(http.StatusOK, laya.GetMessage("Err", 0, language, map[string]interface{}{}))
//		return
//	}
//
//	// 7. 执行发送信息
//	sendResult := SendPhoneCode1(code, p.Phone, p.Zone, p.Type)
//	if sendResult.Code != 0 {
//		c.JSON(http.StatusOK, laya.GetMessage("PhoneCodeFail", 40025, language, sendResult))
//		return
//	}
//
//	// 8. 成功返回
//	c.JSON(http.StatusOK, laya.GetMessage("Success", 1, language, sendResult))
//}
//
//// 修改用户密码
//func (ctrl *controller) EditUserPwd(c *gin.Context) {
//	//获取语言类型
//	lang := c.GetHeader("Accept-Language")
//	language := laya.GetLang(lang)
//	var phone = c.PostForm("Phone")
//	var Code = c.PostForm("Code")
//	var Zone = c.PostForm("Zone")
//	var NewPassword = c.PostForm("NewPassword")
//	var ReNewPassword = c.PostForm("ReNewPassword")
//
//	// 5.短信验证码验证
//	code, err := rdb.Dao.Get("phone:code:" + Zone + phone).Result()
//	if err != nil {
//		c.JSON(http.StatusOK, laya.GetMessage("CodeErr", 40008, language, map[string]interface{}{}))
//		return
//	}
//	if code != Code {
//		c.JSON(http.StatusOK, laya.GetMessage("PhoneCodeErr", 40009, language, map[string]interface{}{}))
//		return
//	}
//
//	// 判断重复密码是否一致
//	if NewPassword != ReNewPassword {
//		c.JSON(http.StatusOK, laya.GetMessage("PWDFail", 40007, language, map[string]interface{}{}))
//		return
//	}
//
//	// 修改密码
//	rdb.Dao.Model(&db.User{}).Where("phone = ?", phone).Where("zone = ?", Zone).Update("password", NewPassword)
//
//	// 返回信息=
//	c.JSON(http.StatusOK, laya.GetMessage("Success", 1, language, map[string]interface{}{}))
//}
//
// 获取用户信息
func (ctrl *controller) GetUserInfo(c *gin.Context) {
	//uid := c.GetInt64("uid")
	//user := db.User{ID: uid}
	//if result := db.Dao.Model(&db.User{}).Where(&user).Find(&user);
	//	errors.Is(result.Error, gorm.ErrRecordNotFound) {
	//	c.Set("$.err", response.NotFoundUser)
	//	return
	//}
	//c.Set("$.response", user)
}

//
//// 修改用户信息
//func (ctrl *controller) EditUser(c *gin.Context) {
//	//获取语言类型
//	lang := c.GetHeader("Accept-Language")
//	language := laya.GetLang(lang)
//	uid := c.GetInt64("uid")
//	var etype = c.PostForm("etype") // 修改类型枚举
//	var nickname = c.PostForm("nickname")
//	var avatar = c.PostForm("avatar")
//	var qq = c.PostForm("qq")
//	var phone = c.PostForm("phone")
//	var zone = c.PostForm("zone")
//
//	user := db.User{}
//	user.ID = uid
//
//	if rdb.Dao.Where(user).Find(&user).RecordNotFound() {
//		c.JSON(http.StatusOK, laya.GetMessage("NoUser", 40001, language, map[string]interface{}{}))
//		return
//	}
//
//	switch etype {
//	// 修改昵称
//	case "1":
//		if nickname == "" {
//			c.JSON(http.StatusOK, laya.GetMessage("Err", 0, language, map[string]interface{}{}))
//			return
//		}
//		user.Nickname = nickname
//		break
//	// 修改头像
//	case "2":
//		if avatar == "" {
//			c.JSON(http.StatusOK, laya.GetMessage("Err", 0, language, map[string]interface{}{}))
//			return
//		}
//		user.Avatar = avatar
//		break
//	// 修改QQ
//	case "3":
//		if qq == "" {
//			c.JSON(http.StatusOK, laya.GetMessage("Err", 0, language, map[string]interface{}{}))
//			return
//		}
//		user.QQ = qq
//		break
//	// 修改手机号
//	case "4":
//		if phone == "" || zone == "" {
//			c.JSON(http.StatusOK, laya.GetMessage("Err", 0, language, map[string]interface{}{}))
//			return
//		}
//		user.Phone = phone
//		user.Zone = zone
//		break
//	}
//
//	if err := rdb.Dao.Save(&user).Error; err != nil {
//		c.JSON(http.StatusOK, laya.GetMessage("Err", 0, language, map[string]interface{}{}))
//		return
//	}
//
//	c.JSON(http.StatusOK, laya.GetMessage("Success", 1, language, map[string]interface{}{}))
//}
//
//// 修改密码
//func (ctrl *controller) EditPwd(c *gin.Context) {
//	//获取语言类型
//	lang := c.GetHeader("Accept-Language")
//	language := laya.GetLang(lang)
//	uid := c.GetInt64("uid")
//	var oldpassword = c.PostForm("oldpassword")
//	var newpassword = c.PostForm("newpassword")
//	var renewpassword = c.PostForm("renewpassword")
//	var Phone = c.PostForm("phone")
//	var Zone = c.PostForm("zone")
//	var PhoneCode = c.PostForm("PhoneCode")
//
//	if PhoneCode == "" {
//		c.JSON(http.StatusOK, laya.GetMessage("PhoneCodeErr", 40009, language, map[string]interface{}{}))
//		return
//	}
//	// 5.短信验证码验证
//	code, err := rdb.Dao.Get("phone:code:" + Zone + Phone).Result()
//	if err != nil {
//		c.JSON(http.StatusOK, laya.GetMessage("CodeErr", 40008, language, map[string]interface{}{}))
//		return
//	}
//	if code != PhoneCode {
//		c.JSON(http.StatusOK, laya.GetMessage("PhoneCodeErr", 40009, language, map[string]interface{}{}))
//		return
//	}
//
//	if len(oldpassword) != 32 || len(newpassword) != 32 || len(renewpassword) != 32 {
//		c.JSON(http.StatusOK, laya.GetMessage("ParamErr", 40004, language, map[string]interface{}{}))
//		return
//	}
//
//	// 判断重复密码是否一致
//	if newpassword != renewpassword {
//		c.JSON(http.StatusOK, laya.GetMessage("PWDFail", 40007, language, map[string]interface{}{}))
//		return
//	}
//
//	user := db.User{ID: uid, Password: oldpassword}
//
//	if rdb.Dao.Where(user).Find(&user).RecordNotFound() {
//		c.JSON(http.StatusOK, laya.GetMessage("PWDErr", 40007, language, map[string]interface{}{}))
//		return
//	}
//
//	// 修改密码
//	user.Password = newpassword
//	if err := rdb.Dao.Save(&user).Error; err != nil {
//		c.JSON(http.StatusOK, laya.GetMessage("Err", 0, language, map[string]interface{}{}))
//		return
//	}
//
//	// 返回信息
//	c.JSON(http.StatusOK, laya.GetMessage("Success", 1, language, map[string]interface{}{}))
//}
//
//// 签到
//func (ctrl *controller) SignIn(c *gin.Context) {
//	//获取语言类型
//	lang := c.GetHeader("Accept-Language")
//	language := laya.GetLang(lang)
//	uid := c.GetInt64("uid")
//	ymd := time.Now().Format("2006-01-02")
//	//千位分割符(金额)
//	ac := laya.Ac
//
//	// 获取锁开启事务
//	lockKey := "SignIn:lock:" + strconv.FormatInt(uid, 10)
//	lock, err := laya.GetLock(rdb.Dao, lockKey, 5*time.Second, 5*time.Second)
//	if err != nil {
//		c.JSON(http.StatusOK, laya.GetMessage("SysErr", 40024, language, map[string]interface{}{}))
//		return
//	}
//
//	// 1.判断是否已经签到了
//	usl := db.USL{Uid: uid, Ymd: ymd}
//	if !rdb.Dao.Model(db.USL{}).Where(usl).Find(&usl).RecordNotFound() {
//		laya.ReleaseLock(rdb.Dao, lockKey, lock)
//		c.JSON(http.StatusOK, laya.GetMessage("Signed", 40059, language, map[string]interface{}{}))
//		return
//	}
//	//response := laya.Response{Code: 1, Msg: map[string]interface{}{"Chinese":"成功","English":"success"}}
//	var Data interface{}
//	tx := rdb.Dao.Begin()
//	// 2. 赠送奖励
//	conf := db.Config{}
//	tx.Model(conf).Where("id = ?", 3).Find(&conf)
//	if conf.Status == 1 {
//		userWallet := db.UW{Uid: uid}
//		tx.Where(userWallet).Find(&userWallet)
//		amount, _ := strconv.ParseInt(conf.Val, 10, 64)
//		//userWallet.Balance += amount
//		usl.Amount = amount
//		walletLog := db.WalletLog{Uid: uid, Wtype: 19, Ttype: 1, Amount: amount, CreateTime: time.Now(), BeforeAmount: userWallet.Balance}
//		if err := tx.Create(&walletLog).Error; err != nil {
//			tx.Rollback()
//			laya.ReleaseLock(rdb.Dao, lockKey, lock)
//			c.JSON(http.StatusOK, laya.GetMessage("Err", 0, language, err))
//			return
//		}
//		if err := tx.Model(db.UW{}).Where(userWallet).UpdateColumn("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
//			tx.Rollback()
//			laya.ReleaseLock(rdb.Dao, lockKey, lock)
//			c.JSON(http.StatusOK, laya.GetMessage("Err", 0, language, err))
//			return
//		}
//
//		// 发送消息通知
//		notice := db.Notice{Uid: uid, Title: "SignedSuccess", Status: 2, CreateTime: time.Now()}
//		notice.Content = "SignedInSuccess|" + ac.FormatMoney(amount)
//		if err := tx.Create(&notice).Error; err != nil {
//			tx.Rollback()
//			laya.ReleaseLock(rdb.Dao, lockKey, lock)
//			c.JSON(http.StatusOK, laya.GetMessage("Err", 0, language, err))
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
//		laya.ReleaseLock(rdb.Dao, lockKey, lock)
//		c.JSON(http.StatusOK, laya.GetMessage("Err", 0, language, err))
//		return
//	}
//
//	tx.Commit()
//	laya.ReleaseLock(rdb.Dao, lockKey, lock)
//
//	c.JSON(http.StatusOK, laya.GetMessage("Success", 1, language, Data))
//}
