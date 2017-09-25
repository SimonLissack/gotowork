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

var (
	configPath, origin, destination string
)

func init() {
	flag.StringVar(&configPath, "config", "", "Path to config file")
	flag.StringVar(&origin, "origin", "", "Path to config file")
	flag.StringVar(&destination, "destination", "", "Path to config file")
	flag.Parse()

	if origin == "" {
		log.Fatal("No origin address given")
	}

	if destination == "" {
		log.Fatal("No destination address given")
	}

	if configPath == "" {
		log.Fatal("No configuration file given")
	}
}

func main() {
	configFile, err := ioutil.ReadFile(configPath)
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
