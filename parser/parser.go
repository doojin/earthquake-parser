// Package parser provides functionality of getting data from REST service of usgs.gov
package parser

import (
	"github.com/doojin/earthquake-parser/earthquake"
	"encoding/xml"
	"strings"
	"strconv"
	"net/http"
	"io/ioutil"
)

// SendRequest sends HTTP requests to the REST API of usgs.gov and
// returns parsed collection of earthquakes
func SendRequest(startTime string, endTime string, minMagnitude int, maxMagnitude int) (collection earthquake.EarthquakeCollection) {
	url := buildUrl(startTime, endTime, minMagnitude, maxMagnitude)
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	responseBody, _ := ioutil.ReadAll(response.Body)
	return parseXML(string(responseBody))
}

func parseXML(xmlString string) (collection earthquake.EarthquakeCollection) {
	xml.Unmarshal([]byte(xmlString), &collection)
	return collection
}

func buildUrl(startTime string, endTime string, minMagnitude int, maxMagnitude int) string {
	url := "http://earthquake.usgs.gov/fdsnws/event/1/query?format=xml&starttime={startTime}&endtime={endTime}&minmagnitude={minMagnitude}&maxmagnitude={maxMagnitude}"
	url = strings.Replace(url, "{startTime}", startTime, 1)
	url = strings.Replace(url, "{endTime}", endTime, 1)
	url = strings.Replace(url, "{minMagnitude}", strconv.Itoa(minMagnitude), 1)
	url = strings.Replace(url, "{maxMagnitude}", strconv.Itoa(maxMagnitude), 1)
	return url
}