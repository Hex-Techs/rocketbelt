/*
Copyright 2021 The.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// HexRoleBindingSpec defines the desired state of HexRoleBinding
type HexRoleBindingSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	RoleRef RoleRef  `json:"roleRef,omitempty"`
	Users   []string `json:"users,omitempty"`
}

type RoleRef struct {
	Name string `json:"name,omitempty"`
}

// HexRoleBindingStatus defines the observed state of HexRoleBinding
// type HexRoleBindingStatus struct {
// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
// Important: Run "make" to regenerate code after modifying this file
// }

// +genclient
// +genclient:nonNamespaced
//+kubebuilder:object:root=true

// HexRoleBinding is the Schema for the hexrolebindings API
type HexRoleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec HexRoleBindingSpec `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// HexRoleBindingList contains a list of HexRoleBinding
type HexRoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HexRoleBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HexRoleBinding{}, &HexRoleBindingList{})
}
