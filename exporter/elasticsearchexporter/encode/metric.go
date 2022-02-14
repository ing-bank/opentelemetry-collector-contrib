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

package encode // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/encode"

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/objmodel"
)

// Gauge serializes gauge metric data points into JSON documents.
func Gauge(resource *pcommon.Resource, scope *pcommon.InstrumentationScope, metric *pmetric.Metric) ([][]byte, []error) {
	var documents [][]byte
	var errs []error
	gauge := metric.Gauge()
	for i := 0; i < gauge.DataPoints().Len(); i++ {
		var document objmodel.Document
		dp := gauge.DataPoints().At(i)
		document.AddString("type", "gauge")
		setCommonFields(&document, resource, scope, metric, dp.Timestamp(), dp.Attributes())
		setNumberValue(&document, dp)
		ser, err := serializeDocument(&document)
		if err == nil {
			documents = append(documents, ser)
		} else {
			errs = append(errs, err)
		}
	}
	return documents, errs
}

// Sum serializes sum metric data points into JSON documents.
func Sum(resource *pcommon.Resource, scope *pcommon.InstrumentationScope, metric *pmetric.Metric) ([][]byte, []error) {
	var documents [][]byte
	var errs []error
	sum := metric.Sum()
	for i := 0; i < sum.DataPoints().Len(); i++ {
		var document objmodel.Document
		dp := sum.DataPoints().At(i)
		if dp.ValueType() == pmetric.NumberDataPointValueTypeNone {
			continue
		}
		if sum.IsMonotonic() {
			document.AddString("type", "counter")
		} else {
			document.AddString("type", "gauge")
		}
		setCommonFields(&document, resource, scope, metric, dp.Timestamp(), dp.Attributes())
		setNumberValue(&document, dp)
		ser, err := serializeDocument(&document)
		if err == nil {
			documents = append(documents, ser)
		} else {
			errs = append(errs, err)
		}
	}
	return documents, errs
}

// Summary serializes summary metric data points into JSON documents.
func Summary(resource *pcommon.Resource, scope *pcommon.InstrumentationScope, metric *pmetric.Metric) ([][]byte, []error) {
	var documents [][]byte
	var errs []error
	summary := metric.Summary()
	for i := 0; i < summary.DataPoints().Len(); i++ {
		var document objmodel.Document
		dp := summary.DataPoints().At(i)
		document.AddUInt("count", dp.Count())
		document.AddDouble("sum", dp.Sum())
		quantiles := pcommon.NewMap()
		qv := dp.QuantileValues()
		for j := 0; j < qv.Len(); j++ {
			qvj := qv.At(j)
			k := strconv.FormatFloat(qvj.Quantile(), 'f', -1, 64)
			// workaround as dedup messes with floating point keys
			k = strings.Replace(k, ".", "_", -1)
			quantiles.InsertDouble(k, qvj.Value())
		}
		document.AddAttributes("quantileValues", quantiles)
		document.AddString("type", "summary")
		setCommonFields(&document, resource, scope, metric, dp.Timestamp(), dp.Attributes())
		ser, err := serializeDocument(&document)
		if err == nil {
			documents = append(documents, ser)
		} else {
			errs = append(errs, err)
		}
	}
	return documents, errs
}

// Histogram serializes sum histogram data points into JSON documents.
func Histogram(resource *pcommon.Resource, scope *pcommon.InstrumentationScope, metric *pmetric.Metric) ([][]byte, []error) {
	var documents [][]byte
	var errs []error
	hist := metric.Histogram()
	for i := 0; i < hist.DataPoints().Len(); i++ {
		dp := hist.DataPoints().At(i)
		explicitBounds := dp.ExplicitBounds()
		maxBucketIDLen := len(fmt.Sprintf("%d", len(dp.BucketCounts())))
		for bucketID, v := range dp.BucketCounts() {
			var document objmodel.Document
			document.AddUInt("count", dp.Count())
			document.AddDouble("sum", dp.Sum())
			document.AddUInt("bucketId", uint64(bucketID))
			// bucket bounds are defined as 'less than', we add +inf as last bucket
			const lastBucket = "+inf"
			if bucketID < len(explicitBounds) {
				bucket := formatBucket(explicitBounds[bucketID])
				document.AddString("bucket", bucket)
				// bucketText can be used to display the histogram visualization in the correct order
				document.AddString("bucketText", fmt.Sprintf("%0*d (%s)", maxBucketIDLen, bucketID, bucket))
			} else {
				document.AddString("bucket", lastBucket)
				document.AddString("bucketText", fmt.Sprintf("%0*d (%s)", maxBucketIDLen, bucketID, lastBucket))
			}
			document.AddUInt("bucketCount", v)
			document.AddString("type", "histogram")
			setCommonFields(&document, resource, scope, metric, dp.Timestamp(), dp.Attributes())
			ser, err := serializeDocument(&document)
			if err == nil {
				documents = append(documents, ser)
			} else {
				errs = append(errs, err)
			}
		}
	}
	return documents, errs
}

func formatBucket(bucket float64) string {
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%f", bucket), "0"), ".")
}

func setCommonFields(document *objmodel.Document, resource *pcommon.Resource, scope *pcommon.InstrumentationScope, metric *pmetric.Metric, timestamp pcommon.Timestamp, attributes pcommon.Map) {
	document.AddTimestamp("@timestamp", timestamp)
	document.AddString("name", metric.Name())
	document.AddAttributes("resource", resource.Attributes())
	document.AddAttributes("attributes", attributes)
	scopeMap := pcommon.NewMap()
	scopeMap.InsertString("name", scope.Name())
	scopeMap.InsertString("version", scope.Version())
	document.AddAttributes("scope", scopeMap)
}

func setNumberValue(document *objmodel.Document, dp pmetric.NumberDataPoint) {
	switch dp.ValueType() {
	case pmetric.NumberDataPointValueTypeDouble:
		document.AddDouble("value", dp.DoubleVal())
	case pmetric.NumberDataPointValueTypeInt:
		document.AddInt("value", dp.IntVal())
	}
}

func serializeDocument(document *objmodel.Document) ([]byte, error) {
	document.Dedup()
	document.Sort()
	var buf bytes.Buffer
	err := document.Serialize(&buf, true)
	return buf.Bytes(), err
}
