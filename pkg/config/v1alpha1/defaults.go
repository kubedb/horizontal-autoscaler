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

package v1alpha1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubectrlmgrconfigv1alpha1 "k8s.io/kube-controller-manager/config/v1alpha1"
	utilpointer "k8s.io/utils/pointer"
)

// RecommendedDefaultHPAControllerConfiguration defaults a pointer to a
// HPAControllerConfiguration struct. This will set the recommended default
// values, but they may be subject to change between API versions. This function
// is intentionally not registered in the scheme as a "normal" `SetDefaults_Foo`
// function to allow consumers of this type to set whatever defaults for their
// embedded configs. Forcing consumers to use these defaults would be problematic
// as defaulting in the scheme is done as part of the conversion, and there would
// be no easy way to opt-out. Instead, if you want to use this defaulting method
// run it in your wrapper struct of this type in its `SetDefaults_` method.
func RecommendedDefaultHPAControllerConfiguration(obj *kubectrlmgrconfigv1alpha1.HPAControllerConfiguration) {
	zero := metav1.Duration{}
	if obj.HorizontalPodAutoscalerUseRESTClients == nil {
		obj.HorizontalPodAutoscalerUseRESTClients = utilpointer.BoolPtr(true)
	}
	if obj.HorizontalPodAutoscalerSyncPeriod == zero {
		obj.HorizontalPodAutoscalerSyncPeriod = metav1.Duration{Duration: 15 * time.Second}
	}
	if obj.HorizontalPodAutoscalerUpscaleForbiddenWindow == zero {
		obj.HorizontalPodAutoscalerUpscaleForbiddenWindow = metav1.Duration{Duration: 3 * time.Minute}
	}
	if obj.HorizontalPodAutoscalerDownscaleStabilizationWindow == zero {
		obj.HorizontalPodAutoscalerDownscaleStabilizationWindow = metav1.Duration{Duration: 5 * time.Minute}
	}
	if obj.HorizontalPodAutoscalerCPUInitializationPeriod == zero {
		obj.HorizontalPodAutoscalerCPUInitializationPeriod = metav1.Duration{Duration: 5 * time.Minute}
	}
	if obj.HorizontalPodAutoscalerInitialReadinessDelay == zero {
		obj.HorizontalPodAutoscalerInitialReadinessDelay = metav1.Duration{Duration: 30 * time.Second}
	}
	if obj.HorizontalPodAutoscalerDownscaleForbiddenWindow == zero {
		obj.HorizontalPodAutoscalerDownscaleForbiddenWindow = metav1.Duration{Duration: 5 * time.Minute}
	}
	if obj.HorizontalPodAutoscalerTolerance == 0 {
		obj.HorizontalPodAutoscalerTolerance = 0.1
	}
}
