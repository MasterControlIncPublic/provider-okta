/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/healthcarecom/provider-okta/apis/admin/v1alpha1"
	v1alpha1app "github.com/healthcarecom/provider-okta/apis/app/v1alpha1"
	v1alpha1auth "github.com/healthcarecom/provider-okta/apis/auth/v1alpha1"
	v1alpha1group "github.com/healthcarecom/provider-okta/apis/group/v1alpha1"
	v1alpha1groupmembership "github.com/healthcarecom/provider-okta/apis/groupmembership/v1alpha1"
	v1alpha1idp "github.com/healthcarecom/provider-okta/apis/idp/v1alpha1"
	v1alpha1okta "github.com/healthcarecom/provider-okta/apis/okta/v1alpha1"
	v1alpha1trustedorigin "github.com/healthcarecom/provider-okta/apis/trustedorigin/v1alpha1"
	v1alpha1user "github.com/healthcarecom/provider-okta/apis/user/v1alpha1"
	v1alpha1apis "github.com/healthcarecom/provider-okta/apis/v1alpha1"
	v1beta1 "github.com/healthcarecom/provider-okta/apis/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1app.SchemeBuilder.AddToScheme,
		v1alpha1auth.SchemeBuilder.AddToScheme,
		v1alpha1group.SchemeBuilder.AddToScheme,
		v1alpha1groupmembership.SchemeBuilder.AddToScheme,
		v1alpha1idp.SchemeBuilder.AddToScheme,
		v1alpha1okta.SchemeBuilder.AddToScheme,
		v1alpha1trustedorigin.SchemeBuilder.AddToScheme,
		v1alpha1user.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
