/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Free Trial License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Free-Trial-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	"kubedb.dev/apimachinery/apis/kubedb"
	cs "kubedb.dev/apimachinery/client/clientset/versioned"
	amc "kubedb.dev/apimachinery/pkg/controller"
	"kubedb.dev/apimachinery/pkg/eventer"

	pcm "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned/typed/monitoring/v1"
	crd_cs "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/metadata"
	"k8s.io/client-go/rest"
	reg_util "kmodules.xyz/client-go/admissionregistration/v1beta1"
	core_util "kmodules.xyz/client-go/core/v1"
	"kmodules.xyz/client-go/discovery"
	appcat_cs "kmodules.xyz/custom-resources/client/clientset/versioned"
)

type OperatorConfig struct {
	amc.Config

	LicenseFile      string
	ClientConfig     *rest.Config
	KubeClient       kubernetes.Interface
	CRDClient        crd_cs.Interface
	DBClient         cs.Interface
	DynamicClient    dynamic.Interface
	AppCatalogClient appcat_cs.Interface
	PromClient       pcm.MonitoringV1Interface
}

func NewOperatorConfig(clientConfig *rest.Config) *OperatorConfig {
	return &OperatorConfig{
		ClientConfig: clientConfig,
	}
}

func (c *OperatorConfig) New() (*Controller, error) {
	if err := discovery.IsDefaultSupportedVersion(c.KubeClient); err != nil {
		return nil, err
	}

	topology, err := core_util.DetectTopology(context.TODO(), metadata.NewForConfigOrDie(c.ClientConfig))
	if err != nil {
		return nil, err
	}

	recorder := eventer.NewEventRecorder(c.KubeClient, "Autoscaler operator")

	ctrl := New(
		c.ClientConfig,
		c.KubeClient,
		c.CRDClient,
		c.DBClient,
		c.DynamicClient,
		c.AppCatalogClient,
		c.PromClient,
		c.Config,
		topology,
		recorder,
	)

	if err := ctrl.EnsureCustomResourceDefinitions(); err != nil {
		return nil, err
	}
	if c.EnableMutatingWebhook {
		if err := reg_util.UpdateMutatingWebhookCABundle(c.ClientConfig, kubedb.MutatorGroupName); err != nil {
			return nil, err
		}
	}
	if c.EnableValidatingWebhook {
		if err := reg_util.UpdateValidatingWebhookCABundle(c.ClientConfig, kubedb.ValidatorGroupName); err != nil {
			return nil, err
		}
	}

	if err := ctrl.Init(); err != nil {
		return nil, err
	}

	return ctrl, nil
}
