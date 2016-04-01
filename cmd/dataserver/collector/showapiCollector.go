package collector

import (
	. "dataserver/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	apikey         string = "c480facf3199797f836e2122a60e3274"                       //开发者key
	servicesUrl    string = "http://apis.baidu.com/showapi_open_bus/weather_showapi" //易源数据服务固定url（show）
	areaidService  string = "/areaid"                                                //获取名字对应cityid服务地址
	addressservice string = "/address"                                               //获取cityid对应天气信息的服务地址
)

//公开方法  运行采集器
func Run() {
	go loop() //开始循环
}

//循环采集

func loop() {
	wcollectortimer := time.NewTimer(time.Minute)
	for {
		select {
		case <-wcollectortimer.C: //1分钟后开始第一次采集
			fmt.Printf("collecting\n")
			collect()
			wcollectortimer.Reset(time.Hour * 12) //以后每12小时采集一次
		}
	}
}

//开始采集
func collect() {
	cityidcontent, err := GetCityContent()
	if err != nil {
		fmt.Printf("getcitycontent err:%s\n", err.Error())
		return
	}
	for i := 0; i < len(cityidcontent.Areas); i++ {
		fmt.Printf("begin area %s \n ", cityidcontent.Areas[i].Area)
		go collectArea(cityidcontent.Areas[i])
	}
	//fmt.Printf("alldone:%b\n",done)
}

//采集省级数据
func collectArea(area Area) {
	for i := 0; i < len(area.City); i++ {
		fmt.Printf("begin %s \n", area.City[i].Cityname)
		collectCity(area.City[i].Cityid)
		fmt.Printf("end %s \n", area.City[i].Cityname)
	}
}

//采集城市数据
func collectCity(cityid string) {
	v2 := url.Values{}
	v2.Set("areaid", cityid)
	v2.Set("needMoreDay", "1")
	v2.Set("needIndex", "1")

	addresstestUrl := fmt.Sprintf("%s%s?%s", servicesUrl, addressservice, v2.Encode())
	//fmt.Printf("areaidtest:%s\n",addresstestUrl) //输出请求天气预报路径

	req, err := http.NewRequest("GET", addresstestUrl, nil)
	if err != nil {
		fmt.Printf("request err:%s\n", err.Error())
		return
	}
	req.Header.Set("apikey", apikey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("client err:%s\n", err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("bodyread err:%s\n", err.Error())
		return
	}
	//fmt.Printf("resp body:%s\n",string(body))

	addressResp := &AddressResp{}
	err = json.Unmarshal(body, addressResp)
	if err != nil {
		fmt.Printf("json err %s:%s\n", cityid, err.Error())
		return
	}
	err = UpsertData(&addressResp.Showapi_res_body)
	if err != nil {
		fmt.Printf("upsert %s fail:%s\n", cityid, err.Error())
		return
	}
	fmt.Printf("upsert %s success\n", cityid)
}
