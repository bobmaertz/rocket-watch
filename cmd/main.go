package main

import (
	"flag"
	"time"
	"os"
	"fmt"

	n "github.com/bou1der/rocket-watch/pkg/provider/nasa"
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

	prov, err := n.NewProvider()

	resp, err := prov.GetLaunches()

	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(2)
	}

	for i, hit := range resp.Hits.Hits {
		fmt.Println("Launch: ", i)
		fmt.Println(hit.Source.Description)
	}
}
