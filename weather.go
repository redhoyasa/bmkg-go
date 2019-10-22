package bmkg

import (
	"encoding/xml"
	"fmt"
)

// ForecastData ...
type ForecastData struct {
	Forecast Forecast `xml:"forecast"`
}

// Forecast ...
type Forecast struct {
	Area []Area `xml:"area"`
}

// Area ...
type Area struct {
	Name      string      `xml:"description,attr"`
	Parameter []Parameter `xml:"parameter"`
}

// Parameter ...
type Parameter struct {
	Description string        `xml:"description,attr"`
	Type        string        `xml:"type,attr"`
	Measurement []Measurement `xml:"timerange"`
}

// Measurement ...
type Measurement struct {
	Type     string `xml:"type,attr"`
	Datetime string `xml:"datetime,attr"`
	Unit     []Unit `xml:"value"`
}

// Unit ...
type Unit struct {
	Unit  string `xml:"unit,attr"`
	Value string `xml:",chardata"`
}

// GetWeatherForecast return weather forecast for certain province
func (b *Bmkg) GetWeatherForecast(province string) (*ForecastData, error) {
	endpoint := fmt.Sprintf("/datamkg/MEWS/DigitalForecast/DigitalForecast-%v.xml", province)
	url := fmt.Sprintf(`%v%v`, b.config.BaseURL, endpoint)

	xmlBytes, err := b.client.GetXMLBytes(url)
	if err != nil {
		return nil, err
	}

	var forecast *ForecastData
	xml.Unmarshal(xmlBytes, &forecast)
	return forecast, nil
}
