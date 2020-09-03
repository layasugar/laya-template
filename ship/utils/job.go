package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"laya-go/ship"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type JobMessage struct {
	ID int64 `json:"id"`
}

type JobMessageRule struct {
	Point     int64 `json:"point"`      // 加得比例
	Rule      int64 `json:"rule"`       // 总比列
	ProjectID int64 `json:"project_id"` // 项目ID
}

type ProjectJobMessage struct {
	ID      int64 `json:"id"`
	Uid     int64 `json:"uid"`
	OrderID int64 `json:"orderID"`
	End     int64 `json:"end"`
}

// delayTime 执行时间秒
func AddJob(JobName string, body JobMessage, delayTime int64) string {
	id, _ := JobPush(JobName, body, delayTime)
	return id
}

// delayTime 执行时间秒
func AddJobRule(JobName string, body JobMessageRule, delayTime int64) string {
	id, b := JobPush(JobName, body, delayTime)
	if b == true {
		return id
	} else {
		return ""
	}
}

// 移出任务
func RemoveJob(msgID string) bool {
	_, err := JobRemove(msgID)
	if err != nil {
		return false
	} else {
		return true
	}
}

// delayTime 执行时间秒
func ProjectAddJob(body ProjectJobMessage, delayTime int64) string {
	id, _ := JobPush("profit", body, delayTime)
	return id
}

type QueuePush struct {
	Topic string `json:"topic"`
	Id    string `json:"id"`
	Delay int64  `json:"delay"`
	TTR   int64  `json:"ttr"`
	Body  string `json:"body" `
}
type QueuePop struct {
	Topic string `json:"topic"`
}
type QueueRemove struct {
	Id string `json:"id"`
}

type ResData struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    ResultData `json:"data"`
}
type ResultData struct {
	Id   string `json:"id"`
	Body string `json:"body"`
}

// 任务push
func JobPush(topic string, body interface{}, delayTime int64) (string, bool) {
	jobMsg, _ := json.Marshal(body)
	msg := QueuePush{
		Topic: topic,
		Id:    fmt.Sprintf("%x", md5.Sum([]byte(strconv.FormatInt(time.Now().UnixNano(), 10)+RandSeqs(5)))),
		Delay: delayTime,
		TTR:   120,
		Body:  string(jobMsg),
	}
	msgJson, _ := json.Marshal(msg)
	urls := ship.DelayServer + "/push"
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urls, bytes.NewBuffer(msgJson))
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	var result ResData
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(res, result)
	if result.Code == 0 {
		fmt.Println("pushres:id:", msg.Id)
		return msg.Id, true
	} else {
		return msg.Id, false
	}
}

// 任务pop
func JobPop(topic string) (ResData, error) {
	var data ResData
	urls := ship.DelayServer + "/pop"
	resp, err := http.Post(urls, "application/json;charset=UTF-8", strings.NewReader("{\"topic\":\""+topic+"\"}"))
	if err != nil {
		return data, err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	// return
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(res, &data)
	return data, err
}

func JobFinish(id string) (string, error) {
	msg := QueueRemove{Id: id}
	msgJson, err := json.Marshal(msg)
	urls := ship.DelayServer + "/finish"
	client := &http.Client{}
	req, err := http.NewRequest("POST", urls, bytes.NewBuffer(msgJson))
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		return msg.Id, err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return msg.Id, err
	}
	var result ResData
	_ = json.Unmarshal(res, result)
	return msg.Id, err
}

// 移出任务
func JobRemove(id string) (string, error) {
	msg := QueueRemove{
		Id: id,
	}
	msgJson, err := json.Marshal(msg)
	urls := ship.DelayServer + "/delete"
	client := &http.Client{}
	req, err := http.NewRequest("POST", urls, bytes.NewBuffer(msgJson))
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		return msg.Id, err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return msg.Id, err
	}

	var result ResData
	_ = json.Unmarshal(res, result)
	return msg.Id, err
}
