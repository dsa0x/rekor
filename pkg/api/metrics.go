//
// Copyright 2021 The Sigstore Authors.
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

package api

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	metricNewEntries = promauto.NewCounter(prometheus.CounterOpts{
		Name: "rekor_new_entries",
		Help: "The total number of new log entries",
	})

	MetricLatency = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "rekor_api_latency",
		Help: "Api Latency on calls",
	}, []string{"path", "code"})

	MetricLatencySummary = promauto.NewSummaryVec(prometheus.SummaryOpts{
		Name: "rekor_api_latency_summary",
		Help: "Api Latency on calls",
	}, []string{"path", "code"})
)
