package tectonicingress

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// Kind is the TypeMeta.Kind for the OperatorConfig.
	Kind = "TectonicIngressOperatorConfig"
	// APIVersion is the TypeMeta.APIVersion for the OperatorConfig.
	APIVersion = "v1"
)

// OperatorConfig defines the configuration needed by the Tectonic Network Operator.
type OperatorConfig struct {
	metav1.TypeMeta `json:",inline"`

	// The installer platform, e.g. "bare-metal"
	InstallerPlatform string `json:"installerPlatform"`
}
