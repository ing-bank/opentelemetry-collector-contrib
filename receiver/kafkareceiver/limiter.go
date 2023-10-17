package kafkareceiver

import (
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/inmemorymetricexporter"
)

type Limiter struct {
	Interval                   time.Duration `mapstructure:"interval"`
	ReceiverAcceptedSpansRate  float64       `mapstructure:"receiver_accepted_spans_rate"`
	ProcessorAcceptedSpansRate float64       `mapstructure:"processor_accepted_spans_rate"`
	ExporterSentSpansRate      float64       `mapstructure:"exporter_sent_spans_rate"`
	ExporterQueueSizeDiff      float64       `mapstructure:"exporter_queue_size_diff"`
	KafkaReceiverMessagesRate  float64       `mapstructure:"kafka_receiver_messages_rate"`
	KafkaReceiverLagDiff       float64       `mapstructure:"kafka_receiver_lag_diff"`
	Enabled                    bool          `mapstructure:"enabled"`
	logger                     *zap.Logger
	rateLimited                bool
}

type limiterValues struct {
	currValue  float64
	prevValue  float64
	limit      float64
	currTs     pcommon.Timestamp
	prevTs     pcommon.Timestamp
	metricType pmetric.MetricType
}

func (c Limiter) run(handler sarama.ConsumerGroup) {

	metricMap := make(map[string]limiterValues)

	metricMap["otelcol_receiver_accepted_spans"] = limiterValues{currTs: currTs(), limit: c.ReceiverAcceptedSpansRate, metricType: pmetric.MetricTypeSum}
	metricMap["otelcol_processor_accepted_spans"] = limiterValues{currTs: currTs(), limit: c.ProcessorAcceptedSpansRate, metricType: pmetric.MetricTypeSum}
	metricMap["otelcol_exporter_sent_spans"] = limiterValues{currTs: currTs(), limit: c.ExporterSentSpansRate, metricType: pmetric.MetricTypeSum}
	metricMap["otelcol_exporter_queue_size"] = limiterValues{currTs: currTs(), limit: c.ExporterQueueSizeDiff, metricType: pmetric.MetricTypeGauge}
	metricMap["otelcol_kafka_receiver_messages"] = limiterValues{currTs: currTs(), limit: c.KafkaReceiverMessagesRate, metricType: pmetric.MetricTypeSum}
	metricMap["otelcol_kafka_receiver_offset_lag"] = limiterValues{currTs: currTs(), limit: c.KafkaReceiverLagDiff, metricType: pmetric.MetricTypeGauge}

	for {
		c.logger.Debug("running limiter")

		var pauseRequired bool

		for metricName, values := range metricMap {
			mTs, mVal := c.getMetricValue(metricName, values.metricType)
			values.prevTs = values.currTs
			values.prevValue = values.currValue
			values.currTs = mTs
			values.currValue = mVal
			metricMap[metricName] = values

			pause, rate := c.shouldPause(values)
			c.logger.Debug(fmt.Sprintf("rate for %s: %f (limit: %f)", metricName, rate, values.limit))

			if pause {
				if pauseRequired {
					c.logger.Warn(fmt.Sprintf("rate for %s exceeded: %f (limit: %f). Was already paused, will continue to pause", metricName, rate, values.limit))
				} else {
					c.logger.Warn(fmt.Sprintf("rate for %s exceeded: %f (limit: %f). Going to start pause", metricName, rate, values.limit))
					pauseRequired = true
				}
			}

		}

		if pauseRequired {
			if c.Enabled {
				handler.PauseAll()
				c.rateLimited = true
			} else {
				c.logger.Debug("rate would be limited but limiter is not enabled")
			}
		} else {
			if c.rateLimited {
				c.rateLimited = false
				c.logger.Warn("lifting rate limit, resuming operations")
			}
			if c.Enabled {
				handler.ResumeAll()
			}
		}

		time.Sleep(c.Interval)
	}
}

func (c Limiter) shouldPause(values limiterValues) (bool, float64) {
	previousTs := values.prevTs.AsTime()
	currentTs := values.currTs.AsTime()
	duration := currentTs.Sub(previousTs)
	seconds := duration.Seconds()

	// if any of the values are 0, do not do anything. This is to avoid unintended pausing at startup.
	if values.prevValue != 0 || values.currValue != 0 {

		if seconds > 0 {
			diff := values.currValue - values.prevValue
			rate := diff / seconds

			if rate > values.limit {
				return true, rate
			}
			return false, rate
		}
	}

	return false, 0
}

func (c Limiter) getMetricValue(metricName string, metricType pmetric.MetricType) (pcommon.Timestamp, float64) {
	inMemExp := inmemorymetricexporter.InMemoryMetricsExporter

	metric, metricOk := inMemExp.GetMetric(metricName)
	if metricOk {
		mName := metric.Name()
		mType := metric.Type()
		mData := pmetric.NumberDataPoint{}

		if metricType == pmetric.MetricTypeGauge {
			mData = metric.Gauge().DataPoints().At(0)
		} else if metricType == pmetric.MetricTypeSum {
			mData = metric.Sum().DataPoints().At(0)
		}

		mValue := mData.DoubleValue()
		mTs := mData.Timestamp()

		c.logger.Debug(fmt.Sprintf("metric %s (type: %s) had value %f at time %s", mName, mType, mValue, mTs.String()))
		return mTs, mValue
	}
	c.logger.Warn(fmt.Sprintf("no data for %s", metricName))
	return currTs(), 0
}

func currTs() pcommon.Timestamp {
	currTs := time.Now().UnixNano()
	return pcommon.Timestamp(uint64(currTs))
}
