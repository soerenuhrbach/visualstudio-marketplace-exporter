package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/soerenuhrbach/visualstudio-marketplace-exporter/internal/exporter"
)

var (
	// config
	metricsPath   = "/metrics"
	listenAddress = ":9719"
)

func main() {
	http.Handle(metricsPath, promhttp.Handler())

	exporter := exporter.NewVisualStudioMarketPlaceExporter("soerenuhrbach.vscode-deepl")
	prometheus.MustRegister(exporter)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>VisualStudio Marketplace Exporter</title></head>
             <body>
             <h1>VisualStudio Marketplace Exporter</h1>
             <p><a href='` + metricsPath + `'>Metrics</a></p>
             </body>
             </html>`))
	})

	log.Fatal(http.ListenAndServe(listenAddress, nil))
}
