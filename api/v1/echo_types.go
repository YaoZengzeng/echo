/*


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
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// EchoSpec defines the desired state of Echo
type EchoSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	APIVersion string `json:"apiVersion"`

	Kind string `json:"kind"`

	Name string `json:"name"`

	// Not all objects are required to be scoped to a namespace - the value of this field for
	// those objects will be empty.
	Namespace string `json:"namespace,omitempty"`

	RefPath string `json:"refPath"`
}

// EchoStatus defines the observed state of Echo
type EchoStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Data v1beta1.JSON `json:"data"`
}

// +kubebuilder:object:root=true

// Echo is the Schema for the echoes API
// +kubebuilder:subresource:status
type Echo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EchoSpec   `json:"spec,omitempty"`
	Status EchoStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EchoList contains a list of Echo
type EchoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Echo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Echo{}, &EchoList{})
}
