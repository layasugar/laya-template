package utils

import (
	"bytes"
	"container/list"
	"crypto/md5"
	crand "crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/oschwald/geoip2-golang"
	uuid "github.com/satori/go.uuid"
	"laya-go/ship"
	"log"
	"math/big"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// md5
func MD5(s string) string {
	m := md5.Sum([]byte(s))
	return hex.EncodeToString(m[:])
}

// 获取随机字符串
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, b[r.Intn(len(b))])
	}
	return string(result)
}

// 获取token
func GetToken() string {
	s := GetRandomString(16)
	return MD5(s)
}

// ip转int
func StringIpToInt(ipstring string) int {
	ipSegs := strings.Split(ipstring, ".")
	var ipInt int = 0
	var pos uint = 24
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	return ipInt
}

// ip int 转 string
func IpIntToString(ipInt int) string {
	ipSegs := make([]string, 4)
	var len int = len(ipSegs)
	buffer := bytes.NewBufferString("")
	for i := 0; i < len; i++ {
		tempInt := ipInt & 0xFF
		ipSegs[len-i-1] = strconv.Itoa(tempInt)
		ipInt = ipInt >> 8
	}
	for i := 0; i < len; i++ {
		buffer.WriteString(ipSegs[i])
		if i < len-1 {
			buffer.WriteString(".")
		}
	}
	return buffer.String()
}

// 获取6位随机字符串
func GetRandomString6(n uint64) []byte {
	baseStr := "0123456789ABCDEFGHJKLMNPQRSTUVWXYZ"
	base := []byte(baseStr)
	quotient := n
	mod := uint64(0)
	l := list.New()
	for quotient != 0 {
		mod = quotient % 34
		quotient = quotient / 34
		l.PushFront(base[int(mod)])
	}
	listLen := l.Len()
	if listLen >= 6 {
		res := make([]byte, 0, listLen)
		for i := l.Front(); i != nil; i = i.Next() {
			res = append(res, i.Value.(byte))
		}
		return res
	} else {
		res := make([]byte, 0, 6)
		for i := 0; i < 6; i++ {
			if i < 6-listLen {
				res = append(res, base[0])
			} else {
				res = append(res, l.Front().Value.(byte))
				l.Remove(l.Front())
			}

		}
		return res
	}

}

// 生成6位随机验证码
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

// RemoteIp 返回远程客户端的 IP，如 192.168.1.1
func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get(ship.XRealIP); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get(ship.XForwardedFor); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

// 时间格式化
type JsonTime time.Time

func (t JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

// 生成订单号
func CreateOrder() int64 {
	return int64(rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

func GetAddressByIP(ipA string) string {
	db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP(ipA)
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	var province string
	if len(record.Subdivisions) > 0 {
		province = record.Subdivisions[0].Names["zh-CN"]
	}

	return record.Country.Names["zh-CN"] + "-" + province + "-" + record.City.Names["zh-CN"]
}

func Abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

type ReData struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

//签名验证 uuid+key  签名验证
func Signature(uid string, key string) string {
	//拼接字符串
	var build strings.Builder
	build.WriteString(uid)
	build.WriteString(key)
	res := build.String()
	//进行md5加密
	return MD5(res)
}

//生成随机字符串
func RandSeqs(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetUUID() uuid.UUID {
	// 创建
	u1 := uuid.NewV4()
	return u1
}

//获取随机字符串
func RandString(width int) string {
	shipStr := "0123456789abcdefghijklmnopqrstuvwxyz"
	randStr := ""
	for i := 0; i < width; i++ {
		n, _ := crand.Int(crand.Reader, big.NewInt(35))
		index := n.Int64()
		randStr += string(shipStr[index])
	}
	return randStr
}