// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package places

import (
	"fmt"
	"flag"
	"os"
	"log"

	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
	"net/http"

	"encoding/json"
)

var (
	apiKey = flag.String("key", "AIzaSyBAG-KEfEG9SyDKCuHvoOIeMD1hOqwtdtk", "API Key for using Google Maps API.")
	location  = flag.String("location", "41.584627,-93.637551", "The latitude/longitude around which to retrieve place information. This must be specified as latitude,longitude.")
	radius    = flag.Uint("radius", 1609, "Defines the distance (in meters) within which to bias place results. The maximum allowed radius is 50,000 meters.")
	placeType = flag.String("type", "restaurant", "Restricts the results to places matching the specified type.")
)


func GetPlacesController(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	flag.Parse()
	var client *maps.Client
	var err error

	if *apiKey != "" {
		client, err = maps.NewClient(maps.WithAPIKey(*apiKey))
	} else {
		usageAndExit("Please specify an API Key, or Client ID and Signature.")
	}

	check(err)

	r := &maps.NearbySearchRequest{
		Radius:    *radius,
		Type: 	 maps.PlaceType(*placeType),
	}

	parseLocation(*location, r)

	resp, err := client.NearbySearch(context.Background(), r)
	check(err)

	json.NewEncoder(writer).Encode(resp)
}

func parseLocation(location string, r *maps.NearbySearchRequest) {
	if location != "" {
		l, err := maps.ParseLatLng(location)
		check(err)
		r.Location = &l
	}
}

func usageAndExit(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	fmt.Println("Flags:")
	flag.PrintDefaults()
	os.Exit(2)
}


func check(err error) {
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
}
