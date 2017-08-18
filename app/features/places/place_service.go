package places

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"strings"

	"googlemaps.github.io/maps"
)

var (
	apiKey    = flag.String("key", "AIzaSyBAG-KEfEG9SyDKCuHvoOIeMD1hOqwtdtk", "API Key for using Google Maps API.")
	radius    = flag.Uint("radius", 1609, "Defines the distance (in meters) within which to bias place results. The maximum allowed radius is 50,000 meters.")
	placeType = flag.String("type", "restaurant", "Restricts the results to places matching the specified type.")
)

// GetPlacesService This has a comment.
func GetPlacesService(requestLocation string) (data []Place, err error) {

	flag.Parse()
	var client *maps.Client

	client, err = maps.NewClient(maps.WithAPIKey(*apiKey))

	if check(err) {
		return
	}

	r := &maps.NearbySearchRequest{
		Radius: *radius,
		Type:   maps.PlaceType(*placeType),
	}

	parseLocation(requestLocation, r)

	resp, err := client.NearbySearch(context.Background(), r)

	if check(err) {
		return
	}
	randomNumber := rand.Intn(len(resp.Results))

	data = append(data,
		Place{
			resp.Results[randomNumber].Name,
			strings.Join(resp.Results[randomNumber].Types, ", "),
			resp.Results[randomNumber].PriceLevel,
			resp.Results[randomNumber].Rating,
			1.2,
			"https://maps.googleapis.com/maps/api/place/photo?maxwidth=400&photoreference=" + resp.Results[randomNumber].Photos[0].PhotoReference + "&key=" + *apiKey,
		},
	)
	return
}

func parseLocation(location string, r *maps.NearbySearchRequest) {
	if location != "" {
		l, _ := maps.ParseLatLng(location)
		r.Location = &l
	}
}

func check(err error) bool {
	if err != nil {
		log.Fatalf("fatal error: %s", err)
		return true
	}
	return false
}
