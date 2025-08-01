package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"

	v1 "github.com/cloudogu/k8s-debug-mode-cr-lib/api/v1"
)

// client wraps the rest.Interface to use as a restClient for the component client.
type client struct {
	restClient rest.Interface
}

// NewForConfig creates a new client for a given rest.Config.
func NewForConfig(c *rest.Config) (DebugModeV1Interface, error) {
	config := *c
	gv := schema.GroupVersion{Group: v1.GroupVersion.Group, Version: v1.GroupVersion.Version}
	config.ContentConfig.GroupVersion = &gv
	config.APIPath = "/apis"

	s := scheme.Scheme
	err := v1.AddToScheme(s)
	if err != nil {
		return nil, err
	}

	metav1.AddToGroupVersion(s, gv)
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	restClient, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &client{restClient: restClient}, nil
}

// DebugMode takes a namespace and returns a debugMode client.
func (c *client) DebugMode(namespace string) DebugModeInterface {
	return &debugModeClient{
		client: c.restClient,
		ns:     namespace,
	}
}
