package global

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type AlarmMsg struct {
	Msgtype  string        `json:"msgtype"`
	Text     AlarmText     `json:"text"`
	Markdown AlarmMarkdown `json:"markdown"`
	At       AlarmAt       `json:"at"`
}

type AlarmText struct {
	Content string `json:"content"`
}

type AlarmMarkdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type AlarmAt struct {
	AtMobiless []string `json:"atMobiles"`
	IsAtAll    bool     `json:"isAtAll"`
}

type AlarmData struct {
	RobotKey    string                 //机器人key
	RobotHost   string                 //机器人域名
	Title       string                 //报警标题
	Description string                 //报警描述
	Content     map[string]interface{} //kv数据
}

func SendAlarm(d *AlarmData) error {
	var s string
	msg := bytes.Buffer{}
	s += fmt.Sprintf("## %s\r\n#### %s\r\n", d.Title, d.Description)
	for k, v := range d.Content {
		s += fmt.Sprintf("> %s：%v\r\n\r\n", k, v)
	}
	msg.WriteString(s)
	err := sendToDingTalk(d.RobotKey, d.RobotHost, msg.String())
	if err != nil {
		return err
	}
	return nil
}

func hmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return base64.URLEncoding.EncodeToString([]byte(h.Sum(nil)))
	//return base64.RawStdEncoding.EncodeToString([]byte(h.Sum(nil)))
	//return hex.EncodeToString(h.Sum(nil))
}

func sendToDingTalk(robotKey, robotHost, msg string) error {
	now := time.Now().Unix()
	now *= 1000
	nowStr := strconv.FormatInt(now, 10)
	sig := hmacSha256(nowStr+"\n"+robotKey, robotKey)
	sig = url.PathEscape(sig)
	sig = strings.Replace(sig, "-", "%2B", -1)
	sig = strings.Replace(sig, "_", "%2F", -1)
	requestUrl := robotHost + "&timestamp=" + nowStr + "&sign=" + sig
	textMsg := &AlarmMsg{}
	textMsg.Msgtype = "markdown"
	textMsg.Markdown = AlarmMarkdown{}
	textMsg.Markdown.Title = "payment post request mq Error"
	textMsg.Markdown.Text = msg
	postdata, _ := json.Marshal(textMsg)
	req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(postdata))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	responseData, _ := ioutil.ReadAll(resp.Body)
	log.Printf("钉钉通知结果：%s", string(responseData))
	return nil
}
