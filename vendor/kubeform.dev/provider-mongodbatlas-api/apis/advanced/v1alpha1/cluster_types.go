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

type Cluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec,omitempty"`
	Status            ClusterStatus `json:"status,omitempty"`
}

type ClusterSpecAdvancedConfiguration struct {
	// +optional
	DefaultReadConcern *string `json:"defaultReadConcern,omitempty" tf:"default_read_concern"`
	// +optional
	DefaultWriteConcern *string `json:"defaultWriteConcern,omitempty" tf:"default_write_concern"`
	// +optional
	FailIndexKeyTooLong *bool `json:"failIndexKeyTooLong,omitempty" tf:"fail_index_key_too_long"`
	// +optional
	JavascriptEnabled *bool `json:"javascriptEnabled,omitempty" tf:"javascript_enabled"`
	// +optional
	MinimumEnabledTlsProtocol *string `json:"minimumEnabledTlsProtocol,omitempty" tf:"minimum_enabled_tls_protocol"`
	// +optional
	NoTableScan *bool `json:"noTableScan,omitempty" tf:"no_table_scan"`
	// +optional
	OplogSizeMb *int64 `json:"oplogSizeMb,omitempty" tf:"oplog_size_mb"`
	// +optional
	SampleRefreshIntervalBiConnector *int64 `json:"sampleRefreshIntervalBiConnector,omitempty" tf:"sample_refresh_interval_bi_connector"`
	// +optional
	SampleSizeBiConnector *int64 `json:"sampleSizeBiConnector,omitempty" tf:"sample_size_bi_connector"`
}

type ClusterSpecBiConnector struct {
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	ReadPreference *string `json:"readPreference,omitempty" tf:"read_preference"`
}

type ClusterSpecConnectionStringsPrivateEndpointEndpoints struct {
	// +optional
	EndpointID *string `json:"endpointID,omitempty" tf:"endpoint_id"`
	// +optional
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name"`
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
}

type ClusterSpecConnectionStringsPrivateEndpoint struct {
	// +optional
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string"`
	// +optional
	Endpoints []ClusterSpecConnectionStringsPrivateEndpointEndpoints `json:"endpoints,omitempty" tf:"endpoints"`
	// +optional
	SrvConnectionString *string `json:"srvConnectionString,omitempty" tf:"srv_connection_string"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type ClusterSpecConnectionStrings struct {
	// +optional
	// Deprecated
	AwsPrivateLink map[string]string `json:"awsPrivateLink,omitempty" tf:"aws_private_link"`
	// +optional
	// Deprecated
	AwsPrivateLinkSrv map[string]string `json:"awsPrivateLinkSrv,omitempty" tf:"aws_private_link_srv"`
	// +optional
	Private *string `json:"private,omitempty" tf:"private"`
	// +optional
	PrivateEndpoint []ClusterSpecConnectionStringsPrivateEndpoint `json:"privateEndpoint,omitempty" tf:"private_endpoint"`
	// +optional
	PrivateSrv *string `json:"privateSrv,omitempty" tf:"private_srv"`
	// +optional
	Standard *string `json:"standard,omitempty" tf:"standard"`
	// +optional
	StandardSrv *string `json:"standardSrv,omitempty" tf:"standard_srv"`
}

type ClusterSpecLabels struct {
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type ClusterSpecReplicationSpecsRegionConfigsAnalyticsSpecs struct {
	// +optional
	DiskIops *int64 `json:"diskIops,omitempty" tf:"disk_iops"`
	// +optional
	EbsVolumeType *string `json:"ebsVolumeType,omitempty" tf:"ebs_volume_type"`
	InstanceSize  *string `json:"instanceSize" tf:"instance_size"`
	// +optional
	NodeCount *int64 `json:"nodeCount,omitempty" tf:"node_count"`
}

type ClusterSpecReplicationSpecsRegionConfigsAutoScaling struct {
	// +optional
	ComputeEnabled *bool `json:"computeEnabled,omitempty" tf:"compute_enabled"`
	// +optional
	ComputeMaxInstanceSize *string `json:"computeMaxInstanceSize,omitempty" tf:"compute_max_instance_size"`
	// +optional
	ComputeMinInstanceSize *string `json:"computeMinInstanceSize,omitempty" tf:"compute_min_instance_size"`
	// +optional
	ComputeScaleDownEnabled *bool `json:"computeScaleDownEnabled,omitempty" tf:"compute_scale_down_enabled"`
	// +optional
	DiskGbEnabled *bool `json:"diskGbEnabled,omitempty" tf:"disk_gb_enabled"`
}

type ClusterSpecReplicationSpecsRegionConfigsElectableSpecs struct {
	// +optional
	DiskIops *int64 `json:"diskIops,omitempty" tf:"disk_iops"`
	// +optional
	EbsVolumeType *string `json:"ebsVolumeType,omitempty" tf:"ebs_volume_type"`
	InstanceSize  *string `json:"instanceSize" tf:"instance_size"`
	// +optional
	NodeCount *int64 `json:"nodeCount,omitempty" tf:"node_count"`
}

type ClusterSpecReplicationSpecsRegionConfigsReadOnlySpecs struct {
	// +optional
	DiskIops *int64 `json:"diskIops,omitempty" tf:"disk_iops"`
	// +optional
	EbsVolumeType *string `json:"ebsVolumeType,omitempty" tf:"ebs_volume_type"`
	InstanceSize  *string `json:"instanceSize" tf:"instance_size"`
	// +optional
	NodeCount *int64 `json:"nodeCount,omitempty" tf:"node_count"`
}

type ClusterSpecReplicationSpecsRegionConfigs struct {
	// +optional
	AnalyticsSpecs *ClusterSpecReplicationSpecsRegionConfigsAnalyticsSpecs `json:"analyticsSpecs,omitempty" tf:"analytics_specs"`
	// +optional
	AutoScaling *ClusterSpecReplicationSpecsRegionConfigsAutoScaling `json:"autoScaling,omitempty" tf:"auto_scaling"`
	// +optional
	BackingProviderName *string `json:"backingProviderName,omitempty" tf:"backing_provider_name"`
	// +optional
	ElectableSpecs *ClusterSpecReplicationSpecsRegionConfigsElectableSpecs `json:"electableSpecs,omitempty" tf:"electable_specs"`
	Priority       *int64                                                  `json:"priority" tf:"priority"`
	ProviderName   *string                                                 `json:"providerName" tf:"provider_name"`
	// +optional
	ReadOnlySpecs *ClusterSpecReplicationSpecsRegionConfigsReadOnlySpecs `json:"readOnlySpecs,omitempty" tf:"read_only_specs"`
	RegionName    *string                                                `json:"regionName" tf:"region_name"`
}

type ClusterSpecReplicationSpecs struct {
	// +optional
	ContainerID *map[string]string `json:"containerID,omitempty" tf:"container_id"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	NumShards     *int64                                     `json:"numShards,omitempty" tf:"num_shards"`
	RegionConfigs []ClusterSpecReplicationSpecsRegionConfigs `json:"regionConfigs" tf:"region_configs"`
	// +optional
	ZoneName *string `json:"zoneName,omitempty" tf:"zone_name"`
}

type ClusterSpec struct {
	State *ClusterSpecResource `json:"state,omitempty" tf:"-"`

	Resource ClusterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ClusterSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AdvancedConfiguration *ClusterSpecAdvancedConfiguration `json:"advancedConfiguration,omitempty" tf:"advanced_configuration"`
	// +optional
	BackupEnabled *bool `json:"backupEnabled,omitempty" tf:"backup_enabled"`
	// +optional
	BiConnector *ClusterSpecBiConnector `json:"biConnector,omitempty" tf:"bi_connector"`
	// +optional
	ClusterID   *string `json:"clusterID,omitempty" tf:"cluster_id"`
	ClusterType *string `json:"clusterType" tf:"cluster_type"`
	// +optional
	ConnectionStrings []ClusterSpecConnectionStrings `json:"connectionStrings,omitempty" tf:"connection_strings"`
	// +optional
	CreateDate *string `json:"createDate,omitempty" tf:"create_date"`
	// +optional
	DiskSizeGb *float64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb"`
	// +optional
	EncryptionAtRestProvider *string `json:"encryptionAtRestProvider,omitempty" tf:"encryption_at_rest_provider"`
	// +optional
	Labels []ClusterSpecLabels `json:"labels,omitempty" tf:"labels"`
	// +optional
	MongoDbMajorVersion *string `json:"mongoDbMajorVersion,omitempty" tf:"mongo_db_major_version"`
	// +optional
	MongoDbVersion *string `json:"mongoDbVersion,omitempty" tf:"mongo_db_version"`
	Name           *string `json:"name" tf:"name"`
	// +optional
	Paused *bool `json:"paused,omitempty" tf:"paused"`
	// +optional
	PitEnabled       *bool                         `json:"pitEnabled,omitempty" tf:"pit_enabled"`
	ProjectID        *string                       `json:"projectID" tf:"project_id"`
	ReplicationSpecs []ClusterSpecReplicationSpecs `json:"replicationSpecs" tf:"replication_specs"`
	// +optional
	RootCertType *string `json:"rootCertType,omitempty" tf:"root_cert_type"`
	// +optional
	StateName *string `json:"stateName,omitempty" tf:"state_name"`
	// +optional
	VersionReleaseSystem *string `json:"versionReleaseSystem,omitempty" tf:"version_release_system"`
}

type ClusterStatus struct {
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

// ClusterList is a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Cluster CRD objects
	Items []Cluster `json:"items,omitempty"`
}
