// A module to test platform-builder

package main

import (
	"context"
	"fmt"
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
	platformBuilder := dag.PlatformBuilder(k3s.Config())
	// Install GitOps
	// dagger call install --name gitops
	_, err = platformBuilder.Install(ctx, "gitops")
	if err != nil {
		return "", err
	}

	// Get Pods
	// dagger call get-pods
	pods, err := platformBuilder.GetPods(ctx)
	if err != nil {
		return "", err
	}

	// Status
	// dagger call status
	status, err := platformBuilder.Status(ctx)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("PODS:\n%s\n\nSTATUS:\n%s\n", pods, status), nil
}
