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

type BackupSnapshotExportJob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupSnapshotExportJobSpec   `json:"spec,omitempty"`
	Status            BackupSnapshotExportJobStatus `json:"status,omitempty"`
}

type BackupSnapshotExportJobSpecComponents struct {
	// +optional
	ExportID *string `json:"exportID,omitempty" tf:"export_id"`
	// +optional
	ReplicaSetName *string `json:"replicaSetName,omitempty" tf:"replica_set_name"`
}

type BackupSnapshotExportJobSpecCustomData struct {
	Key   *string `json:"key" tf:"key"`
	Value *string `json:"value" tf:"value"`
}

type BackupSnapshotExportJobSpec struct {
	State *BackupSnapshotExportJobSpecResource `json:"state,omitempty" tf:"-"`

	Resource BackupSnapshotExportJobSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type BackupSnapshotExportJobSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ClusterName *string `json:"clusterName" tf:"cluster_name"`
	// +optional
	Components []BackupSnapshotExportJobSpecComponents `json:"components,omitempty" tf:"components"`
	// +optional
	CreatedAt  *string                                 `json:"createdAt,omitempty" tf:"created_at"`
	CustomData []BackupSnapshotExportJobSpecCustomData `json:"customData" tf:"custom_data"`
	// +optional
	ErrMsg         *string `json:"errMsg,omitempty" tf:"err_msg"`
	ExportBucketID *string `json:"exportBucketID" tf:"export_bucket_id"`
	// +optional
	ExportJobID *string `json:"exportJobID,omitempty" tf:"export_job_id"`
	// +optional
	ExportStatusExportedCollections *int64 `json:"exportStatusExportedCollections,omitempty" tf:"export_status_exported_collections"`
	// +optional
	ExportStatusTotalCollections *int64 `json:"exportStatusTotalCollections,omitempty" tf:"export_status_total_collections"`
	// +optional
	FinishedAt *string `json:"finishedAt,omitempty" tf:"finished_at"`
	// +optional
	Prefix     *string `json:"prefix,omitempty" tf:"prefix"`
	ProjectID  *string `json:"projectID" tf:"project_id"`
	SnapshotID *string `json:"snapshotID" tf:"snapshot_id"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
}

type BackupSnapshotExportJobStatus struct {
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

// BackupSnapshotExportJobList is a list of BackupSnapshotExportJobs
type BackupSnapshotExportJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BackupSnapshotExportJob CRD objects
	Items []BackupSnapshotExportJob `json:"items,omitempty"`
}