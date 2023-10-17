// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate mdatagen metadata.yaml

package inmemorymetricexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudexporter"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/featuregate"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/inmemorymetricexporter/internal/metadata"
)

var _ = featuregate.GlobalRegistry().MustRegister(
	"exporter.nmemorymetric",
	featuregate.StageStable,
	featuregate.WithRegisterDescription("When enabled, the pdata metrics are stored in memory and made available to consumption."),
	featuregate.WithRegisterToVersion("v0.86.0"),
)

// NewFactory creates a factory for the googlecloud exporter
func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		metadata.Type,
		createDefaultConfig,
		exporter.WithMetrics(createMetricsExporter, metadata.MetricsStability),
	)
}

// createDefaultConfig creates the default configuration for exporter.
func createDefaultConfig() component.Config {
	return &Config{
		Enabled:          false,
		MetricExpiration: 60 * time.Second,
	}
}

var InMemoryMetricsExporter MetricsExporter

// createMetricsExporter creates a metrics exporter based on this config.
func createMetricsExporter(
	ctx context.Context,
	params exporter.CreateSettings,
	cfg component.Config) (exporter.Metrics, error) {
	eCfg := cfg.(*Config)
	mExp := MetricsExporter{logger: params.Logger, accumulator: newAccumulator(params.Logger, eCfg.MetricExpiration)}
	InMemoryMetricsExporter = mExp
	mExp.logger.Info("Starting in-memory Metric Exporter")
	return exporterhelper.NewMetricsExporter(
		ctx,
		params,
		eCfg,
		mExp.storeMetrics,
		exporterhelper.WithShutdown(mExp.Shutdown),
		exporterhelper.WithTimeout(exporterhelper.TimeoutSettings{Timeout: 0}),
	)
}
