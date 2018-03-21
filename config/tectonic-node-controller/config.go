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
	metav1.TypeMeta       `json:",inline"`
	HTTPProxy             string `json:"HTTPProxy"`
	HTTPSProxy            string `json:"HTTPSProxy"`
	NoProxy               string `json:"NoProxy"`
	KubeletImageURL       string `json:"KubeletImageURL"`
	KubeletImageTag       string `json:"KubeletImageTag"`
	IscsiEnabled          string `json:"IscsiEnabled"`
	KubeconfigFetchCmd    string `json:"KubeconfigFetchCmd"`
	TectonicTorcxImageURL string `json:"TectonicTorcxImageURL"`
	TectonicTorcxImageTag string `json:"TectonicTorcxImageTag"`
	BootstrapUpgradeCl    string `json:"BootstrapUpgradeCl"`
	TorcxStoreURL         string `json:"TorcxStoreURL"`
	TorcxSkipSetup        string `json:"TorcxSkipSetup"`
	MasterNodeLabel       string `json:"MasterNodeLabel"`
	WorkerNodeLabel       string `json:"WorkerNodeLabel"`
	NodeTaintsParam       string `json:"NodeTaintsParam"`
	ClusterDNSIP          string `json:"ClusterDNSIP"`
	CloudProvider         string `json:"CloudProvider"`
	CloudProviderConfig   string `json:"CloudProviderConfig"`
	DebugConfig           string `json:"DebugConfig"`
	ClusterName           string `json:"ClusterName"`

	// User customizations, list of node configs in the cluster to collect and apply as part of the final config
	AdditionalConfigs []string `json:"AddtionalConfigs"`
}
