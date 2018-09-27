package go_base_libs

import (
	"encoding/json"
	"net/http"
	"crypto/tls"
	"bytes"
	"io/ioutil"
	log "github.com/cihub/seelog"
)

func SendDingDingWebHook(sendDataStruct [] byte, url string)  {

	postJson,err := json.Marshal(sendDataStruct)

	if err != nil{
		log.Error("转换结构体为json格式失败，请检查程序", err)
		return
	}
	//https
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//http 主体
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postJson))
	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Transport:tr}
	resp, herr := client.Do(req)
	if herr != nil {
		log.Error("给叮叮发送告警信息失败", herr)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Info("叮叮的返回信息为:", string(body))
}