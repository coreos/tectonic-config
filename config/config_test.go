package config

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	kubecore "github.com/coreos/tectonic-config/config/kube-core"
)

var config = &kubecore.OperatorConfig{
	TypeMeta: metav1.TypeMeta{
		Kind:       kubecore.Kind,
		APIVersion: kubecore.APIVersion,
	},
	ClusterConfig: kubecore.ClusterConfig{
		APIServerURL: "https://kco-test-api.coreservices.team.coreos.systems",
	},
	AuthConfig: kubecore.AuthConfig{
		OIDCClientID:      "tectonic-kubectl",
		OIDCIssuerURL:     "https://kco-test.coreservices.team.coreos.systems/identity",
		OIDCGroupsClaim:   "groups",
		OIDCUsernameClaim: "email",
	},
	DNSConfig: kubecore.DNSConfig{
		ClusterIP: "10.3.0.10",
	},
	RoutingConfig: kubecore.RoutingConfig{
		Subdomain: "kco-test.coreservices.team.coreos.systems",
	},
	CloudProviderConfig: kubecore.CloudProviderConfig{
		CloudConfigPath:      "/cloud/config/path",
		CloudProviderProfile: "aws",
	},
	NetworkConfig: kubecore.NetworkConfig{
		AdvertiseAddress: "0.0.0.0",
		ClusterCIDR:      "10.2.0.0/16",
		EtcdServers:      "https://kco-test-etcd-0.coreservices.team.coreos.systems:2379",
		ServiceCIDR:      "10.3.0.0/16",
	},
}

func TestToFile(t *testing.T) {
	got := &kubecore.OperatorConfig{}
	err := FromFile("testdata/kco-config.yaml", got)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(config, got) {
		t.Fatalf("Expected unmarshaled config to look like:\n\n%#v\n\ngot:\n\n%#v", config, got)
	}
}
func TestFromFile(t *testing.T) {
	dir, err := ioutil.TempDir("", "config-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	path := filepath.Join(dir, "kco-config.yaml")

	if err := ToFile(path, config); err != nil {
		t.Fatal(err)
	}
	got, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	want, err := ioutil.ReadFile("testdata/kco-config.yaml")
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(got, want) {
		t.Fatalf("Expected config to look like:\n\n%s\n\ngot:\n\n%s", want, got)
	}
}
