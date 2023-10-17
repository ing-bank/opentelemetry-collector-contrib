package inmemorymetricexporter

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

type MetricsExporter struct {
	metrics     pmetric.Metrics
	accumulator accumulator

	shutdownC chan struct{}
	logger    *zap.Logger
}

func (e *MetricsExporter) storeMetrics(ctx context.Context, metrics pmetric.Metrics) error {
	e.metrics = metrics
	e.consumeMetrics(ctx, metrics)
	return nil
}

func (e *MetricsExporter) processMetrics(rm pmetric.ResourceMetrics) (n int) {
	return e.accumulator.Accumulate(rm)
}

func (e *MetricsExporter) consumeMetrics(_ context.Context, md pmetric.Metrics) {
	n := 0
	rmetrics := md.ResourceMetrics()
	for i := 0; i < rmetrics.Len(); i++ {
		n += e.processMetrics(rmetrics.At(i))
	}
}

func (e *MetricsExporter) GetMetric(metricName string) (pmetric.Metric, bool) {
	metric, ok := e.accumulator.(*lastValueAccumulator).registeredMetrics.Load(metricName)
	if ok {
		return metric.(*accumulatedValue).value, true
	}
	return pmetric.Metric{}, false
}

func (e *MetricsExporter) Shutdown(ctx context.Context) error {
	// TODO: pass ctx to goroutines so that we can use its deadline
	close(e.shutdownC)
	c := make(chan struct{})
	go func() {
		close(c)
	}()
	select {
	case <-ctx.Done():
		err := errors.New("error waiting for async tasks to finish")
		if err != nil {
			return err
		}
	case <-c:
	}
	return nil
}
