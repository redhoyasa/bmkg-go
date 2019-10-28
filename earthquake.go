package bmkg

import (
	"encoding/xml"
	"fmt"
)

// RecentEarthquakes ...
type RecentEarthquakes struct {
	Earthquakes []Earthquake `xml:"gempa"`
}

// Earthquake ...
type Earthquake struct {
	Date      string `xml:"Tanggal"`
	Time      string `xml:"Jam"`
	Latitude  string `xml:"Lintang"`
	Longitude string `xml:"Bujur"`
	Magnitude string `xml:"Magnitude"`
	Depth     string `xml:"Kedalaman"`
	Area      string `xml:"Wilayah"`
}

// GetRecentEarthquakes ...
func (b *Bmkg) GetRecentEarthquakes() (*RecentEarthquakes, error) {
	url := fmt.Sprintf(`%v/en_gempaterkini.xml`, b.config.BaseURL)

	xmlBytes, err := b.client.GetXMLBytes(url)
	if err != nil {
		return nil, err
	}

	var earthquakes *RecentEarthquakes
	xml.Unmarshal(xmlBytes, &earthquakes)
	return earthquakes, nil
}
