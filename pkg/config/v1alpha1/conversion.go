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
	"kubedb.dev/autoscaler/pkg/config"

	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/kube-controller-manager/config/v1alpha1"
)

// Important! The public back-and-forth conversion functions for the types in this package
// with HPAControllerConfiguration types need to be manually exposed like this in order for
// other packages that reference this package to be able to call these conversion functions
// in an autogenerated manner.
// TODO: Fix the bug in conversion-gen so it automatically discovers these Convert_* functions
// in autogenerated code as well.

// Convert_config_HPAControllerConfiguration_To_v1alpha1_HPAControllerConfiguration is an autogenerated conversion function.
func Convert_config_HPAControllerConfiguration_To_v1alpha1_HPAControllerConfiguration(in *config.HPAControllerConfiguration, out *v1alpha1.HPAControllerConfiguration, s conversion.Scope) error {
	return autoConvert_config_HPAControllerConfiguration_To_v1alpha1_HPAControllerConfiguration(in, out, s)
}

// Convert_v1alpha1_HPAControllerConfiguration_To_config_HPAControllerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_HPAControllerConfiguration_To_config_HPAControllerConfiguration(in *v1alpha1.HPAControllerConfiguration, out *config.HPAControllerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_HPAControllerConfiguration_To_config_HPAControllerConfiguration(in, out, s)
}
