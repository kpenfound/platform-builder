// A module to test platform-builder

package main

import (
	"context"
)

type Test struct{}

// Tests the platform-builder in k3s
func (m *Test) Test(ctx context.Context) (string, error) {
	// 1. setup k3s as a service
	k3s := dag.K3S("test")
	kServer := k3s.Server()

	kServer, err := kServer.Start(ctx)
	if err != nil {
		return "", err
	}

	// 2. run the platform-builder against k3s
	return dag.PlatformBuilder(k3s.Config()).CheckConfig(ctx)
}
