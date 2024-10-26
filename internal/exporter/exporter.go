package exporter

import (
	"crypto/tls"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/soerenuhrbach/visualstudio-marketplace-exporter/internal/visualstudiomarketplace"
)

const namespace = "visualstudio_marketplace"

var (
	installs = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "installs"),
		"Amount of installations of the extension",
		[]string{"extension", "extensionId"},
		nil,
	)
	updates = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "updates"),
		"Amount of updates of the extension",
		[]string{"extension", "extensionId"},
		nil,
	)
	averageRating = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "average_rating"),
		"Average rating of the extension",
		[]string{"extension", "extensionId"},
		nil,
	)
	weightedRating = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "weighted_rating"),
		"Weighted rating of the extension",
		[]string{"extension", "extensionId"},
		nil,
	)
	ratings = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "ratings"),
		"Amount of ratings of the extension",
		[]string{"extension", "extensionId"},
		nil,
	)
	downloads = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "downloads"),
		"Amount of manual extension downloads via web interface",
		[]string{"extension", "extensionId"},
		nil,
	)
	trendingDaily = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "trending_daily"),
		"Daily trending score of the extension",
		[]string{"extension", "extensionId"},
		nil,
	)
	trendingWeekly = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "trending_weekly"),
		"Weekly trending score of the extensions",
		[]string{"extension", "extensionId"},
		nil,
	)
	trendingMonthly = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "trending_monthly"),
		"Monthly trending score of the extension ",
		[]string{"extension", "extensionId"},
		nil,
	)

	// http client
	tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}
)

type VisualStudioMarketPlaceExporter struct {
	extensionNames []string
}

func (e *VisualStudioMarketPlaceExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- installs
}

func (e *VisualStudioMarketPlaceExporter) Collect(ch chan<- prometheus.Metric) {
	client := &visualstudiomarketplace.VisualStudioMarketplaceClient{
		BaseUrl:    "https://marketplace.visualstudio.com",
		HttpClient: client,
	}

	statistics := client.GetStatistics(e.extensionNames)

	var scrapedMetrics = 0

	for i := range statistics {
		statistic := statistics[i]

		var metric *prometheus.Desc

		switch statistic.Name {
		case "install":
			metric = installs
		case "averagerating":
			metric = averageRating
		case "ratingcount":
			metric = ratings
		case "trendingdaily":
			metric = trendingDaily
		case "trendingmonthly":
			metric = trendingMonthly
		case "trendingweekly":
			metric = trendingWeekly
		case "updateCount":
			metric = updates
		case "weightedRating":
			metric = weightedRating
		case "downloadCount":
			metric = downloads
		}

		if metric != nil {
			ch <- prometheus.MustNewConstMetric(
				metric,
				prometheus.GaugeValue,
				statistic.Value,
				statistic.ExtensionName,
				statistic.ExtensionId,
			)
			scrapedMetrics += 1
		}
	}

	log.Info(fmt.Sprintf("Scraped %d metrics!", scrapedMetrics))
}

func NewVisualStudioMarketPlaceExporter(extensions []string) *VisualStudioMarketPlaceExporter {
	return &VisualStudioMarketPlaceExporter{
		extensionNames: extensions,
	}
}
