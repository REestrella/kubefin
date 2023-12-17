//go:build e2e
// +build e2e

/*
Copyright 2022 The KubeFin Authors

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

package e2e

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"kubefin.dev/kubefin/pkg/api"
	"kubefin.dev/kubefin/test/e2e/utils"
)

func TestAllClustersMetricsSummary(t *testing.T) {
	t.Parallel()

	body, code, err := utils.DoGetRequest(utils.E2ETestEndpoint, utils.AllClusterMetricsSummaryPath)
	if err != nil || code != http.StatusOK {
		t.Fatalf("Get all clusters metrics summary error:%v, %d", err, code)
	}
	allClustersSummary := api.ClusterResourcesSummaryList{}
	err = json.Unmarshal(body, &allClustersSummary)
	if err != nil {
		t.Fatalf("Marshal clusters metrics summary error:%v", err)
	}
	if !utils.ValidateAllClustersMetricsSummary(&allClustersSummary) {
		t.Fatalf("Validate clusters metrics summary response error:%s", string(body))
	}
}

func TestSpecificClusterMetricsSummary(t *testing.T) {
	t.Parallel()

	getClusterMetricsSummary := func(clusterId string) {
		path := fmt.Sprintf(utils.SpecificClusterMetricsSummaryPath, clusterId)
		body, code, err := utils.DoGetRequest(utils.E2ETestEndpoint, path)
		if err != nil || code != http.StatusOK {
			t.Fatalf("Get specific cluster metrics summary error:%v, %d", err, code)
		}
		clusterSummary := api.ClusterResourcesSummary{}
		err = json.Unmarshal(body, &clusterSummary)
		if err != nil {
			t.Fatalf("Marshal cluster metrics summary error:%v", err)
		}
		if !utils.ValidateSpecificClusterMetricsSummary(&clusterSummary) {
			t.Fatalf("Validate cluster metrics summary response error:%s", string(body))
		}
	}
	primaryClusterID, err := utils.GetPrimaryClusterID()
	if err != nil {
		t.Fatalf("Get primary cluster id error:%v", err)
	}
	getClusterMetricsSummary(primaryClusterID)

	secondaryClusterID, err := utils.GetSecondaryClusterID()
	if err != nil {
		t.Fatalf("Get primary cluster id error:%v", err)
	}
	getClusterMetricsSummary(secondaryClusterID)
}

func TestAllClustersCostsSummary(t *testing.T) {
	t.Parallel()

	body, code, err := utils.DoGetRequest(utils.E2ETestEndpoint, utils.AllClusterCostsSummaryPath)
	if err != nil || code != http.StatusOK {
		t.Fatalf("Get all clusters costs summary error:%v, %d", err, code)
	}
	allClustersSummary := api.ClusterCostsSummaryList{}
	err = json.Unmarshal(body, &allClustersSummary)
	if err != nil {
		t.Fatalf("Marshal clusters costs summary error:%v", err)
	}
	if !utils.ValidateAllClustersCostsSummary(&allClustersSummary) {
		t.Fatalf("Validate clusters costs summary response error:%s", string(body))
	}
}

func TestSpecificClusterCostsSummary(t *testing.T) {
	t.Parallel()

	getClusterCostsSummary := func(clusterId string) {
		path := fmt.Sprintf(utils.SpecificClusterCostsSummaryPath, clusterId)
		body, code, err := utils.DoGetRequest(utils.E2ETestEndpoint, path)
		if err != nil || code != http.StatusOK {
			t.Fatalf("Get specific cluster costs summary error:%v, %d", err, code)
		}
		clusterSummary := api.ClusterCostsSummary{}
		err = json.Unmarshal(body, &clusterSummary)
		if err != nil {
			t.Fatalf("Marshal cluster costs summary error:%v", err)
		}
		if !utils.ValidateSpecificClusterCostsSummary(&clusterSummary) {
			t.Fatalf("Validate cluster costs summary response error:%s", string(body))
		}
	}
	primaryClusterID, err := utils.GetPrimaryClusterID()
	if err != nil {
		t.Fatalf("Get primary cluster id error:%v", err)
	}
	getClusterCostsSummary(primaryClusterID)

	secondaryClusterID, err := utils.GetSecondaryClusterID()
	if err != nil {
		t.Fatalf("Get primary cluster id error:%v", err)
	}
	getClusterCostsSummary(secondaryClusterID)
}
