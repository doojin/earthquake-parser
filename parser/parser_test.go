package parser

import (
	"testing"
	"github.com/doojin/earthquake-parser/earthquake"
	"github.com/stretchr/testify/assert"
)

func Test_ParseXML_ShouldParseXMLCorrectly(t *testing.T) {
	xmlString := "<q><eventParameters><event><origin><time><value>2014-01-01T23:51:36.020Z</value></time><longitude><value>-116.7776667</value></longitude><latitude><value>33.6633333</value></latitude></origin><magnitude><mag><value>1.29</value></mag></magnitude></event><event><origin><time><value>2014-01-01T23:51:36.020Z</value></time><longitude><value>-2.2</value></longitude><latitude><value>222.222</value></latitude></origin><magnitude><mag><value>2.22</value></mag></magnitude></event></eventParameters></q>"

	result := parseXML(xmlString)

	expectedResult := earthquake.EarthquakeCollection{
		Earthquakes: []earthquake.Earthquake{
			earthquake.Earthquake{
				Time: "2014-01-01T23:51:36.020Z",
				Longitude: -116.7776667,
				Latitude: 33.6633333,
				Magnitude: 1.29,
			},
			earthquake.Earthquake{
				Time: "2014-01-01T23:51:36.020Z",
				Longitude: -2.2,
				Latitude: 222.222,
				Magnitude: 2.22,
			},
		},
	}
	assert.Equal(t, expectedResult, result)
}

func Test_buildUrl_ShouldBuildValidUrl(t *testing.T) {
	url := buildUrl("2014-03-21", "2014-03-22", 1, 10)

	assert.Equal(t,
		"http://earthquake.usgs.gov/fdsnws/event/1/query?format=xml&starttime=2014-03-21&endtime=2014-03-22&minmagnitude=1&maxmagnitude=10",
		url,
	)
}
