// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package trace

import (
	"testing"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper/operatortest"
)

func TestConfig(t *testing.T) {
	cases := []operatortest.ConfigUnmarshalTest{
		{
			Name:   "default",
			Expect: defaultCfg(),
		},
		{
			Name: "on_error_drop",
			Expect: func() *Config {
				cfg := defaultCfg()
				cfg.OnError = "drop"
				return cfg
			}(),
		},
		{
			Name: "spanid",
			Expect: func() *Config {
				parseFrom := entry.NewBodyField("app_span_id")
				cfg := helper.SpanIDConfig{}
				cfg.ParseFrom = &parseFrom

				c := defaultCfg()
				c.SpanID = &cfg
				return c
			}(),
		},
		{
			Name: "traceid",
			Expect: func() *Config {
				parseFrom := entry.NewBodyField("app_trace_id")
				cfg := helper.TraceIDConfig{}
				cfg.ParseFrom = &parseFrom

				c := defaultCfg()
				c.TraceID = &cfg
				return c
			}(),
		},
		{
			Name: "trace_flags",
			Expect: func() *Config {
				parseFrom := entry.NewBodyField("app_trace_flags_id")
				cfg := helper.TraceFlagsConfig{}
				cfg.ParseFrom = &parseFrom

				c := defaultCfg()
				c.TraceFlags = &cfg
				return c
			}(),
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			tc.Run(t, defaultCfg())
		})
	}
}

func defaultCfg() *Config {
	return NewConfig("trace_parser")
}
