package services

import (
	"fmt"
	"go.uber.org/zap"
	"time"
	"webhooks/io"

	"github.com/go-resty/resty/v2"
	. "github.com/virgoC0der/go-base/logging"
)

const (
	apiKey     = "a2c6cd5e8a35f8e9351da20aff6d1935"
	cityCode   = "440305"
	weatherApi = "https://restapi.amap.com/v3/weather/weatherInfo"
	hookApi    = "https://open.feishu.cn/open-apis/bot/v2/hook/3b1c12ee-00a5-4f59-9f95-97354a9d52e8"
	text       = "早上好，今天是%s星期%s，天气%s，最高温度%s，最低温度%s，%s风%s级"
)

var weekdayMap = map[string]string{
	"1": "一",
	"2": "二",
	"3": "三",
	"4": "四",
	"5": "五",
	"6": "六",
	"7": "日",
}

func WeatherHook() {
	t := time.NewTicker(time.Hour)
	retryT := time.NewTicker(5 * time.Second)
	client := resty.New()
	for {
		if time.Now().Hour() != 9 {
			<-t.C
			continue
		}

		param := map[string]string{
			"key":        apiKey,
			"city":       cityCode,
			"extensions": "all",
		}
		weatherResp := &io.WeatherResp{}
		resp, err := client.R().SetQueryParams(param).SetResult(weatherResp).Get(weatherApi)
		if err != nil {
			Logger.Warn("get weather forecast err", zap.Error(err))
			<-retryT.C
			continue
		}

		if resp.StatusCode() != 200 {
			Logger.Warn("get weather forecast err", zap.Error(err), zap.Int("code", resp.StatusCode()))
			<-retryT.C
			continue
		}
		weather := weatherResp.Forecasts[0].Casts[0]
		condition := weather.DayWeather
		if weather.DayWeather != weather.NightWeather {
			condition = fmt.Sprintf("%s转%s", weather.DayWeather, weather.NightWeather)
		}
		hookText := fmt.Sprintf(text, weather.Date, weekdayMap[weather.Week],
			condition, weather.DayTemp, weather.NightTemp, weather.DayWind, weather.DayPower)

		hookReq := &io.FeiShuReq{
			MsgType: "post",
			Content: &io.ContentInfo{
				Post: &io.PostInfo{
					ZhCn: &io.TextInfo{
						Title: "今天又是元气满满的一天",
						Content: []interface{}{
							[]*io.Text{
								{
									Tag:      "text",
									Text:     hookText,
									UnEscape: true,
								},
							},
						},
					},
				},
			},
		}
		var result io.FeiShuResp
		hookResp, err := client.R().SetHeader("Content-Type", "application/json").SetBody(hookReq).SetResult(&result).Post(hookApi)
		if err != nil {
			Logger.Warn("post hook err", zap.Error(err))
			<-retryT.C
			continue
		}

		Logger.Info("resp", zap.Any("resp", result))

		if hookResp.StatusCode() != 200 {
			Logger.Warn("post hook err", zap.Error(err), zap.Int("code", hookResp.StatusCode()))
			<-retryT.C
			continue
		}
		<-t.C
	}

}
