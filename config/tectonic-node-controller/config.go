package tnc

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// Kind is the TypeMeta.Kind for the TNC config.
	Kind = "TectonicNodeControllerConfig"
	// APIVersion is the TypeMeta.APIVersion for the TNC config.
	APIVersion = "v1"
)

// ControllerConfig is the config structure TNC to generate NodeConfigs
type ControllerConfig struct {
	metav1.TypeMeta     `json:",inline"`
	KubeconfigFetchCmd  string `json:"KubeconfigFetchCmd"` // TODO(yifan): Try to remove this.
	ClusterDNSIP        string `json:"ClusterDNSIP"`
	CloudProvider       string `json:"CloudProvider"`
	CloudProviderConfig string `json:"CloudProviderConfig"`
	ClusterName         string `json:"ClusterName"`

	BaseDomain string `json:"BaseDomain"`

	// Etcd vars
	EtcdInitialCluster string `json:"EtcdInitialCluster"` // TODO(yifan): Calculate this based on 'BaseDomain'.

	// User customizations, list of node configs in the cluster to collect and apply as part of the final config
	AdditionalConfigs []string `json:"AddtionalConfigs"`
}
