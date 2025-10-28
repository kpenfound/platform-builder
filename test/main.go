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
	gitops, err := platformBuilder.
		InstallGitOps(ctx)
	if err != nil {
		return "", err
	}

	// Check Config
	config, err := platformBuilder.
		CheckConfig(ctx)
	if err != nil {
		return "", err
	}

	// Get Pods
	pods, err := platformBuilder.
		GetPods(ctx)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("GITOPS:\n%s\n\nCONFIG:\n%s\n\nPODS:\n%s\n", gitops, config, pods), nil
}
