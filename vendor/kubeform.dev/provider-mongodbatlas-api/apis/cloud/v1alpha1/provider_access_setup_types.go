/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type ProviderAccessSetup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProviderAccessSetupSpec   `json:"spec,omitempty"`
	Status            ProviderAccessSetupStatus `json:"status,omitempty"`
}

type ProviderAccessSetupSpecAwsConfig struct {
	// +optional
	AtlasAssumedRoleExternalID *string `json:"atlasAssumedRoleExternalID,omitempty" tf:"atlas_assumed_role_external_id"`
	// +optional
	AtlasAwsAccountArn *string `json:"atlasAwsAccountArn,omitempty" tf:"atlas_aws_account_arn"`
}

type ProviderAccessSetupSpec struct {
	State *ProviderAccessSetupSpecResource `json:"state,omitempty" tf:"-"`

	Resource ProviderAccessSetupSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ProviderAccessSetupSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// Deprecated
	Aws *map[string]string `json:"aws,omitempty" tf:"aws"`
	// +optional
	AwsConfig []ProviderAccessSetupSpecAwsConfig `json:"awsConfig,omitempty" tf:"aws_config"`
	// +optional
	CreatedDate  *string `json:"createdDate,omitempty" tf:"created_date"`
	ProjectID    *string `json:"projectID" tf:"project_id"`
	ProviderName *string `json:"providerName" tf:"provider_name"`
	// +optional
	RoleID *string `json:"roleID,omitempty" tf:"role_id"`
}

type ProviderAccessSetupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ProviderAccessSetupList is a list of ProviderAccessSetups
type ProviderAccessSetupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ProviderAccessSetup CRD objects
	Items []ProviderAccessSetup `json:"items,omitempty"`
}
