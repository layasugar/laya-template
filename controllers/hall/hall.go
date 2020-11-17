package hall

//// 客户端心跳做在线处理
//func Ping(c *gin.Context) {
//	//获取语言类型
//	lang := c.GetHeader("Accept-Language")
//	language := ship.GetLang(lang)
//	uid := c.GetInt64("uid")
//	var timeUnix = float64(time.Now().Unix())
//	score := redis.Z{
//		Score:  timeUnix,
//		Member: uid,
//	}
//	ship.Redis.ZAdd("online_users", &score)
//
//	// 查询用户信息，判断账号是否是正常状态
//	userInfo := GetUserInfoByID(uid)
//
//	if userInfo.ID == 0 || userInfo.Status == 1 {
//		c.JSON(http.StatusOK, common.GetMessage("TokenErr",40003,language,map[string]interface{}{}))
//		return
//	}
//
//	c.JSON(http.StatusOK, common.GetMessage("Success",1,language,map[string]interface{}{}))
//}
//// 会员列表
//func GetUserList(c *gin.Context) {
//	//获取语言类型
//	lang := c.GetHeader("Accept-Language")
//	language := common.GetLang(lang)
//	var data validate.GetUserList
//	if err := c.ShouldBind(&data); err != nil {
//		c.JSON(http.StatusOK, common.GetMessage("ParamErr", 40004, language, map[string]interface{}{}))
//		return
//	}
//
//	now := time.Now().Unix()
//	tx := common.DB
//	if data.ID != 0 {
//		tx = tx.Where("tb_user.id = ?", data.ID)
//	} else {
//		if data.OnLine != 0 {
//			strNowMax := strconv.FormatInt(now, 10)
//			strNowMin := strconv.FormatInt(now-35, 10)
//			ZRangeBy := redis.ZRangeBy{
//				Min: strNowMin,
//				Max: strNowMax,
//			}
//			list, _ := common.Redis.ZRangeByScore("online_users", &ZRangeBy).Result()
//			if data.OnLine == 1 {
//				// 取出在线列表
//				tx = tx.Where(list)
//			} else {
//				tx = tx.Not(list)
//			}
//		}
//	}
//
//	var whereLevel string
//	if data.Level != "" {
//		whereLevel = " AND ul.level = " + data.Level
//	} else {
//		whereLevel = ""
//	}
//
//	if data.Nickname != "" {
//		tx = tx.Where("nickname = ?", data.Nickname)
//	}
//
//	if data.Relname != "" {
//		tx = tx.Where("real_name = ?", data.Relname)
//	}
//
//	if data.Phone != "" {
//		tx = tx.Where("phone = ?", data.Phone)
//	}
//
//	if data.IP != "" {
//		tx = tx.Where("last_address like ?", "%"+data.IP+"%")
//	}
//
//	//定义返回数据结构体
//	type reDate struct {
//		model.User
//		LastIp      string               // 最后登录ip
//		LastAddress string               // 最后登录地址
//		Level       model.UserLevel      `gorm:"foreignkey:uid;association_foreignkey:id"`
//		Online      int64                // 1在线，2没在线
//		Balance     model.UW             `gorm:"foreignkey:uid;association_foreignkey:id"`
//		Invite      model.UserInviteInfo `gorm:"foreignkey:uid;association_foreignkey:id"`
//	}
//	var datalist []reDate
//	var count int64
//	tx.Model(&model.User{}).Joins("JOIN tb_user_level ul ON ul.uid = tb_user.id" + whereLevel).Count(&count)
//	tx.Model(&model.User{}).Preload("Level").Preload("Balance").Preload("Invite.UserInfo").Joins("JOIN tb_user_level ul ON ul.uid = tb_user.id" + whereLevel).Offset((data.Page - 1) * data.Size).Limit(data.Size).Order("id DESC").Find(&datalist)
//
//	for k, v := range datalist {
//		tid := strconv.FormatInt(v.ID, 10)
//		score, _ := common.Redis.ZScore("online_users", tid).Result()
//		if (now - int64(score)) > 35 {
//			datalist[k].Online = 2
//		} else {
//			datalist[k].Online = 1
//		}
//	}
//
//	//返回数据
//	c.JSON(http.StatusOK, common.GetPageMessage("Success", 1, language, datalist, map[string]interface{}{}, common.PageRes{Size: data.Size, CurPage: data.Page, Total: count}))
//}