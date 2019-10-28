package bmkg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	mockBaseURL         = "http://data.bmkg.go.id"
	mockWeatherEndpoint = "/datamkg/MEWS/DigitalForecast/DigitalForecast-%v.xml"
)

func TestGetWeatherForecast(t *testing.T) {
	setup()
	ass := assert.New(t)

	bmkgService := NewBmkg(nil)
	bmkgService.client = testClient

	// mock data file
	mock, err := getMockResponse("forecast.xml")
	if err != nil {
		t.Errorf("unexpected error in getMockResponse: %v", err)
	}
	testClient.On("GetXMLBytes", mockBaseURL+fmt.Sprintf(mockWeatherEndpoint, "DIYogyakarta")).Return(mock, nil)
	testClient.On("GetXMLBytes", mockBaseURL+fmt.Sprintf(mockWeatherEndpoint, "Dressrosa")).Return(nil, fmt.Errorf("data not found"))

	res, err := bmkgService.GetWeatherForecast("DIYogyakarta")
	if err != nil {
		t.Errorf("unexpected error in GetWeatherForecast: %v", err)
	}

	ass.Equal(res.Forecast.Area[0].Name, "Bantul")
	ass.Equal(res.Forecast.Area[0].Parameter[0].Description, "Humidity")

	res, err = bmkgService.GetWeatherForecast("Dressrosa")
	ass.Nil(res)
	ass.EqualError(err, "data not found")

	testClient.AssertExpectations(t)
}
