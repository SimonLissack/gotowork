package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/simonlissack/gotowork/config"
	"github.com/simonlissack/gotowork/directions"
)

func collectFlags() (configurationPath, origin, destination string) {
	flag.StringVar(&configurationPath, "config", "config.json", "Path to config file")
	flag.StringVar(&origin, "origin", "", "Path to config file")
	flag.StringVar(&destination, "destination", "", "Path to config file")

	flag.Parse()

	return
}

func main() {
	configurationPath, origin, destination := collectFlags()

	configFile, err := ioutil.ReadFile(configurationPath)
	logFatal(err)

	config, err := config.Load(configFile)
	logFatal(err)

	journey := directions.Journey{
		Origin:      origin,
		Destination: destination,
	}

	route := directions.GetRoute(journey, *config)[0]
	timeAsString := (route.ExpectedTravelTime / time.Nanosecond).String()

	fmt.Println("Via:", route.Summary)
	fmt.Println("distance:", route.Distance/1000, "km")
	fmt.Println("Estimate time:", timeAsString)
	fmt.Println(route.Copyright)
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
