package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/util/log"
	"laya-go/ship"
	r "laya-go/ship/response"
	"laya-go/ship/utils"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

var Key = "f21a12a5679078397ab2b98bd4f7b284"

func SortParams(p url.Values) string {
	var signStr string
	var lst []string
	for k, _ := range p {
		lst = append(lst, k)
	}
	// ToLower sort
	sort.Slice(lst, func(i, j int) bool { return strings.ToLower(lst[i]) < strings.ToLower(lst[j]) })
	for _, v := range lst {
		if v != "Sign" {
			if p.Get(v) != "" {
				signStr += v + "=" + p.Get(v) + "&"
			}
		}
	}
	signStr += Key
	return signStr
}

func GetParams(c *gin.Context) url.Values {
	_, _ = c.MultipartForm()
	var params = c.Request.PostForm
	return params
}

func Validate(sign string, params url.Values, c *gin.Context) {
	t := params.Get("T")
	intT, _ := strconv.ParseInt(t, 10, 64)
	uuid := params.Get("U")

	exist, _ := ship.Redis.SIsMember("user:uuid", uuid).Result()
	if exist {
		c.Set("$.RequestFrequentUuid.code", r.RequestFrequentUuid)
		return
	}
	ship.Redis.SAdd("user:uuid", uuid)
	log.Info(time.Now().UnixNano()/1e6 - intT)
	if time.Now().UnixNano()/1e6-intT > 3000 {
		c.Set("$.RequestFrequentTime.code", r.RequestFrequentTime)
		return
	}

	if sign != params.Get("Sign") {
		c.Set("$.RequestFrequentSign.code", r.RequestFrequentSign)
		return
	}
}

func RunSign(s string) string {
	return utils.MD5(s)
}

// 这里对一些方法进行屏蔽
// 文件上传不用接口签名和加密
func RunWithoutSign(c *gin.Context) bool {
	var route = c.Request.URL.Path
	if route == "/hall/files/upload" {
		return false
	}
	if route == "/hall/user/pay/notify" {
		return false
	}
	if strings.Index(route, "/hall/files") == 0 {
		return false
	}

	return true
}

func New(c *gin.Context) {
	if RunWithoutSign(c) {
		params := GetParams(c)
		str := SortParams(params)
		sign := RunSign(str)
		Validate(sign, params, c)
	}
}

func Sign() gin.HandlerFunc {
	return func(c *gin.Context) {
		New(c)
		c.Next()
	}
}
