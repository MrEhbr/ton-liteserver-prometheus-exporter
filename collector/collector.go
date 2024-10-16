// collector/collector.go
package collector

import (
	"log"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

type MytonCollector struct {
	metrics       []MetricDef
	parsingErrors prometheus.Counter
	mutex         sync.Mutex
	parser        *Parser
}

const (
	MetricNamespace = "ton_liteserver"
	MetricSubsystem = "exporter"
)

func NewMytonCollector(parser *Parser) *MytonCollector {
	return &MytonCollector{
		metrics: Metrics,
		parsingErrors: prometheus.NewCounter(prometheus.CounterOpts{
			Name: prometheus.BuildFQName(MetricNamespace, MetricSubsystem, "parsing_errors_total"),
			Help: "Total number of parsing errors encountered during metric collection",
		}),
		parser: parser,
	}
}

func (collector *MytonCollector) Describe(ch chan<- *prometheus.Desc) {
	for _, mDef := range collector.metrics {
		ch <- mDef.desc
	}
	ch <- collector.parsingErrors.Desc()
}

func (collector *MytonCollector) Collect(ch chan<- prometheus.Metric) {
	collector.mutex.Lock()
	defer collector.mutex.Unlock()

	metrics, err := collector.parser.Parse()
	if err != nil {
		log.Printf("Error collecting metrics: %v", err)
		collector.parsingErrors.Inc()
		return
	}

	for _, mDef := range collector.metrics {
		value, labels := mDef.getValue(metrics)
		ch <- prometheus.MustNewConstMetric(mDef.desc, prometheus.GaugeValue, value, labels...)
	}

	ch <- collector.parsingErrors
}
