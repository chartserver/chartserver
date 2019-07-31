/*
Copyright 2019 Replicated, Inc..

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ChartVersionStatus defines the observed state of ChartVersion
type ChartVersionStatus struct {
}

// ChartVersionSpec defines the desired state of ChartVersion
type ChartVersionSpec struct {
	AppVersion   string   `json:"appVersion"`
	ChartVersion string   `json:"chartVersion"`
	Created      string   `json:"created"`
	Description  string   `json:"description"`
	Digest       string   `json:"digest"`
	Home         string   `json:"home"`
	Icon         string   `json:"icon"`
	Maintainers  []string `json:"maintainers"`
	Name         string   `json:"name"`
	Sources      []string `json:"sources"`
	URLs         []string `json:"urls"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ChartVersion is the Schema for the ChartVersion API
// +k8s:openapi-gen=true
type ChartVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ChartVersionSpec   `json:"spec,omitempty"`
	Status ChartVersionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ChartVersionList contains a list of ChartVersion
type ChartVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChartVersion `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ChartVersion{}, &ChartVersionList{})
}
