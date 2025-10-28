// A module for PlatformBuilder functions

package main

import (
	"context"
	"dagger/platform-builder/internal/dagger"
)

type PlatformBuilder struct {
	// +private
	Kubeconfig *dagger.File
}

func New(
	// +optional
	kubeconfig *dagger.File,
) *PlatformBuilder {
	return &PlatformBuilder{
		Kubeconfig: kubeconfig,
	}
}

// Checks the configuration of the platform
func (m *PlatformBuilder) CheckConfig(ctx context.Context) (string, error) {
	return dag.Container().From("alpine:latest").Stdout(ctx)
}

// Installs GitOps into your cluster
func (m *PlatformBuilder) InstallGitOps(ctx context.Context) (string, error) {
	return dag.Container().From("alpine:latest").Stdout(ctx)
}

// Installs Foo into your cluster
func (m *PlatformBuilder) InstallFoo(ctx context.Context) (string, error) {
	return dag.Container().From("alpine:latest").Stdout(ctx)
}
