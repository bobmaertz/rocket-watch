package nasa

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ProviderImpl struct {
	client *http.Client
	url    string
}

// NewProvider Creates a new launch provider source
func NewProvider() (*ProviderImpl, error) {
	client := &http.Client{}

 	return &ProviderImpl{client: client, url: "https://www.nasa.gov/api/2/"}, nil
}

// GetLaunches gets upcoming launches from Nasa Calendar
func (p *ProviderImpl) GetLaunches() (*Response, error) {
	r, err := p.client.Get("https://www.nasa.gov/api/2/calendar-event/_search?size=100&from=0&q=calendar-name:6089")

	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	res := &Response{}

	temp, _ := ioutil.ReadAll(r.Body)

	err = json.Unmarshal(temp, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//Nasa Calendar Response

type Response struct {
	Took     int     `json:"took"`
	TimedOut bool    `json:"timed_out"`
	Shards   Shards  `json:"_shards"`
	Hits     BigHits `json:"Hits"`
}
type Shards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Failed     int `json:"failed"`
}
type AdditionalLink1 struct {
	URL        string        `json:"url"`
	Title      string        `json:"title"`
	Attributes []interface{} `json:"attributes"`
}
type EventDate struct {
	Value      string      `json:"value"`
	Value2     string      `json:"value2"`
	Rrule      interface{} `json:"rrule"`
	Timezone   string      `json:"timezone"`
	TimezoneDb string      `json:"timezone_db"`
	DateType   string      `json:"date_type"`
}
type MasterImage struct {
	ID     string `json:"id"`
	Fid    string `json:"fid"`
	URI    string `json:"uri"`
	Width  string `json:"width"`
	Height string `json:"height"`
}

type Source struct {
	EventDateCount  int               `json:"event-date-count"`
	Title           string            `json:"title"`
	Nid             string            `json:"nid"`
	Type            string            `json:"type"`
	Changed         string            `json:"changed"`
	UUID            string            `json:"uuid"`
	Name            string            `json:"name"`
	URI             string            `json:"uri"`
	AdditionalLink1 []AdditionalLink1 `json:"additional-link1"`
	CalendarName    []string          `json:"calendar-name"`
	Description     string            `json:"description"`
	EventDate       []EventDate       `json:"event-date"`
	MasterImage     MasterImage       `json:"master-image"`
}

type Hits struct {
	Index  string  `json:"_index"`
	Type   string  `json:"_type"`
	ID     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Source Source  `json:"_source,omitempty"`
}

type BigHits struct {
	Total    int     `json:"total"`
	MaxScore float64 `json:"max_score"`
	Hits     []Hits  `json:"hits"`
}
