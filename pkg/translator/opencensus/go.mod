module github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/opencensus

go 1.22.0

require (
	github.com/census-instrumentation/opencensus-proto v0.4.1
	github.com/golang/protobuf v1.5.4
	github.com/google/go-cmp v0.6.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.117.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest v0.117.1-0.20250117002813-e970f8bb1258
	github.com/stretchr/testify v1.10.0
	go.opencensus.io v0.24.0
	go.opentelemetry.io/collector/pdata v1.23.1-0.20250119231113-f07ebc3afb51
	go.opentelemetry.io/collector/semconv v0.117.1-0.20250119231113-f07ebc3afb51
	go.uber.org/goleak v1.3.0
	google.golang.org/protobuf v1.36.2
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.18.1 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.117.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241015192408-796eee8c2d53 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241104194629-dd2ea8efbc28 // indirect
	google.golang.org/grpc v1.69.4 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../../pkg/pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../../internal/coreinternal

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../golden
