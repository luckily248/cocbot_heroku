package tests

import (
	_ "crypto/hmac"
	_ "crypto/sha1"
	"dataserver/models"
	_ "dataserver/routers"
	_ "encoding/base64"
	"encoding/json"
	"fmt"
	_ "github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"net/url"
	_ "path/filepath"
	_ "runtime"
	"testing"
	_ "time"
)

//func init() {
//	_, file, _, _ := runtime.Caller(1)
//	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
//	beego.TestBeegoInit(apppath)
//}
//const(
//	base64Table="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
//)

//// TestGet is a sample to run an endpoint test
//func TestGet(t *testing.T) {

//	r, _ := http.NewRequest("GET", requesturl, nil)
//	w := httptest.NewRecorder()
//	beego.BeeApp.Handlers.ServeHTTP(w, r)

//	beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())

//	Convey("Subject: Test Station Endpoint\n", t, func() {
//	        Convey("Status Code Should Be 200", func() {
//	                So(w.Code, ShouldEqual, 200)
//	        })
//	        Convey("The Result Should Not Be Empty", func() {
//	                So(w.Body.Len(), ShouldBeGreaterThan, 0)
//	        })
//	})
//}

//func TestCnApi(t *testing.T){
//	areaid:="101280101"   //广州
//	apitype:="forecast_f"   //常规
//	date:="201507101613" //客户端时间
//	appidfull:="e2169407f962168b" //完整appid 生成公钥用
//	appid:="e21694" //前6位 传输参数用
//	private_key:="a0889d_SmartWeatherAPI_d66d004" //密钥
//	homeurl:="http://open.weather.com.cn/data/" //固定url
//	public_key:=fmt.Sprintf("%s?areaid=%s&type=%s&date=%s&appid=%s",homeurl,areaid,apitype,date,appidfull) //除key以外为公钥

//	t.Logf("public_key:%s\n",public_key) //输出公钥

//	h:=hmac.New(sha1.New,[]byte(private_key)) //hmac-sha1加密
//	h.Write([]byte(public_key))
//	coder:=base64.NewEncoding(base64Table)  //base64加密
//	key:=coder.EncodeToString(h.Sum(nil))
//	keybyjava:="DOItq0SaBIM8zl1eexsyq6SuWhc%3D"

//	t.Logf("key:%s\n",key) //输出key

//	requesturl:=fmt.Sprintf("%s?areaid=%s&type=%s&date=%s&appid=%s&key=%s",homeurl,areaid,apitype,date,appid,keybyjava) //完整请求路径

//	t.Logf("requesturl:%s\n",requesturl) //输出路径

//	resp,err:=http.Get(requesturl)
//	if err!=nil{
//		t.Fatalf("httpGet err:%s",err.Error())
//	}
//	defer resp.Body.Close()
//	body,err:=ioutil.ReadAll(resp.Body)
//	if err!=nil{
//		t.Fatalf("bodyread err:%s",err.Error())
//	}
//	t.Logf("resp body:%s\n",string(body))
//}

//type AreaResp struct{
//	Showapi_res_code int
//	Showapi_res_error string
//	Showapi_res_body ResBody
//}

//type ResBody struct{
//	List []ResBodyAreaInfo
//	Ret_code int
//}

//type ResBodyAreaInfo struct{
//	Area string  //地区名  ex 广州
//	Areaid string  //地区id  ex 101280101
//	CityInfo CityInfo  //城市信息
//	Distric string //市  ex 广州
//	Prov string  //省  ex 广东
//}

//type CityInfo struct{
//	C1 string    //地区id  ex 101280101
//	C10 string
//	C11 string
//	C12 string   //邮编 ex  510000
//	C15 string
//	C16 string
//	C17 string
//	C2 string   //拼音  ex guangzhou
//	C3 string   //中文  ex 广州
//	C4 string
//	C5 string
//	C6 string   //省拼音 ex guangdong
//	C7 string	//省中文 ex 广东
//	C8 string	//国家英文 ex china
//	C9 string	//国家中文 ex 中国
//	Latitude float32  //纬度 23.108
//	Longitude float32 //经度 113.265
//}

//type AddressResp struct{
//	Showapi_res_code int
//	Showapi_res_error string
//	Showapi_res_body AddressResBody
//}
//type AddressResBody struct{
//	CityInfo CityInfo
//	F1 ForeCast
//	F2 ForeCast
//	F3 ForeCast
//	F4 ForeCast
//	F5 ForeCast
//	F6 ForeCast
//	F7 ForeCast
//	Now Nowreport
//	Ret_code int
//	Time string
//}
//type ForeCast struct{
//	Day string
//	Day_air_temperature string
//	Day_weather string
//	Day_weather_pic string
//	Day_weather_direction string
//	Day_wind_power string
//	Index Index
//	Night_air_temperature string
//	Night_weather string
//	Night_weather_pic string
//	Night_wind_direction string
//	Night_wind_power string
//	Sun_begin_end string
//	Weekday int
//}
//type Nowreport struct{
//	Aqi int32
//	AqiDetail AqiDetail
//	Sd string
//	Temperature string
//	Temperature_time string
//	Weather string
//	Weather_pic string
//	Wind_direction string
//	Wind_power string
//}
//type Index struct{
//	Beauty Indexcontent
//	Clothes Indexcontent
//	Cold Indexcontent
//	Comfort Indexcontent
//	Glass Indexcontent
//	Sports Indexcontent
//	Travel Indexcontent
//	Uv Indexcontent
//	Wash_car Indexcontent
//}
//type Indexcontent struct{
//	Desc string
//	Title string
//}
//type AqiDetail struct{
//	Aqi int
//	Area string
//	Co float32
//	No2 int
//	O3 int
//	O3_8h int
//	Pm10 int
//	Pm2_5 int
//	Primary_pollutant string
//	Quality string
//	So2 int
//}
var apikey = "c480facf3199797f836e2122a60e3274"
var servicesUrl = "http://apis.baidu.com/showapi_open_bus/weather_showapi"
var areaidService = "/areaid"

func getAreaId(area string) (resultareaid string, err error) {
	values := url.Values{}
	values.Set("area", area)
	resultareaid = ""
	err = nil

	areaidtestUrl := fmt.Sprintf("%s%s?%s", servicesUrl, areaidService, values.Encode())
	//t.Logf("areaidtest:%s\n",areaidtestUrl) //输出请求地区id路径

	req, err := http.NewRequest("GET", areaidtestUrl, nil)
	if err != nil {
		return
	}
	req.Header.Set("apikey", apikey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	//t.Logf("resp body:%s\n",string(body))

	areaResp := &models.AddressResp{}
	err = json.Unmarshal(body, areaResp)
	if err != nil {
		return
	}

	resultareaid = areaResp.Showapi_res_body.CityInfo.C1
	//t.Logf("areaid:%s\n",resultareaid)//输出地区码
	return
}

//测试易源api
func TestApi(t *testing.T) {

	//testarea := "安阳"
	//testareaid, err := getAreaId(testarea)
	//if err != nil {
	//	t.Fatalf("%s\n", err.Error())
	//}
	testareaid := "101010100"
	addressservice := "/address"
	v2 := url.Values{}
	v2.Set("areaid", testareaid)
	v2.Set("needMoreDay", "1")
	v2.Set("needIndex", "1")

	addresstestUrl := fmt.Sprintf("%s%s?%s", servicesUrl, addressservice, v2.Encode())
	t.Logf("areaidtest:%s\n", addresstestUrl) //输出请求天气预报路径

	req, err := http.NewRequest("GET", addresstestUrl, nil)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}
	req.Header.Set("apikey", apikey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("bodyread err:%s", err.Error())
	}
	//t.Logf("resp body:%s\n",string(body))

	addressResp := &models.AddressResp{}
	err = json.Unmarshal(body, addressResp)
	if err != nil {
		t.Fatalf("%s\n", err.Error())
		return
	}
	t.Logf("addressResp:%s\n", addressResp.Showapi_res_body.Now.AqiDetail.Area)
	t.Logf("addressResp index:%s\n", addressResp.Showapi_res_body.F1.Index.Glass.Desc)
}

//测试upsert
func TestUpsert(t *testing.T) {
	body := &models.ShowWeatherModel{}
	body.CityInfo.C1 = string("101280101")
	body.Now.Temperature = string("37%")
	t.Logf("content _id:%s", body.CityInfo.C1)
	err := models.UpsertData(body)
	if err != nil {
		t.Fatalf("data save error:%s", err.Error())
	}
	result, err := models.GetOneData(string("101280101"))
	if err != nil {
		t.Fatalf("data get error:%s", err.Error())
	}
	t.Logf("content get _id:%s", result.Now.Temperature)
}
