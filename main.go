// A module for PlatformBuilder functions

package main

import (
	"context"
	"dagger/platform-builder/internal/dagger"
)

type PlatformBuilder struct {
	// +private
	Kubeconfig *dagger.File
	// +private
	GitOps *dagger.PlatformComponent
}

func New(
	kubeconfig *dagger.File,
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
		WithExec([]string{"kubectl", "config", "view"}).
		CombinedOutput(ctx)
}

// List all pods in the platform
func (m *PlatformBuilder) GetPods(ctx context.Context) (string, error) {
	return dag.Kubectl(m.Kubeconfig).
		Container().
		WithExec([]string{"kubectl", "get", "pods", "--all-namespaces", "-o", "wide"}).
		CombinedOutput(ctx)
}

// Install GitOps
func (m *PlatformBuilder) InstallGitOps(ctx context.Context) (string, error) {
	return m.GitOps.Install(ctx, m.Kubeconfig)
}

// GitOps Status
func (m *PlatformBuilder) StatusGitOps(ctx context.Context) (string, error) {
	return m.GitOps.Status(ctx, m.Kubeconfig)
}
