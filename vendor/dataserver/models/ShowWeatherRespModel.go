package models

import (
	"gopkg.in/mgo.v2/bson"
)

type AddressResp struct {
	Showapi_res_code  int
	Showapi_res_error string
	Showapi_res_body  ShowWeatherModel
}
type ShowWeatherModel struct {
	BaseDBmodel
	CityInfo CityInfo
	F1       ForeCast
	F2       ForeCast
	F3       ForeCast
	F4       ForeCast
	F5       ForeCast
	F6       ForeCast
	F7       ForeCast
	Now      Nowreport
	Ret_code int
	Time     string
}
type ForeCast struct {
	Day                   string
	Day_air_temperature   string
	Day_weather           string
	Day_weather_pic       string
	Day_weather_direction string
	Day_wind_power        string
	Index                 Index
	Night_air_temperature string
	Night_weather         string
	Night_weather_pic     string
	Night_wind_direction  string
	Night_wind_power      string
	Sun_begin_end         string
	Weekday               int
}
type Nowreport struct {
	Aqi              interface{}
	AqiDetail        AqiDetail
	Sd               string
	Temperature      string
	Temperature_time string
	Weather          string
	Weather_pic      string
	Wind_direction   string
	Wind_power       string
}
type Index struct {
	Beauty   Indexcontent
	Clothes  Indexcontent
	Cold     Indexcontent
	Comfort  Indexcontent
	Glass    Indexcontent
	Sports   Indexcontent
	Travel   Indexcontent
	Uv       Indexcontent
	Wash_car Indexcontent
}
type Indexcontent struct {
	Desc  string
	Title string
}
type AqiDetail struct {
	Aqi               int
	Area              string
	Co                float32
	No2               int
	O3                int
	O3_8h             int
	Pm10              int
	Pm2_5             int
	Primary_pollutant string
	Quality           string
	So2               int
}
type CityInfo struct {
	C1        string //地区id  ex 101280101   `bson:"_id"`
	C10       string
	C11       string
	C12       string //邮编 ex  510000
	C15       string
	C16       string
	C17       string
	C2        string //拼音  ex guangzhou
	C3        string //中文  ex 广州
	C4        string
	C5        string
	C6        string  //省拼音 ex guangdong
	C7        string  //省中文 ex 广东
	C8        string  //国家英文 ex china
	C9        string  //国家中文 ex 中国
	Latitude  float32 //纬度 23.108
	Longitude float32 //经度 113.265
}

func (this *ShowWeatherModel) Tablename() string {
	return "addressResBody"
}

func (this *ShowWeatherModel) init() (err error) {
	err = this.BaseDBmodel.init()
	if err != nil {
		return
	}
	this.c = this.db.C(this.Tablename())
	return
}

func GetOneData(cityid string) (content *ShowWeatherModel, err error) {
	content = &ShowWeatherModel{}
	err = content.init()
	if err != nil {
		return
	}
	defer content.session.Close()
	err = content.c.FindId(cityid).One(&content)
	return
}

func UpsertData(content *ShowWeatherModel) (err error) {
	err = content.init()
	if err != nil {
		return
	}
	defer content.session.Close()
	_, err = content.c.Upsert(bson.M{"_id": content.CityInfo.C1}, &content)
	return err
}
