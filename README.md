# VisualStudio Marketplace Prometheus Exporter

[![Build](https://github.com/soerenuhrbach/visualstudio-marketplace-exporter/actions/workflows/build.yml/badge.svg)](https://github.com/soerenuhrbach/visualstudio-marketplace-exporter/actions/workflows/build.yml)
[![GoDoc](https://godoc.org/github.com/soerenuhrbach/visualstudio-marketplace-exporter?status.png)](https://godoc.org/github.com/soerenuhrbach/visualstudio-marketplace-exporter)
[![GoReportCard](https://goreportcard.com/badge/github.com/soerenuhrbach/visualstudio-marketplace-exporter)](https://goreportcard.com/report/github.com/soerenuhrbach/visualstudio-marketplace-exporter)

Prometheus exporter to scrape metrics about extensions of the visual studio marketplace.

## Usage

The exporter is available as a [Docker image](https://github.com/soerenuhrbach/visualstudio-marketplace-exporter/pkgs/container/visualstudio-marketplace-exporter).
You can run it using the following example and pass configuration environment variables:

```bash
docker run \
  -e 'EXTENSIONS=soerenuhrbach.vscode-deepl' \
  -p 9719:9719 \
  ghcr.io/soerenuhrbach/visualstudio-marketplace-exporter:latest
```

Once the exporter is running, you can configure your collector to scrape the metrics. 

Adjust your `prometheus.yml` configuration to let it scrape the exporter like this:

```yaml
scrape_configs:
  - job_name: 'visualstudio_marketplace'
    static_configs:
      - targets: ['localhost:9719']
```

## Available configurations

|Configuration|Environment variable|Description|Required|Default|
|---|---|---|---|---|
|Extensions|EXTENSIONS|Comma-separated list of extensions that should be scraped|Required|-|
|Metric path|METRICSPATH|Port to be used for the exporter|false|/metrics|
|Port|PORT|Port to be used for the exporter|false|9719|
|Binding Address|BINDADDRESS|Address to be used for the exporter|false|0.0.0.0|

Example with all possible configurations:

```bash
docker run \
  -e 'EXTENSIONS=soerenuhrbach.vscode-deepl' \
  -e 'METRICSPATH=/metrics' \
  -e 'PORT=9719' \
  -e 'BINDADDRESS=127.0.0.1' \
  -p 9719:9719 \
  ghcr.io/soerenuhrbach/visualstudio-marketplace-exporter:latest
```

## Available metrics

|Metric|Description|
|---|---|
|visualstudio_marketplace_installs|Number of installations of the extension|
|visualstudio_marketplace_updates|Number of updates of the extension|
|visualstudio_marketplace_ratings|Number of ratings of the extension|
|visualstudio_marketplace_average_rating|Average rating of the extension|
|visualstudio_marketplace_weighted_rating|Weighted rating of the extension|
|visualstudio_marketplace_trending_daily|Daily trending score of the extension|
|visualstudio_marketplace_trending_weekly|Weekly trending score of the extension|
|visualstudio_marketplace_trending_monthly|Monthly trending score of the extension|
|visualstudio_marketplace_downloads|Number of manual extension downloads via web interface|