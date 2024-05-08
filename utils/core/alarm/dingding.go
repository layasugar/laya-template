package alarm

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type dingAlarmMsg struct {
	MsgType  string            `json:"msgtype"`
	Text     dingAlarmText     `json:"text"`
	Markdown dingAlarmMarkdown `json:"markdown"`
	At       dingAlarmAt       `json:"at"`
}

type dingAlarmText struct {
	Content string `json:"content"`
}

type dingAlarmMarkdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type dingAlarmAt struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

type DingContext struct {
	robotKey  string
	robotHost string
}

func (ctx *DingContext) Push(title string, content string, data map[string]interface{}) {
	var d = Data{
		Title:       title,
		Description: content,
		Content:     data,
	}
	go DingAlarm(ctx.robotKey, ctx.robotHost, d)
}

func DingAlarm(robotKey, robotHost string, ad Data) {
	var s string
	msg := bytes.Buffer{}
	s += fmt.Sprintf("## %s\r\n#### %s\r\n", ad.Title, ad.Description)
	for k, v := range ad.Content {
		s += fmt.Sprintf("> %s：%v\r\n\r\n", k, v)
	}
	msg.WriteString(s)
	err := sendToDingTalk(robotKey, robotHost, msg.String())
	if err != nil {
		log.Printf("dingding 推送失败, err: %s", err.Error())
	}
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
	textMsg := &dingAlarmMsg{}
	textMsg.MsgType = "markdown"
	textMsg.Markdown = dingAlarmMarkdown{}
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
	responseData, _ := io.ReadAll(resp.Body)
	log.Printf("钉钉通知请求结果：%s", string(responseData))
	return nil
}
