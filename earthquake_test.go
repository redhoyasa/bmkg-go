package bmkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	mockEarthquakeEndpoint = "/en_gempaterkini.xml"
)

func TestGetRecentEarthquakes(t *testing.T) {
	setup()
	ass := assert.New(t)

	bmkgService := NewBmkg(nil)
	bmkgService.client = testClient

	// mock data file
	mock, err := getMockResponse("earthquake.xml")
	if err != nil {
		t.Errorf("unexpected error in getMockResponse: %v", err)
	}
	testClient.On("GetXMLBytes", mockBaseURL+mockEarthquakeEndpoint).Return(mock, nil)

	res, err := bmkgService.GetRecentEarthquakes()
	if err != nil {
		t.Errorf("unexpected error in GetRecentEarthquakes: %v", err)
	}
	ass.Equal(res.Earthquakes[0].Time, "10:50:45 WIB")

	testClient.AssertExpectations(t)
}
