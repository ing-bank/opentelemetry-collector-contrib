// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package inmemorymetricexporter

import (
	"fmt"
	"time"

	"go.opentelemetry.io/collector/component"
)

type Config struct {
	Enabled          bool          `mapstructure:"enabled"`
	MetricExpiration time.Duration `mapstructure:"metric_expiration"`
}

var _ component.Config = (*Config)(nil)

func (cfg *Config) Validate() error {
	if cfg == nil {
		return fmt.Errorf("Invalid in-memory metric exporter config")
	}
	return nil
}
