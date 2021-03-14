// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/hyperledger/fabric/orderer/common/follower"
	"github.com/hyperledger/fabric/orderer/common/types"
)

type ChannelParticipationMetricsReporter struct {
	ReportConsensusRelationAndStatusMetricsStub        func(string, types.ConsensusRelation, types.Status)
	reportConsensusRelationAndStatusMetricsMutex       sync.RWMutex
	reportConsensusRelationAndStatusMetricsArgsForCall []struct {
		arg1 string
		arg2 types.ConsensusRelation
		arg3 types.Status
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ChannelParticipationMetricsReporter) ReportConsensusRelationAndStatusMetrics(arg1 string, arg2 types.ConsensusRelation, arg3 types.Status) {
	fake.reportConsensusRelationAndStatusMetricsMutex.Lock()
	fake.reportConsensusRelationAndStatusMetricsArgsForCall = append(fake.reportConsensusRelationAndStatusMetricsArgsForCall, struct {
		arg1 string
		arg2 types.ConsensusRelation
		arg3 types.Status
	}{arg1, arg2, arg3})
	fake.recordInvocation("ReportConsensusRelationAndStatusMetrics", []interface{}{arg1, arg2, arg3})
	fake.reportConsensusRelationAndStatusMetricsMutex.Unlock()
	if fake.ReportConsensusRelationAndStatusMetricsStub != nil {
		fake.ReportConsensusRelationAndStatusMetricsStub(arg1, arg2, arg3)
	}
}

func (fake *ChannelParticipationMetricsReporter) ReportConsensusRelationAndStatusMetricsCallCount() int {
	fake.reportConsensusRelationAndStatusMetricsMutex.RLock()
	defer fake.reportConsensusRelationAndStatusMetricsMutex.RUnlock()
	return len(fake.reportConsensusRelationAndStatusMetricsArgsForCall)
}

func (fake *ChannelParticipationMetricsReporter) ReportConsensusRelationAndStatusMetricsCalls(stub func(string, types.ConsensusRelation, types.Status)) {
	fake.reportConsensusRelationAndStatusMetricsMutex.Lock()
	defer fake.reportConsensusRelationAndStatusMetricsMutex.Unlock()
	fake.ReportConsensusRelationAndStatusMetricsStub = stub
}

func (fake *ChannelParticipationMetricsReporter) ReportConsensusRelationAndStatusMetricsArgsForCall(i int) (string, types.ConsensusRelation, types.Status) {
	fake.reportConsensusRelationAndStatusMetricsMutex.RLock()
	defer fake.reportConsensusRelationAndStatusMetricsMutex.RUnlock()
	argsForCall := fake.reportConsensusRelationAndStatusMetricsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ChannelParticipationMetricsReporter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.reportConsensusRelationAndStatusMetricsMutex.RLock()
	defer fake.reportConsensusRelationAndStatusMetricsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ChannelParticipationMetricsReporter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ follower.ChannelParticipationMetricsReporter = new(ChannelParticipationMetricsReporter)
