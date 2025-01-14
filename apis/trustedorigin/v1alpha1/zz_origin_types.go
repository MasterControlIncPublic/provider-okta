/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type OriginObservation struct {

	// Whether the Trusted Origin is active or not - can only be issued post-creation
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Unique name for this trusted origin
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Unique origin URL for this trusted origin
	Origin *string `json:"origin,omitempty" tf:"origin,omitempty"`

	// Scopes of the Trusted Origin - can either be CORS or REDIRECT only
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`
}

type OriginParameters struct {

	// Whether the Trusted Origin is active or not - can only be issued post-creation
	// +kubebuilder:validation:Optional
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// Unique name for this trusted origin
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Unique origin URL for this trusted origin
	// +kubebuilder:validation:Optional
	Origin *string `json:"origin,omitempty" tf:"origin,omitempty"`

	// Scopes of the Trusted Origin - can either be CORS or REDIRECT only
	// +kubebuilder:validation:Optional
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`
}

// OriginSpec defines the desired state of Origin
type OriginSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OriginParameters `json:"forProvider"`
}

// OriginStatus defines the observed state of Origin.
type OriginStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OriginObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Origin is the Schema for the Origins API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,okta}
type Origin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.origin)",message="origin is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.scopes)",message="scopes is a required parameter"
	Spec   OriginSpec   `json:"spec"`
	Status OriginStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OriginList contains a list of Origins
type OriginList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Origin `json:"items"`
}

// Repository type metadata.
var (
	Origin_Kind             = "Origin"
	Origin_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Origin_Kind}.String()
	Origin_KindAPIVersion   = Origin_Kind + "." + CRDGroupVersion.String()
	Origin_GroupVersionKind = CRDGroupVersion.WithKind(Origin_Kind)
)

func init() {
	SchemeBuilder.Register(&Origin{}, &OriginList{})
}
