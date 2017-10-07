package directions

import (
	"strconv"
	"time"

	"github.com/simonlissack/gotowork/config"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

// Journey defines the origin and destination of a route
type Journey struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}

// Route describes the journey (distance is measured in meters)
type Route struct {
	Summary            string        `json:"summary"`
	Origin             string        `json:"origin"`
	Destination        string        `json:"destination"`
	Distance           int           `json:"distance"`
	UsualTravelTime    time.Duration `json:"usualTravelTime"`
	ExpectedTravelTime time.Duration `json:"expectedTravelTime"`
	Copyright          string        `json:"copyright"`
}

// GetRoute gets the route for a journey
func GetRoute(journey Journey, config config.GTWConfiguration) ([]Route, error) {
	client, _ := maps.NewClient(maps.WithAPIKey(config.APIKey))
	departureTime := strconv.FormatInt(time.Now().Unix(), 10)

	directions := maps.DirectionsRequest{
		Origin:        journey.Origin,
		Destination:   journey.Destination,
		DepartureTime: departureTime,
	}

	response, _, err := client.Directions(context.Background(), &directions)

	if err != nil {
		return nil, err
	}

	routes := make([]Route, len(response))

	for i, route := range response {
		var expecetedTime, usualTime time.Duration
		rt := Route{}
		distance := 0

		for _, leg := range route.Legs {
			usualTime += leg.Duration
			expecetedTime += leg.DurationInTraffic
			distance += leg.Distance.Meters
		}

		rt.Origin = journey.Origin
		rt.Summary = route.Summary
		rt.Destination = journey.Destination
		rt.Distance = distance
		rt.ExpectedTravelTime = expecetedTime
		rt.UsualTravelTime = usualTime
		rt.Copyright = route.Copyrights

		routes[i] = rt
	}

	return routes, nil
}
