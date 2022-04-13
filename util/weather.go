package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type RspGetWeather struct {
	Cityid     string `json:"cityid"`
	City       string `json:"city"`
	CityEn     string `json:"cityEn"`
	Country    string `json:"country"`
	CountryEn  string `json:"countryEn"`
	UpdateTime string `json:"update_time"`
	Data       []struct {
		Day         string   `json:"day"`
		Date        string   `json:"date"`
		Week        string   `json:"week"`
		Wea         string   `json:"wea"`
		WeaImg      string   `json:"wea_img"`
		WeaDay      string   `json:"wea_day"`
		WeaDayImg   string   `json:"wea_day_img"`
		WeaNight    string   `json:"wea_night"`
		WeaNightImg string   `json:"wea_night_img"`
		Tem         string   `json:"tem"`
		Tem1        string   `json:"tem1"`
		Tem2        string   `json:"tem2"`
		Humidity    string   `json:"humidity"`
		Visibility  string   `json:"visibility"`
		Pressure    string   `json:"pressure"`
		Win         []string `json:"win"`
		WinSpeed    string   `json:"win_speed"`
		WinMeter    string   `json:"win_meter"`
		Sunrise     string   `json:"sunrise"`
		Sunset      string   `json:"sunset"`
		Air         string   `json:"air"`
		AirLevel    string   `json:"air_level"`
		AirTips     string   `json:"air_tips"`
		Alarm       struct {
			AlarmType    string `json:"alarm_type"`
			AlarmLevel   string `json:"alarm_level"`
			AlarmTitle   string `json:"alarm_title,omitempty"`
			AlarmContent string `json:"alarm_content"`
		} `json:"alarm"`
		Hours []struct {
			Hours    string `json:"hours"`
			Wea      string `json:"wea"`
			WeaImg   string `json:"wea_img"`
			Tem      string `json:"tem"`
			Win      string `json:"win"`
			WinSpeed string `json:"win_speed"`
		} `json:"hours"`
		Index []struct {
			Title string `json:"title"`
			Level string `json:"level"`
			Desc  string `json:"desc"`
		} `json:"index"`
	} `json:"data"`
	Aqi struct {
		UpdateTime string `json:"update_time"`
		Cityid     string `json:"cityid"`
		City       string `json:"city"`
		CityEn     string `json:"cityEn"`
		Country    string `json:"country"`
		CountryEn  string `json:"countryEn"`
		Air        string `json:"air"`
		AirLevel   string `json:"air_level"`
		AirTips    string `json:"air_tips"`
		Pm25       string `json:"pm25"`
		Pm25Desc   string `json:"pm25_desc"`
		Pm10       string `json:"pm10"`
		Pm10Desc   string `json:"pm10_desc"`
		O3         string `json:"o3"`
		O3Desc     string `json:"o3_desc"`
		No2        string `json:"no2"`
		No2Desc    string `json:"no2_desc"`
		So2        string `json:"so2"`
		So2Desc    string `json:"so2_desc"`
		Co         string `json:"co"`
		CoDesc     string `json:"co_desc"`
		Kouzhao    string `json:"kouzhao"`
		Yundong    string `json:"yundong"`
		Waichu     string `json:"waichu"`
		Kaichuang  string `json:"kaichuang"`
		Jinghuaqi  string `json:"jinghuaqi"`
	} `json:"aqi"`
}

const (
	TD  = `<td style="padding:6px 10px 6px 10px">`
	TD2 = `<td style="padding:2px 10px 2px 10px">`
)

// GetWeather 获取天气信息
func GetWeather() (string, error) {
	city, err := GetCity()
	if err != nil {
		return "", err
	}
	url := "https://www.tianqiapi.com/api?appid=24242169&appsecret=fUgcxMb2&version=v9&unescape=1&city=" + city
	println("查询天气：", url)
	defer func() {
		if err != nil {
			println("查询天气失败：", err)
		}
	}()
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	rsp := &RspGetWeather{}
	err = json.Unmarshal(data, rsp)
	if err != nil {
		return "", err
	}
	now := time.Now()
	table := `<table>`
	for k, item := range rsp.Data {
		if k == 0 {
			table += fmt.Sprintf("<tr>%v%v</td>%v%v</td>%v%v℃</td>%v%v</td></tr>\n", TD, now.Format("15:04"), TD, city, TD, item.Tem, TD, item.Wea)
			table += fmt.Sprintf("<tr>%v%v</td>%v%v</td>%v%v~%v℃</td>%v%v</td></tr>\n", TD, item.Date[5:], TD, item.Week, TD, item.Tem2, item.Tem1, TD, item.Wea)
			fmt.Printf("%v	%v	%v℃	%v\n", now.Format("15:04"), city, item.Tem, item.Wea)
			fmt.Printf("%v	%v	%v~%v℃	%v\n", item.Date[5:], item.Week, item.Tem2, item.Tem1, item.Wea)
			fmt.Printf("--------------------------------\n")
			for _, hour := range item.Hours {
				if ShowHour(hour.Hours) {
					table += fmt.Sprintf("<tr>%v%v</td>%v%v℃</td>%v%v</td>%v%v%v</td></tr>\n", TD2, hour.Hours, TD2, hour.Tem, TD2, hour.Wea, TD2, hour.Win, hour.WinSpeed)
					fmt.Printf("%v	%v℃	%v	%v%v\n", hour.Hours, hour.Tem, hour.Wea, hour.Win, hour.WinSpeed)
				}
			}
			fmt.Printf("--------------------------------\n")
		} else {
			table += fmt.Sprintf("<tr>%v%v</td>%v%v</td>%v%v~%v℃</td>%v%v</td></tr>\n", TD, item.Date[5:], TD, item.Week, TD, item.Tem2, item.Tem1, TD, item.Wea)
			fmt.Printf("%v	%v	%v~%v℃	%v\n", item.Date[5:], item.Week, item.Tem2, item.Tem1, item.Wea)
		}
	}
	table += `</table>`
	return table, nil
}

// ShowHour 展示当前时间以后的预报
func ShowHour(hour string) bool {
	comHour := 0
	nowHour := time.Now().Hour()
	list := strings.Split(hour[:2], "")
	num, _ := strconv.Atoi(list[0])
	if num == 0 {
		comHour, _ = strconv.Atoi(list[1])
	} else {
		comHour, _ = strconv.Atoi(hour[:2])
	}
	if comHour > nowHour {
		return true
	}
	return false
}
