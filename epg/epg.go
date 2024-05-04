package epg

import (
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Channel struct to represent the channel element
type Channel struct {
	ID          string      `xml:"id,attr"`
	DisplayName DisplayName `xml:"display-name"`
	Icon        Icon        `xml:"icon"`
}

// DisplayName struct to represent the display-name element
type DisplayName struct {
	// Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

// Icon struct to represent the icon element
type Icon struct {
	Src string `xml:"src,attr"`
}

// Programme struct to represent the programme element
type Programme struct {
	Start   string `xml:"start,attr"`
	Stop    string `xml:"stop,attr"`
	Channel string `xml:"channel,attr"`
	Title   Title  `xml:"title"`
	Desc    Desc   `xml:"desc"`
}

// Title struct to represent the title element
type Title struct {
	// Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

// Desc struct to represent the desc element
type Desc struct {
	// Lang  string `xml:"lang,attr"`
	Value string `xml:",chardata"`
}

// TV struct to represent the root tv element
type TV struct {
	Channels   []Channel   `xml:"channel"`
	Programmes []Programme `xml:"programme"`
}

func ParseEPG(url string) (*TV, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Printf("Redirected from %s to %s\n", via[len(via)-1].URL, req.URL)
			return nil // Allow all redirects
		},
		Timeout: time.Second * 100, // Set a timeout to avoid indefinite waits
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Wrap the io.ReadCloser with a gzip reader
	gzipReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return nil, err
	}

	defer gzipReader.Close()

	// Create an xml.Decoder
	decoder := xml.NewDecoder(gzipReader)
	// Create a variable to hold the data
	var tv TV

	// Use the decoder to decode the XML stream into the tv variable
	err = decoder.Decode(&tv)
	if err != nil {
		return nil, err
	}

	return &tv, nil
}

const EpgTimeLayout = "20060102150405 -0700"

func (programme *Programme) GetStart() (time.Time, error) {
	start, err := time.Parse(EpgTimeLayout, programme.Start)
	if err != nil {
		log.Println("Cannot parse time", programme.Start, err)
		return start, err
	}

	return start, nil
}

func (programme *Programme) GetStop() (time.Time, error) {
	stop, err := time.Parse(EpgTimeLayout, programme.Stop)
	if err != nil {
		log.Println("Cannot parse time", programme.Stop, err)
		return stop, err
	}

	return stop, nil
}
