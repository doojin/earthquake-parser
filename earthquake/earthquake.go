// Package earthquake provides functionality to operate with earthquake data
package earthquake

// Earthquake represents an earthquake entity
type Earthquake struct {
	Time      string  	`xml:"origin>time>value"`
	Longitude float64 	`xml:"origin>longitude>value"`
	Latitude  float64 	`xml:"origin>latitude>value"`
	Magnitude float64 	`xml:"magnitude>mag>value"`
}

// Earthquake collection is a container for Earthquake slice
type EarthquakeCollection struct {
	Earthquakes []Earthquake `xml:"eventParameters>event"`
}