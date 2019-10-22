package bmkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	mockBaseURL  = "http://data.bmkg.go.id"
	mockEndpoint = "/datamkg/MEWS/DigitalForecast/DigitalForecast-DIYogyakarta.xml"
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
	testClient.On("GetXMLBytes", mockBaseURL+mockEndpoint).Return(string(mock), nil)

	res, err := bmkgService.GetWeatherForecast("DIYogyakarta")
	if err != nil {
		t.Errorf("unexpected error in GetWeatherForecast: %v", err)
	}

	ass.Equal(res.Forecast.Area[0].Name, "Bantul")
	ass.Equal(res.Forecast.Area[0].Parameter[0].Description, "Humidity")
	testClient.AssertExpectations(t)
}
