package ltsvExporter

import (
	"io/ioutil"
	"net/http"
)

type remoteLTSVScraper struct {
	ltsvURL string
}

func newRemoteLTSVScraper(ltsvURL string) *remoteLTSVScraper {
	return &remoteLTSVScraper{
		ltsvURL: ltsvURL,
	}
}

func (e *remoteLTSVScraper) scrape() ([]byte, error) {
	ltsvResp, err := http.Get(e.ltsvURL)
	if err != nil {
		return nil, err
	}
	defer ltsvResp.Body.Close()

	body, err := ioutil.ReadAll(ltsvResp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
