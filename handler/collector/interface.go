package collector

import (
	"context"
	"github.com/chrisarends09/weather-forecast/model"
)

type WeatherService interface {
	Forecast(ctx context.Context, city string, region string, days int) ([]model.Weather, error)
}
