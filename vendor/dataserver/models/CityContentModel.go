package models

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"sync"
)

//cityids
type citycontent struct {
	Areas []Area
}

type Area struct {
	Area string
	City []City
}

type City struct {
	Cityname string
	Cityid   string
}

var cityidcontent *citycontent //内存驻留cityid列表
var mutex sync.Mutex           //互斥锁

func CheckCitycontentInstance() (err error) {
	mutex.Lock()
	defer mutex.Unlock()
	if cityidcontent == nil {
		err = confInit()
		if err != nil {
			return
		}
	}
	return
}

//初始化cityid列表
func confInit() (err error) {
	//fmt.Printf("conf init\n")
	buf, err := ioutil.ReadFile("./conf/cityid.json") //读取文件
	if err != nil {
		return
	}
	cityidcontent = &citycontent{}
	err = json.Unmarshal(buf, cityidcontent) //序列化文件为对象
	if err != nil {
		return
	}
	//fmt.Printf("Area:%s\n",cityread.Cityids[0].City[0].Cityid)
	return
}

func GetOneCityid(cityname string) (result City, err error) {
	if cityname == "" {
		err = errors.New("cityname is empty")
		return
	}
	err = CheckCitycontentInstance()
	if err != nil {
		return
	}
	for _, Area := range cityidcontent.Areas {
		for _, City := range Area.City {
			if cityname == City.Cityname {
				result = City
				return
			}
		}
	}
	err = errors.New("cityname not found")
	return
}

func GetCityContent() (result *citycontent, err error) {
	err = CheckCitycontentInstance()
	if err != nil {
		return
	}
	return cityidcontent, nil

}
