// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package encode

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func Test_Gauge(t *testing.T) {
	resource := pcommon.NewResource()
	metric := pmetric.NewMetric()
	metric.SetName("gauge-metric")
	metric.SetDataType(pmetric.MetricDataTypeGauge)

	dp := metric.Gauge().DataPoints().AppendEmpty()
	dp.SetIntVal(1)
	dp.SetTimestamp(createTimestamp())

	scope := createInstrumentationScope()

	documents, _ := Gauge(&resource, &scope, &metric)
	assert.Equal(t, `{"@timestamp":"2022-01-01T10:00:05.000000123Z","name":"gauge-metric","scope":{"name":"scope1","version":"v1"},"type":"gauge","value":1}`, string(documents[0]))
}

func Test_SumCounter(t *testing.T) {
	resource := pcommon.NewResource()
	metric := pmetric.NewMetric()
	metric.SetName("counter-metric")
	metric.SetDataType(pmetric.MetricDataTypeSum)

	sum := metric.Sum()
	sum.SetIsMonotonic(true)

	dp := sum.DataPoints().AppendEmpty()
	dp.SetIntVal(1)
	dp.SetTimestamp(createTimestamp())

	scope := createInstrumentationScope()

	documents, _ := Sum(&resource, &scope, &metric)
	assert.Equal(t, `{"@timestamp":"2022-01-01T10:00:05.000000123Z","name":"counter-metric","scope":{"name":"scope1","version":"v1"},"type":"counter","value":1}`, string(documents[0]))
}

func Test_SumGauge(t *testing.T) {
	resource := pcommon.NewResource()
	metric := pmetric.NewMetric()
	metric.SetName("gauge-metric")
	metric.SetDataType(pmetric.MetricDataTypeSum)

	sum := metric.Sum()
	sum.SetIsMonotonic(false)

	dp := sum.DataPoints().AppendEmpty()
	dp.SetIntVal(1)
	dp.SetTimestamp(createTimestamp())

	scope := createInstrumentationScope()

	documents, _ := Sum(&resource, &scope, &metric)
	assert.Equal(t, `{"@timestamp":"2022-01-01T10:00:05.000000123Z","name":"gauge-metric","scope":{"name":"scope1","version":"v1"},"type":"gauge","value":1}`, string(documents[0]))
}

func Test_Summary(t *testing.T) {
	resource := pcommon.NewResource()
	metric := pmetric.NewMetric()
	metric.SetDataType(pmetric.MetricDataTypeSummary)
	metric.SetName("summary-metric")

	dp := metric.Summary().DataPoints().AppendEmpty()
	dp.SetCount(1)
	dp.SetSum(10.1)
	dp.SetTimestamp(createTimestamp())
	dp.Attributes().InsertString("dpLabel", "val1")

	quantile1 := dp.QuantileValues().AppendEmpty()
	quantile1.SetValue(10.2)
	quantile1.SetQuantile(0.9)
	quantile2 := dp.QuantileValues().AppendEmpty()
	quantile2.SetValue(10.5)
	quantile2.SetQuantile(0.95)

	scope := createInstrumentationScope()

	documents, _ := Summary(&resource, &scope, &metric)
	assert.Equal(t, 1, len(documents))
	assert.Equal(t, `{"@timestamp":"2022-01-01T10:00:05.000000123Z","attributes":{"dpLabel":"val1"},"count":1,"name":"summary-metric","quantileValues":{"0_9":10.2,"0_95":10.5},"scope":{"name":"scope1","version":"v1"},"sum":10.1,"type":"summary"}`, string(documents[0]))
}

func Test_Histogram(t *testing.T) {
	resource := pcommon.NewResource()
	metric := pmetric.NewMetric()
	metric.SetName("hist-metric")
	metric.SetDataType(pmetric.MetricDataTypeHistogram)

	dp := metric.Histogram().DataPoints().AppendEmpty()
	dp.SetTimestamp(createTimestamp())
	dp.SetCount(4)
	dp.SetBucketCounts([]uint64{1, 3, 0})
	dp.SetExplicitBounds([]float64{5.87, 10})

	scope := createInstrumentationScope()

	documents, _ := Histogram(&resource, &scope, &metric)
	assert.Equal(t, 3, len(documents))
	assert.Equal(t, `{"@timestamp":"2022-01-01T10:00:05.000000123Z","bucket":"5.87","bucketCount":1,"bucketId":0,"bucketText":"0 (5.87)","count":4,"name":"hist-metric","scope":{"name":"scope1","version":"v1"},"sum":0,"type":"histogram"}`, string(documents[0]))
	assert.Equal(t, `{"@timestamp":"2022-01-01T10:00:05.000000123Z","bucket":"10","bucketCount":3,"bucketId":1,"bucketText":"1 (10)","count":4,"name":"hist-metric","scope":{"name":"scope1","version":"v1"},"sum":0,"type":"histogram"}`, string(documents[1]))
	assert.Equal(t, `{"@timestamp":"2022-01-01T10:00:05.000000123Z","bucket":"+inf","bucketCount":0,"bucketId":2,"bucketText":"2 (+inf)","count":4,"name":"hist-metric","scope":{"name":"scope1","version":"v1"},"sum":0,"type":"histogram"}`, string(documents[2]))
}

func createInstrumentationScope() pcommon.InstrumentationScope {
	lib := pcommon.NewInstrumentationScope()
	lib.SetName("scope1")
	lib.SetVersion("v1")
	return lib
}

func createTimestamp() pcommon.Timestamp {
	return pcommon.NewTimestampFromTime(time.Date(2022, time.January, 1, 10, 00, 5, 123, time.UTC))
}
