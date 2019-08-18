package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	nasa "github.com/bou1der/rocket-watch/pkg/provider/nasa"
)

// DataSource describes where the launch data came from.
type DataSource struct {
	name string
	url  string
	date time.Time
}

// Launch represents rocket launch events.
type Launch struct {
	origin   DataSource
	name     string
	location string
	date     time.Time
}

var (
	//config = flag.String("config", "", "The path to the config file")
	checkNasa = flag.Bool("nasa", false, "Check nasa for upcoming launches")
)

func main() {
	flag.Parse()

	sourceName := "NASA"
	prov, err := nasa.NewProvider()
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(2)
	}

	resp, err := prov.GetLaunches(0, 10)
	fmt.Println("--------------------------")
	fmt.Println("Requesting from source: ", sourceName)

	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(2)
	}

	for i, hit := range resp.Hits.Hits {
		fmt.Println("--------------------------")
		fmt.Println("Launch: ", i)
		fmt.Println("Title: ", hit.Source.Title)
		fmt.Println("Source: ", sourceName)
		fmt.Println("Expected Launch Date: ", hit.Source.EventDate[0].Value)
		fmt.Println("Description: ", hit.Source.Description)
	}
}
