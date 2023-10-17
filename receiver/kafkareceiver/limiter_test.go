package kafkareceiver

import (
	"context"
	"github.com/IBM/sarama"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type MockConsumerGroup struct {
	mock.Mock
}

func (m *MockConsumerGroup) Consume(ctx context.Context, topics []string, handler sarama.ConsumerGroupHandler) error {
	m.Called()
	return nil
}

func (m *MockConsumerGroup) Errors() <-chan error {
	m.Called()
	return nil
}

func (m *MockConsumerGroup) Close() error {
	m.Called()
	return nil
}

func (m *MockConsumerGroup) Pause(partitions map[string][]int32) {
	m.Called()
}

func (m *MockConsumerGroup) Resume(partitions map[string][]int32) {
	m.Called()
}

func (m *MockConsumerGroup) PauseAll() {
	m.Called()
}

func (m *MockConsumerGroup) ResumeAll() {
	m.Called()
}

func TestShouldPause(t *testing.T) {
	limiter := Limiter{
		Enabled: true,
	}

	tests := []struct {
		values limiterValues
	}{
		{limiterValues{currValue: 100, prevValue: 50, limit: 1, currTs: currTs(), prevTs: currTs()}},
	}

	for _, test := range tests {
		prevTs := test.values.currTs.AsTime()
		prevTs = prevTs.Add(-30 * time.Second)
		test.values.prevTs = pcommon.Timestamp(uint64(prevTs.UnixNano()))
		pause, rate := limiter.shouldPause(test.values)
		assert.Equal(t, pause, true)
		assert.Greater(t, rate, float64(1))
	}
}
