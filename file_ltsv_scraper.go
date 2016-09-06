package ltsvExporter

import "io/ioutil"

type fileLTSVScraper struct {
	ltsvPath string
}

func newFileLTSVScraper(ltsvPath string) *fileLTSVScraper {
	return &fileLTSVScraper{
		ltsvPath: ltsvPath,
	}
}

func (e *fileLTSVScraper) scrape() ([]byte, error) {
	return ioutil.ReadFile(e.ltsvPath)
}
