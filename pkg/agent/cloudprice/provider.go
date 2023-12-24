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

package cloudprice

import (
	"context"
	"fmt"
	"strings"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"

	"kubefin.dev/kubefin/cmd/kubefin-agent/app/options"
	"kubefin.dev/kubefin/pkg/agent/cloudprice/ack"
	"kubefin.dev/kubefin/pkg/agent/cloudprice/defaultcloud"
	"kubefin.dev/kubefin/pkg/agent/cloudprice/eks"
	"kubefin.dev/kubefin/pkg/agent/types"
	"kubefin.dev/kubefin/pkg/values"
)

type CloudProviderInterface interface {
	GetNodeHourlyPrice(node *v1.Node) (*types.InstancePriceInfo, error)
	Start(ctx context.Context)
}

func NewCloudProvider(client kubernetes.Interface, agentOptions *options.AgentOptions) (CloudProviderInterface, error) {
	switch agentOptions.CloudProvider {
	case values.CloudProviderACK:
		return ack.NewACKCloudProvider(client, agentOptions)
	case values.CloudProviderEKS:
		return eks.NewEKSCloudProvider(client, agentOptions)
	case values.CloudProviderAuto:
		// If config cloud provider is empty or cannot retrieve, checking it automatically
		return initCloudProviderAuto(client, agentOptions)
	}
	return nil, fmt.Errorf("cloud provider %s is not supported", agentOptions.CloudProvider)
}

func initCloudProviderAuto(client kubernetes.Interface, agentOptions *options.AgentOptions) (CloudProviderInterface, error) {
	nodes, err := client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		klog.Errorf("failed to list nodes: %v", err)
		return nil, err
	}
	if len(nodes.Items) == 0 {
		return nil, fmt.Errorf("no nodes found")
	}

	cloudProviderID := strings.ToLower(nodes.Items[0].Spec.ProviderID)
	if strings.HasPrefix(cloudProviderID, "aws") {
		return eks.NewEKSCloudProvider(client, agentOptions)
	}
	if strings.HasPrefix(cloudProviderID, "gce") {
		klog.Warning("KubeFin doesn't support GCE yet, default pricing data will be used")
	}
	if strings.HasPrefix(cloudProviderID, "azure") {
		klog.Warning("KubeFin doesn't support Azure yet, default pricing data will be used")
	}
	agentOptions.CloudProvider = values.CloudProviderOnPremise
	return defaultcloud.NewDefaultCloudProvider(client, agentOptions)
}
