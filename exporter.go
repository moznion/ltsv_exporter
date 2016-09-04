package ltsvExporter

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/najeira/ltsv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

const namespace = "ltsv"

type exporter struct {
	mu             sync.Mutex
	ltsvURL        string
	value          prometheus.GaugeVec
	scrapeFailures prometheus.Counter
}

func newExporter(ltsvURL string) *exporter {
	e := &exporter{
		value: *prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "value",
			Help:      "LTSV value",
		}, []string{"key"}),
		scrapeFailures: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "exporter_scrape_failures_total",
			Help:      "Number of errors while scraping",
		}),
		ltsvURL: ltsvURL,
	}

	return e
}

func (e *exporter) Describe(ch chan<- *prometheus.Desc) {
	e.value.Describe(ch)
	e.scrapeFailures.Describe(ch)
}

func (e *exporter) Collect(ch chan<- prometheus.Metric) {
	e.mu.Lock() // To protect metrics from concurrent collects.
	defer e.mu.Unlock()

	err := e.collect(ch)
	if err != nil {
		log.Error(err)
		e.incrementFailures(ch)
	}

	e.value.Collect(ch)
	e.scrapeFailures.Collect(ch)
}

func (e *exporter) collect(ch chan<- prometheus.Metric) error {
	ltsvResp, err := http.Get(e.ltsvURL)
	if err != nil {
		return err
	}
	defer ltsvResp.Body.Close()

	ltsvReader := ltsv.NewReader(ltsvResp.Body)
	ltsvRecords, err := ltsvReader.ReadAll()
	if err != nil {
		return err
	}

	for _, ltsvRecord := range ltsvRecords {
		for k, v := range ltsvRecord {
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				log.Debug("Not a number, ignore [key=%s, value=%s]", k, v)
				continue
			}
			e.value.WithLabelValues(k).Set(f)
		}
	}

	return nil
}

func (e *exporter) incrementFailures(ch chan<- prometheus.Metric) {
	e.scrapeFailures.Inc()
	e.scrapeFailures.Collect(ch)
}