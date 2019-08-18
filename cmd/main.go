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
	date time.Time
}

// Launch represents a rocket launch event.
type Launch struct {
	origin      *DataSource
	name        string
	location    string
	date        time.Time
	description string
}

func main() {
	checkNasa := flag.Bool("nasa", false, "Check nasa for upcoming launches")

	flag.Parse()

	if *checkNasa == true {
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

		//TODO: Sort in chronological order
		launches := make([]Launch, len(resp.Hits.Hits))
		for i, hit := range resp.Hits.Hits {
			t, _ := time.Parse(time.RFC3339, hit.Source.EventDate[0].Value)
			launches[i] = Launch{
				origin: &DataSource{
					name: sourceName,
					date: time.Now(),
				},
				name:        hit.Source.Title,
				date:        t,
				description: hit.Source.Description,
			}
		}

		for _, launch := range launches {
			fmt.Println("--------------------------")
			fmt.Println("Title: ", launch.name)
			fmt.Println("Source: ", launch.origin.name)
			fmt.Println("Expected Launch Date: ", launch.date)
			fmt.Println("Description: ", launch.description)
		}
	}
}
