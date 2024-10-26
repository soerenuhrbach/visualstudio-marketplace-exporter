package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/soerenuhrbach/visualstudio-marketplace-exporter/config"
	"github.com/soerenuhrbach/visualstudio-marketplace-exporter/internal/exporter"
)

func main() {
	cfg := config.Load()

	http.Handle(cfg.MetricsPath, promhttp.Handler())

	exporter := exporter.NewVisualStudioMarketPlaceExporter(cfg.Extensions)
	prometheus.MustRegister(exporter)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>VisualStudio Marketplace Exporter</title></head>
             <body>
             <h1>VisualStudio Marketplace Exporter</h1>
             <p><a href='` + cfg.MetricsPath + `'>Metrics</a></p>
             </body>
             </html>`))
	})

	listenAddress := fmt.Sprintf("%s:%d", cfg.BindAddress, cfg.Port)
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}
