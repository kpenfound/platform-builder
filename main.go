// A module for PlatformBuilder functions

package main

import (
	"context"
	"dagger/platform-builder/internal/dagger"
)

type PlatformBuilder struct {
	// +private
	Kubeconfig *dagger.Secret
	GitOps     *dagger.PlatformComponent
}

func New(
	// +optional
	kubeconfig *dagger.Secret,
) *PlatformBuilder {
	return &PlatformBuilder{
		Kubeconfig: kubeconfig,
		// GitOps = Argo CD
		GitOps: dag.Argocd().AsPlatformComponent(),
	}
}

// Checks the configuration of the platform
func (m *PlatformBuilder) CheckConfig(ctx context.Context) (string, error) {
	return dag.Kubectl(m.Kubeconfig).
		Container().
		WithExec([]string{"kubectl", "version"}).
		Stdout(ctx)
}
