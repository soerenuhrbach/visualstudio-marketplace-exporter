package config

import (
	"context"
	"fmt"
	"reflect"
	"runtime"

	log "github.com/sirupsen/logrus"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/flags"
)

type ExporterConfig struct {
	Extensions  []string `config:"extensions"`
	MetricsPath string   `config:"metricsPath"`
	BindAddress string   `config:"bindAddress"`
	Port        uint16   `config:"port"`
}

func getDefaultConfig() *ExporterConfig {
	return &ExporterConfig{
		Extensions:  []string{},
		BindAddress: "0.0.0.0",
		Port:        9719,
		MetricsPath: "/metrics",
	}
}

func Load() *ExporterConfig {
	loaders := []backend.Backend{
		env.NewBackend(),
		flags.NewBackend(),
	}

	loader := confita.NewLoader(loaders...)
	cfg := getDefaultConfig()
	err := loader.Load(context.Background(), cfg)
	if err != nil {
		panic(err)
	}

	cfg.show()

	return cfg
}

func (c ExporterConfig) show() {
	val := reflect.ValueOf(&c).Elem()
	log.Info("-----------------------------------")
	log.Info("-  		Exporter configuration  		-")
	log.Info("-----------------------------------")
	log.Info("Go version: ", runtime.Version())
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		log.Info(fmt.Sprintf("%s : %v", typeField.Name, valueField.Interface()))
	}
	log.Info("-----------------------------------")
}
