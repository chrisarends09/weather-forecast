package model

import "time"

type Weather struct {
	StartTime             *time.Time `json:"start_time"`
	EndTime               *time.Time `json:"end_time"`
	Country               string     `json:"country"`
	City                  string     `json:"city"`
	Timezone              string     `json:"timezone"`
	TimezoneOffsetSeconds int64      `json:"timezone_offset_seconds"`
	Condition             string     `json:"condition"`
	Icon                  string     `json:"icon"`
	AvgTempF              float64    `json:"temp_f"`
	MinTempF              float64    `json:"min_temp_f"`
	MaxTempF              float64    `json:"max_temp_f"`
	AvgWindKph            float64    `json:"avg_wind_kph"`
	MinWindKph            float64    `json:"min_wind_kph"`
	MaxWindKph            float64    `json:"max_wind_kph"`
	HourlyWeathers        []Weather  `json:"hourly_weathers"`
}
