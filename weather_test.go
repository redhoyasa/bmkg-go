package bmkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWeatherForecast(t *testing.T) {
	setup()
	ass := assert.New(t)

	res := testClient.GetWeatherForecast()

	ass.Equal(res, 0)
}
