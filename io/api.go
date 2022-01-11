package io

type WeatherResp struct {
	Status    string      `json:"status"`
	Count     string      `json:"count"`
	Info      string      `json:"info"`
	InfoCode  string      `json:"infocode"`
	Forecasts []*Forecast `json:"forecasts"`
}

type Forecast struct {
	City       string  `json:"city"`
	AdCode     string  `json:"adcode"`
	Province   string  `json:"province"`
	ReportTime string  `json:"reporttime"`
	Casts      []*Cast `json:"casts"`
}

type Cast struct {
	Date         string `json:"date"`
	Week         string `json:"week"`
	DayWeather   string `json:"dayweather"`
	NightWeather string `json:"nightweather"`
	DayTemp      string `json:"daytemp"`
	NightTemp    string `json:"nighttemp"`
	DayWind      string `json:"daywind"`
	NightWind    string `json:"nightwind"`
	DayPower     string `json:"daypower"`
	NightPower   string `json:"nightpower"`
}
