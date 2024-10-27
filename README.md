# VisualStudio Marketplace Prometheus Exporter

[![Build](https://github.com/soerenuhrbach/visualstudio-marketplace-exporter/actions/workflows/ci.yml/badge.svg)](https://github.com/soerenuhrbach/visualstudio-marketplace-exporter/actions/workflows/ci.yml)
[![GoDoc](https://godoc.org/github.com/soerenuhrbach/visualstudio-marketplace-exporter?status.png)](https://godoc.org/github.com/soerenuhrbach/visualstudio-marketplace-exporter)
[![GoReportCard](https://goreportcard.com/badge/github.com/soerenuhrbach/visualstudio-marketplace-exporter)](https://goreportcard.com/report/github.com/soerenuhrbach/visualstudio-marketplace-exporter)

Prometheus exporter to scrape metrics about extensions of the visual studio marketplace.

## Usage

### Using release binaries 

You can download the latest version of the binary built for your architecture here:

* Architecture **i386** [
    [Linux](https://github.com/soerenuhrbach/visualstudio-marketplace-exporter/releases/latest/download/visualstudio_marketplace_exporter-linux-386) /
    [Windows](https://github.com/soerenuhrbach/visualstudio-marketplace-exporter/releases/latest/download/visualstudio_marketplace_exporter-windows-386.exe)
]
* Architecture **amd64** [
    [Darwin](https://github.com/soerenuhrbach/visualstudio-marketplace-exporter/releases/latest/download/visualstudio_marketplace_exporter-darwin-amd64) /
    [Linux](https://github.com/soerenuhrbach/visualstudio-marketplace-exporter/releases/latest/download/visualstudio_marketplace_exporter-linux-amd64) /
    [Windows](https://github.com/soerenuhrbach/visualstudio-marketplace-exporter/releases/latest/download/visualstudio_marketplace_exporter-windows-amd64.exe)
]
* Architecture **arm** [
    [Darwin](https://github.com/soerenuhrbach/visualstudio-marketplace-exporter/releases/latest/download/visualstudio_marketplace_exporter-darwin-arm64) /
    [Linux](https://github.com/soerenuhrbach/visualstudio-marketplace-exporter/releases/latest/download/visualstudio_marketplace_exporter-linux-arm)
]

You can run it using the following example:

```bash
./visualstudio_marketplace_exporter -extensions "soerenuhrbach.vscode"
```

### Docker 

The exporter is also available as a [Docker image](https://github.com/soerenuhrbach/visualstudio-marketplace-exporter/pkgs/container/visualstudio-marketplace-exporter).
You can run it using the following example and pass configuration environment variables:

```bash
docker run \
  -e 'EXTENSIONS=soerenuhrbach.vscode-deepl' \
  -p 9719:9719 \
  ghcr.io/soerenuhrbach/visualstudio-marketplace-exporter:latest
```

## Scrape metrics

Once the exporter is running, you can configure your collector to scrape the metrics. 

Adjust your `prometheus.yml` configuration to let it scrape the exporter like this:

```yaml
scrape_configs:
  - job_name: 'visualstudio_marketplace'
    static_configs:
      - targets: ['localhost:9719']
```

## Available configurations

|Configuration name|Description|Required|Argument|Environment variable|Default|
|---|---|---|---|---|---|
|Extensions|Comma-separated list of extensions that should be scraped|Required|extensions|EXTENSIONS|-|
|Metric path|Port to be used for the exporter|false|metricsPath|METRICSPATH|/metrics|
|Port|Port to be used for the exporter|false|port|PORT|9719|
|Binding Address|Address to be used for the exporter|false|bindAddress|BINDADDRESS|0.0.0.0|

Examples with all possible configurations:

```bash
./visualstudio_marketplace_exporter \
  -extensions "soerenuhrbach.vscode" \
  -metricsPath "/metrics" \
  -port 9719 \
  -bindAddress 0.0.0.0
```
or using docker:

```bash
docker run \
  -e 'EXTENSIONS=soerenuhrbach.vscode-deepl' \
  -e 'METRICSPATH=/metrics' \
  -e 'PORT=9719' \
  -e 'BINDADDRESS=0.0.0.0' \
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