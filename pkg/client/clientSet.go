package client

import (
	"k8s.io/client-go/rest"

	v1 "github.com/cloudogu/k8s-debug-mode-cr-lib/pkg/client/v1"
)

// DebugModeEcosystemInterface exposes the clients for all the custom resources of this library.
type DebugModeEcosystemInterface interface {
	DebugModeV1() v1.DebugModeV1Interface
}

type clientSet struct {
	clientV1 v1.DebugModeV1Interface
}

// NewDebugModeClientSet creates a new instance of the debug mode client set.
func NewDebugModeClientSet(config *rest.Config) (DebugModeEcosystemInterface, error) {
	clientV1, err := v1.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &clientSet{
		clientV1: clientV1,
	}, nil
}

// DebugModeV1 returns the debug mode v1 client.
func (cswc *clientSet) DebugModeV1() v1.DebugModeV1Interface {
	return cswc.clientV1
}
