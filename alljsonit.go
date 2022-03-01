/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package main

import (
	jsoniter "github.com/json-iterator/go"
	"k8s.io/apimachinery/pkg/runtime/schema"
	alertv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/alert/v1alpha1"
	auditingv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/auditing/v1alpha1"
	cloudv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/cloud/v1alpha1"
	clusterv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/cluster/v1alpha1"
	customv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/custom/v1alpha1"
	datav1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/data/v1alpha1"
	databasev1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/database/v1alpha1"
	encryptionv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/encryption/v1alpha1"
	eventv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/event/v1alpha1"
	globalv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/global/v1alpha1"
	ldapv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/ldap/v1alpha1"
	maintenancev1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/maintenance/v1alpha1"
	networkv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/network/v1alpha1"
	onlinev1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/online/v1alpha1"
	privatev1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/private/v1alpha1"
	privatelinkv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/privatelink/v1alpha1"
	projectv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/project/v1alpha1"
	searchv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/search/v1alpha1"
	teamv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/team/v1alpha1"
	teamsv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/teams/v1alpha1"
	thirdv1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/third/v1alpha1"
	x509v1alpha1 "kubeform.dev/provider-mongodbatlas-api/apis/x509/v1alpha1"
	"kubeform.dev/provider-mongodbatlas-controller/controllers"
)

type Data struct {
	JsonIt       jsoniter.API
	ResourceType string
}

var allJsonIt = map[schema.GroupVersionResource]Data{
	{
		Group:    "alert.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "configurations",
	}: {
		JsonIt:       controllers.GetJSONItr(alertv1alpha1.GetEncoder(), alertv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_alert_configuration",
	},
	{
		Group:    "auditing.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "auditings",
	}: {
		JsonIt:       controllers.GetJSONItr(auditingv1alpha1.GetEncoder(), auditingv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_auditing",
	},
	{
		Group:    "cloud.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "provideraccesses",
	}: {
		JsonIt:       controllers.GetJSONItr(cloudv1alpha1.GetEncoder(), cloudv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_cloud_provider_access",
	},
	{
		Group:    "cloud.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "provideraccessauthorizations",
	}: {
		JsonIt:       controllers.GetJSONItr(cloudv1alpha1.GetEncoder(), cloudv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_cloud_provider_access_authorization",
	},
	{
		Group:    "cloud.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "provideraccesssetups",
	}: {
		JsonIt:       controllers.GetJSONItr(cloudv1alpha1.GetEncoder(), cloudv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_cloud_provider_access_setup",
	},
	{
		Group:    "cloud.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "providersnapshots",
	}: {
		JsonIt:       controllers.GetJSONItr(cloudv1alpha1.GetEncoder(), cloudv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_cloud_provider_snapshot",
	},
	{
		Group:    "cloud.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "providersnapshotbackuppolicies",
	}: {
		JsonIt:       controllers.GetJSONItr(cloudv1alpha1.GetEncoder(), cloudv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_cloud_provider_snapshot_backup_policy",
	},
	{
		Group:    "cloud.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "providersnapshotrestorejobs",
	}: {
		JsonIt:       controllers.GetJSONItr(cloudv1alpha1.GetEncoder(), cloudv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_cloud_provider_snapshot_restore_job",
	},
	{
		Group:    "cluster.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "clusters",
	}: {
		JsonIt:       controllers.GetJSONItr(clusterv1alpha1.GetEncoder(), clusterv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_cluster",
	},
	{
		Group:    "custom.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "dbroles",
	}: {
		JsonIt:       controllers.GetJSONItr(customv1alpha1.GetEncoder(), customv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_custom_db_role",
	},
	{
		Group:    "custom.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "dnsconfigurationclusteraws",
	}: {
		JsonIt:       controllers.GetJSONItr(customv1alpha1.GetEncoder(), customv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_custom_dns_configuration_cluster_aws",
	},
	{
		Group:    "data.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "lakes",
	}: {
		JsonIt:       controllers.GetJSONItr(datav1alpha1.GetEncoder(), datav1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_data_lake",
	},
	{
		Group:    "database.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "users",
	}: {
		JsonIt:       controllers.GetJSONItr(databasev1alpha1.GetEncoder(), databasev1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_database_user",
	},
	{
		Group:    "encryption.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "atrests",
	}: {
		JsonIt:       controllers.GetJSONItr(encryptionv1alpha1.GetEncoder(), encryptionv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_encryption_at_rest",
	},
	{
		Group:    "event.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "triggers",
	}: {
		JsonIt:       controllers.GetJSONItr(eventv1alpha1.GetEncoder(), eventv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_event_trigger",
	},
	{
		Group:    "global.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "clusterconfigs",
	}: {
		JsonIt:       controllers.GetJSONItr(globalv1alpha1.GetEncoder(), globalv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_global_cluster_config",
	},
	{
		Group:    "ldap.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "configurations",
	}: {
		JsonIt:       controllers.GetJSONItr(ldapv1alpha1.GetEncoder(), ldapv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_ldap_configuration",
	},
	{
		Group:    "ldap.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "verifies",
	}: {
		JsonIt:       controllers.GetJSONItr(ldapv1alpha1.GetEncoder(), ldapv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_ldap_verify",
	},
	{
		Group:    "maintenance.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "windows",
	}: {
		JsonIt:       controllers.GetJSONItr(maintenancev1alpha1.GetEncoder(), maintenancev1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_maintenance_window",
	},
	{
		Group:    "network.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "containers",
	}: {
		JsonIt:       controllers.GetJSONItr(networkv1alpha1.GetEncoder(), networkv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_network_container",
	},
	{
		Group:    "network.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "peerings",
	}: {
		JsonIt:       controllers.GetJSONItr(networkv1alpha1.GetEncoder(), networkv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_network_peering",
	},
	{
		Group:    "online.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "archives",
	}: {
		JsonIt:       controllers.GetJSONItr(onlinev1alpha1.GetEncoder(), onlinev1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_online_archive",
	},
	{
		Group:    "private.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "ipmodes",
	}: {
		JsonIt:       controllers.GetJSONItr(privatev1alpha1.GetEncoder(), privatev1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_private_ip_mode",
	},
	{
		Group:    "privatelink.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "endpoints",
	}: {
		JsonIt:       controllers.GetJSONItr(privatelinkv1alpha1.GetEncoder(), privatelinkv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_privatelink_endpoint",
	},
	{
		Group:    "privatelink.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "endpointservices",
	}: {
		JsonIt:       controllers.GetJSONItr(privatelinkv1alpha1.GetEncoder(), privatelinkv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_privatelink_endpoint_service",
	},
	{
		Group:    "project.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "projects",
	}: {
		JsonIt:       controllers.GetJSONItr(projectv1alpha1.GetEncoder(), projectv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_project",
	},
	{
		Group:    "project.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "ipaccesslists",
	}: {
		JsonIt:       controllers.GetJSONItr(projectv1alpha1.GetEncoder(), projectv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_project_ip_access_list",
	},
	{
		Group:    "search.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "indices",
	}: {
		JsonIt:       controllers.GetJSONItr(searchv1alpha1.GetEncoder(), searchv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_search_index",
	},
	{
		Group:    "team.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "teams",
	}: {
		JsonIt:       controllers.GetJSONItr(teamv1alpha1.GetEncoder(), teamv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_team",
	},
	{
		Group:    "teams.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "teams",
	}: {
		JsonIt:       controllers.GetJSONItr(teamsv1alpha1.GetEncoder(), teamsv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_teams",
	},
	{
		Group:    "third.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "partyintegrations",
	}: {
		JsonIt:       controllers.GetJSONItr(thirdv1alpha1.GetEncoder(), thirdv1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_third_party_integration",
	},
	{
		Group:    "x509.mongodbatlas.kubeform.com",
		Version:  "v1alpha1",
		Resource: "authenticationdatabaseusers",
	}: {
		JsonIt:       controllers.GetJSONItr(x509v1alpha1.GetEncoder(), x509v1alpha1.GetDecoder()),
		ResourceType: "mongodbatlas_x509_authentication_database_user",
	},
}

func getJsonItAndResType(gvr schema.GroupVersionResource) Data {
	return allJsonIt[gvr]
}
