// pkg/apis/mygroup.example.com/v1alpha1/doc.go
// +k8s:deepcopy-gen=package
package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MyResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MyResourceSpec `json:"spec"`
}
type MyResourceSpec struct {
	Image  string            `json:"image"`
	Memory resource.Quantity `json:"memory"`
}
type MyResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MyResource `json:"items"`
}
