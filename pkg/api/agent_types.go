/*
Copyright 2023 The KubeFin Authors

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

package api

import (
	appv1 "k8s.io/client-go/listers/apps/v1"
	v1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"

	insightlister "github.com/kubefin/kubefin/pkg/generated/listers/insight/v1alpha1"
)

type InstancePriceInfo struct {
	NodeTotalHourlyPrice float64
	CPUCore              float64
	CPUCoreHourlyPrice   float64
	RamGiB               float64
	RAMGiBHourlyPrice    float64
	GPUCards             float64
	GPUCardHourlyPrice   float64
	InstanceType         string
	BillingMode          string

	// Only need for monthly/yearly billing mode
	BillingPeriod int
	Region        string
	CloudProvider string
}

type CoreResourceInformerLister struct {
	NodeInformer              cache.SharedIndexInformer
	NamespaceInformer         cache.SharedIndexInformer
	PodInformer               cache.SharedIndexInformer
	DeploymentInformer        cache.SharedIndexInformer
	StatefulSetInformer       cache.SharedIndexInformer
	DaemonSetInformer         cache.SharedIndexInformer
	CustomWorkloadCfgInformer cache.SharedIndexInformer
	NodeLister                v1.NodeLister
	PodLister                 v1.PodLister
	DeploymentLister          appv1.DeploymentLister
	StatefulSetLister         appv1.StatefulSetLister
	DaemonSetLister           appv1.DaemonSetLister
	CustomWorkloadCfgLister   insightlister.CustomAllocationConfigurationLister
}
