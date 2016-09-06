package ltsvExporter

type ltsvScraper interface {
	scrape() ([]byte, error)
}
