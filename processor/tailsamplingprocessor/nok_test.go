// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// This file contains tests for the Nok (not ok) configuration extension. It
// is kept separate from the upstream test files since this extension is not
// intended to be merged back upstream.

package tailsamplingprocessor

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/client"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumertest"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor/processortest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/pkg/samplingpolicy"
)

// TestNokForwardsNotSampledTraces generates multiple traces that result in
// both a sampled and not-sampled decision, and verifies - using the actual
// trace data received by the next consumer rather than telemetry metrics -
// that:
//   - the sampled trace is forwarded, as usual.
//   - the not-sampled traces are also forwarded, because Nok is enabled,
//     instead of being dropped.
//   - the decision made for each trace is indeed the expected one, confirmed
//     via the sampled/non-sampled decision hooks.
func TestNokForwardsNotSampledTraces(t *testing.T) {
	controller := newTestTSPController()

	var sampledHookCalls []struct {
		id pcommon.TraceID
		td *TraceData
	}
	var nonSampledHookCalls []struct {
		id pcommon.TraceID
		td *TraceData
	}

	sampledHook := func(_ context.Context, id pcommon.TraceID, td *TraceData) {
		sampledHookCalls = append(sampledHookCalls, struct {
			id pcommon.TraceID
			td *TraceData
		}{id: id, td: td})
	}

	nonSampledHook := func(_ context.Context, id pcommon.TraceID, td *TraceData) {
		nonSampledHookCalls = append(nonSampledHookCalls, struct {
			id pcommon.TraceID
			td *TraceData
		}{id: id, td: td})
	}

	cfg := Config{
		SamplingStrategy: samplingStrategyTraceComplete,
		DecisionWait:     defaultTestDecisionWait,
		NumTraces:        defaultNumTraces,
		PolicyCfgs: []PolicyCfg{
			{
				sharedPolicyCfg: sharedPolicyCfg{
					Name: "prod-only",
					Type: StringAttribute,
					StringAttributeCfg: StringAttributeCfg{
						Key:    "env",
						Values: []string{"prod"},
					},
				},
			},
		},
		Nok: NokConfig{
			Enabled:      true,
			ContextKey:   "tenant",
			ContextValue: "dev-st",
			DefaultValue: "dev",
			Action:       POSTFIX,
		},
		Options: []Option{
			withTestController(controller),
			WithSampledHooks(sampledHook),
			WithNonSampledHooks(nonSampledHook),
		},
	}

	msp := new(consumertest.TracesSink)
	p, err := newTracesProcessor(t.Context(), processortest.NewNopSettings(metadata.Type), msp, cfg)
	require.NoError(t, err)
	require.NoError(t, p.Start(t.Context(), componenttest.NewNopHost()))
	defer func() {
		require.NoError(t, p.Shutdown(t.Context()))
	}()

	// One trace matches the policy and is sampled.
	sampledTraceID := uInt64ToTraceID(1)
	sampledTrace := simpleTracesWithID(sampledTraceID)
	sampledTrace.ResourceSpans().At(0).ScopeSpans().At(0).Spans().At(0).Attributes().PutStr("env", "prod")

	// Two traces don't match the policy and are not sampled; they are
	// nevertheless expected to be forwarded because Nok is enabled.
	nonSampledTraceID1 := uInt64ToTraceID(2)
	nonSampledTrace1 := simpleTracesWithID(nonSampledTraceID1)
	nonSampledTrace1.ResourceSpans().At(0).ScopeSpans().At(0).Spans().At(0).Attributes().PutStr("env", "staging")

	nonSampledTraceID2 := uInt64ToTraceID(3)
	nonSampledTrace2 := simpleTracesWithID(nonSampledTraceID2)
	nonSampledTrace2.ResourceSpans().At(0).ScopeSpans().At(0).Spans().At(0).Attributes().PutStr("env", "dev")

	require.NoError(t, p.ConsumeTraces(t.Context(), sampledTrace))
	require.NoError(t, p.ConsumeTraces(t.Context(), nonSampledTrace1))
	require.NoError(t, p.ConsumeTraces(t.Context(), nonSampledTrace2))

	controller.waitForTick() // the first tick always gets an empty batch
	controller.waitForTick() // the second tick makes the decisions

	// verify the decision made for each trace.
	require.Len(t, sampledHookCalls, 1, "only the prod trace should be sampled")
	assert.Equal(t, sampledTraceID, sampledHookCalls[0].id)
	assert.Equal(t, samplingpolicy.Sampled, sampledHookCalls[0].td.FinalDecision)

	require.Len(t, nonSampledHookCalls, 2, "the staging and dev traces should not be sampled")
	nonSampledIDs := map[pcommon.TraceID]struct{}{}
	for _, call := range nonSampledHookCalls {
		assert.Equal(t, samplingpolicy.NotSampled, call.td.FinalDecision)
		nonSampledIDs[call.id] = struct{}{}
	}
	assert.Contains(t, nonSampledIDs, nonSampledTraceID1)
	assert.Contains(t, nonSampledIDs, nonSampledTraceID2)

	// verify all three traces were forwarded to the next consumer: the
	// sampled one as usual, and the not-sampled ones because Nok is enabled.
	allTraces := msp.AllTraces()
	require.Len(t, allTraces, 3)

	gotSampled := findTrace(t, allTraces, sampledTraceID)
	gotSampledSpan := gotSampled.ResourceSpans().At(0).ScopeSpans().At(0).Spans().At(0)
	env, ok := gotSampledSpan.Attributes().Get("env")
	require.True(t, ok)
	assert.Equal(t, "prod", env.Str())

	gotNonSampled1 := findTrace(t, allTraces, nonSampledTraceID1)
	gotNonSampled1Span := gotNonSampled1.ResourceSpans().At(0).ScopeSpans().At(0).Spans().At(0)
	env, ok = gotNonSampled1Span.Attributes().Get("env")
	require.True(t, ok)
	assert.Equal(t, "staging", env.Str())

	gotNonSampled2 := findTrace(t, allTraces, nonSampledTraceID2)
	gotNonSampled2Span := gotNonSampled2.ResourceSpans().At(0).ScopeSpans().At(0).Spans().At(0)
	env, ok = gotNonSampled2Span.Attributes().Get("env")
	require.True(t, ok)
	assert.Equal(t, "dev", env.Str())
}

// capturingConsumer records the context each batch is forwarded with, so tests
// can assert on the client metadata injected by Nok.
type capturingConsumer struct {
	mu       sync.Mutex
	captured []capturedTrace
}

type capturedTrace struct {
	ctx context.Context
	td  ptrace.Traces
}

func (*capturingConsumer) Capabilities() consumer.Capabilities {
	return consumer.Capabilities{MutatesData: false}
}

func (c *capturingConsumer) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	clone := ptrace.NewTraces()
	td.CopyTo(clone)
	c.captured = append(c.captured, capturedTrace{ctx: ctx, td: clone})
	return nil
}

func (c *capturingConsumer) all() []capturedTrace {
	c.mu.Lock()
	defer c.mu.Unlock()
	return append([]capturedTrace(nil), c.captured...)
}

// metadataFor returns the client metadata values the given trace ID was
// forwarded with.
func (c *capturingConsumer) metadataFor(t *testing.T, id pcommon.TraceID, key string) []string {
	t.Helper()
	for _, got := range c.all() {
		rss := got.td.ResourceSpans()
		for i := 0; i < rss.Len(); i++ {
			sss := rss.At(i).ScopeSpans()
			for j := 0; j < sss.Len(); j++ {
				spans := sss.At(j).Spans()
				for k := 0; k < spans.Len(); k++ {
					if spans.At(k).TraceID() == id {
						return client.FromContext(got.ctx).Metadata.Get(key)
					}
				}
			}
		}
	}
	t.Fatalf("trace %s was not forwarded", id)
	return nil
}

// TestNokInjectsContextValue verifies that not-sampled traces forwarded by Nok
// carry the configured context value as client metadata, for both the regular
// decision path and the late-arriving span path.
func TestNokInjectsContextValue(t *testing.T) {
	controller := newTestTSPController()

	cfg := Config{
		SamplingStrategy: samplingStrategyTraceComplete,
		DecisionWait:     defaultTestDecisionWait,
		NumTraces:        defaultNumTraces,
		PolicyCfgs: []PolicyCfg{
			{
				sharedPolicyCfg: sharedPolicyCfg{
					Name: "prod-only",
					Type: StringAttribute,
					StringAttributeCfg: StringAttributeCfg{
						Key:    "env",
						Values: []string{"prod"},
					},
				},
			},
		},
		Nok: NokConfig{
			Enabled:      true,
			ContextKey:   "tenant",
			ContextValue: "-nok",
			DefaultValue: "dev",
			Action:       POSTFIX,
		},
		Options: []Option{withTestController(controller)},
	}

	sink := &capturingConsumer{}
	p, err := newTracesProcessor(t.Context(), processortest.NewNopSettings(metadata.Type), sink, cfg)
	require.NoError(t, err)
	require.NoError(t, p.Start(t.Context(), componenttest.NewNopHost()))
	defer func() {
		require.NoError(t, p.Shutdown(t.Context()))
	}()

	sampledTraceID := uInt64ToTraceID(1)
	sampledTrace := simpleTracesWithID(sampledTraceID)
	sampledTrace.ResourceSpans().At(0).ScopeSpans().At(0).Spans().At(0).Attributes().PutStr("env", "prod")

	nonSampledTraceID := uInt64ToTraceID(2)
	nonSampledTrace := simpleTracesWithID(nonSampledTraceID)
	nonSampledTrace.ResourceSpans().At(0).ScopeSpans().At(0).Spans().At(0).Attributes().PutStr("env", "staging")

	require.NoError(t, p.ConsumeTraces(t.Context(), sampledTrace))
	require.NoError(t, p.ConsumeTraces(t.Context(), nonSampledTrace))

	controller.waitForTick() // the first tick always gets an empty batch
	controller.waitForTick() // the second tick makes the decisions

	// The not-sampled trace is forwarded with the nok value injected. No value
	// is present on the incoming context, so the default value is used as base.
	assert.Equal(t, []string{"dev-nok"}, sink.metadataFor(t, nonSampledTraceID, "tenant"))
	// The sampled trace is forwarded untouched.
	assert.Empty(t, sink.metadataFor(t, sampledTraceID, "tenant"))

	// Late-arriving spans on the not-sampled trace must be injected as well.
	lateTrace := simpleTracesWithID(nonSampledTraceID)
	lateTrace.ResourceSpans().At(0).ScopeSpans().At(0).Spans().At(0).Attributes().PutStr("env", "staging")
	before := len(sink.all())
	require.NoError(t, p.ConsumeTraces(t.Context(), lateTrace))
	controller.waitForTick()

	forwarded := sink.all()[before:]
	require.Len(t, forwarded, 1, "late spans must be forwarded exactly once")
	assert.Equal(t, []string{"dev-nok"}, client.FromContext(forwarded[0].ctx).Metadata.Get("tenant"))
	assert.Equal(t, 1, forwarded[0].td.SpanCount())
}
