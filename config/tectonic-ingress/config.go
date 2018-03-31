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

const (
	// TypeHAProxy is haproxy based openshift-router.
	TypeHAProxy = "haproxy-router"
)

// OperatorConfig defines the configuration needed by the Tectonic Network Operator.
type OperatorConfig struct {
	metav1.TypeMeta `json:",inline"`

	Type           string `json:"type"`
	StatsUsername  string `json:"statsUsername"`
	StatsPassword  string `json:"statsPassword"`
	ForceSubdomain string `json:"forceSubdomain"`
}
