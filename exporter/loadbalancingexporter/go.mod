module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter

go 1.22.0

require (
	github.com/aws/aws-sdk-go-v2/config v1.28.11
	github.com/aws/aws-sdk-go-v2/service/servicediscovery v1.34.4
	github.com/aws/smithy-go v1.22.1
	github.com/json-iterator/go v1.1.12
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics v0.117.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchpersignal v0.117.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden v0.117.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest v0.117.1-0.20250117002813-e970f8bb1258
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/component v0.117.1-0.20250119231113-f07ebc3afb51
	go.opentelemetry.io/collector/component/componenttest v0.117.1-0.20250119231113-f07ebc3afb51
	go.opentelemetry.io/collector/config/configretry v1.23.1-0.20250119231113-f07ebc3afb51
	go.opentelemetry.io/collector/config/configtelemetry v0.117.1-0.20250119231113-f07ebc3afb51
	go.opentelemetry.io/collector/confmap v1.23.1-0.20250119231113-f07ebc3afb51
	go.opentelemetry.io/collector/consumer v1.23.1-0.20250119231113-f07ebc3afb51
	go.opentelemetry.io/collector/consumer/consumertest v0.117.1-0.20250119231113-f07ebc3afb51
	go.opentelemetry.io/collector/exporter v0.117.1-0.20250119231113-f07ebc3afb51
	go.opentelemetry.io/collector/exporter/exportertest v0.117.1-0.20250119231113-f07ebc3afb51
	go.opentelemetry.io/collector/exporter/otlpexporter v0.117.1-0.20250119231113-f07ebc3afb51
	go.opentelemetry.io/collector/otelcol/otelcoltest v0.117.1-0.20250119231113-f07ebc3afb51
	go.opentelemetry.io/collector/pdata v1.23.1-0.20250119231113-f07ebc3afb51
	go.opentelemetry.io/collector/semconv v0.117.1-0.20250119231113-f07ebc3afb51
	go.opentelemetry.io/otel v1.32.0
	go.opentelemetry.io/otel/metric v1.32.0
	go.opentelemetry.io/otel/sdk v1.32.0
	go.opentelemetry.io/otel/sdk/metric v1.32.0
	go.opentelemetry.io/otel/trace v1.32.0
	go.uber.org/goleak v1.3.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.31.3
	k8s.io/apimachinery v0.31.3
	k8s.io/client-go v0.31.3
	k8s.io/utils v0.0.0-20240711033017-18e509b52bc8
	sigs.k8s.io/controller-runtime v0.19.4
)

require (
	github.com/aws/aws-sdk-go-v2 v1.32.8 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.17.52 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.23 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.27 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.27 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.12.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.12.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.24.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.28.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.33.7 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/ebitengine/purego v0.8.1 // indirect
	github.com/emicklei/go-restful/v3 v3.11.0 // indirect
	github.com/evanphx/json-patch/v5 v5.9.0 // indirect
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/fxamacker/cbor/v2 v2.7.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.22.4 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.23.0 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/imdario/mergo v0.3.6 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.2 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mostynb/go-grpc-compression v1.2.3 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil v0.117.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/prometheus/client_golang v1.20.5 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.61.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/shirou/gopsutil/v4 v4.24.12 // indirect
	github.com/spf13/cobra v1.8.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.opentelemetry.io/collector/client v1.23.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/component/componentstatus v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/config/configauth v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/config/configcompression v1.23.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/config/configgrpc v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/config/confignet v1.23.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/config/configopaque v1.23.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/config/configtls v1.23.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/confmap/provider/envprovider v1.23.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/confmap/provider/fileprovider v1.23.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/confmap/provider/httpprovider v1.23.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/confmap/provider/yamlprovider v1.23.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/connector v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/connector/connectortest v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/connector/xconnector v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/consumer/consumererror v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/consumer/consumererror/xconsumererror v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/consumer/xconsumer v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/exporter/exporterhelper/xexporterhelper v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/exporter/xexporter v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/extension v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/extension/auth v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/extension/extensioncapabilities v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/extension/extensiontest v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/extension/xextension v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/featuregate v1.23.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/internal/fanoutconsumer v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/otelcol v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/pdata/testdata v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/pipeline v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/pipeline/xpipeline v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/processor v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/processor/processortest v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/processor/xprocessor v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/receiver v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/receiver/receivertest v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/receiver/xreceiver v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/collector/service v0.117.1-0.20250119231113-f07ebc3afb51 // indirect
	go.opentelemetry.io/contrib/bridges/otelzap v0.6.0 // indirect
	go.opentelemetry.io/contrib/config v0.10.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.56.0 // indirect
	go.opentelemetry.io/contrib/propagators/b3 v1.31.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp v0.7.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v1.32.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v1.32.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.31.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.31.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.31.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.54.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutlog v0.7.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.32.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.31.0 // indirect
	go.opentelemetry.io/otel/log v0.8.0 // indirect
	go.opentelemetry.io/otel/sdk/log v0.7.0 // indirect
	go.opentelemetry.io/proto/otlp v1.3.1 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/oauth2 v0.24.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/term v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/time v0.4.0 // indirect
	gonum.org/v1/gonum v0.15.1 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241104194629-dd2ea8efbc28 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241104194629-dd2ea8efbc28 // indirect
	google.golang.org/grpc v1.69.4 // indirect
	google.golang.org/protobuf v1.36.2 // indirect
	gopkg.in/evanphx/json-patch.v4 v4.12.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/klog/v2 v2.130.1 // indirect
	k8s.io/kube-openapi v0.0.0-20240228011516-70dd3763d340 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchpersignal => ../../pkg/batchpersignal

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

// ambiguous import: found package cloud.google.com/go/compute/metadata in multiple modules
replace cloud.google.com/go v0.65.0 => cloud.google.com/go v0.110.10

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../pkg/golden

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics => ../../internal/exp/metrics
