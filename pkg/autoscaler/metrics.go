/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package autoscaler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
)

// InitializeMetrics and export metrics.
func InitializeMetrics(options *Options) {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "ok (%v)\n", time.Now())
	})

	go func() {
		err := http.ListenAndServe(
			fmt.Sprintf("%s:%d", options.PrometheusAddr, options.PrometheusPort), nil)
		if err != nil {
			glog.Fatalf("Error starting metrics server: %v", err)
		}
	}()
}
