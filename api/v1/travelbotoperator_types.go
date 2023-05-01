/*
Copyright 2023.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TravelbotOperatorSpec defines the desired state of TravelbotOperator
type TravelbotOperatorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	ShipName      string `json:"shipName,omitempty"`
	ShipSpeed     string `json:"shipSpeed,omitempty"`
	SleepDuration string `json:"sleepDuration,omitempty"`
	DeployImage   string `json:"deployImage,omitempty"`
	Replicas      int32  `json:"replicas,omitempty"`
}

// TravelbotOperatorStatus defines the observed state of TravelbotOperator
type TravelbotOperatorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TravelbotOperator is the Schema for the travelbotoperators API
type TravelbotOperator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TravelbotOperatorSpec   `json:"spec,omitempty"`
	Status TravelbotOperatorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TravelbotOperatorList contains a list of TravelbotOperator
type TravelbotOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TravelbotOperator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TravelbotOperator{}, &TravelbotOperatorList{})
}
