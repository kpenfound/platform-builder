// A module for managing Argocd installations

package main

import (
	"context"
	"dagger/argocd/internal/dagger"
)

type Argocd struct{}

const (
	NAME  = "argocd"
	CHART = "oci://ghcr.io/argoproj/argo-helm/argo-cd"
)

// Installs Argocd into a cluster
func (m *Argocd) Install(ctx context.Context, kubeconfig *dagger.File) (string, error) {
	return dag.Helm(kubeconfig).UpgradeInstall(ctx, NAME, CHART)
}

// Checks the status of Argocd in a cluster
func (m *Argocd) Status(ctx context.Context, kubeconfig *dagger.File) (string, error) {
	return dag.Helm(kubeconfig).Status(ctx, NAME)
}

// Configures Argocd with a git repository
func (m *Argocd) ConfigureRepository(ctx context.Context, kubeconfig *dagger.File, gitRepo string) (string, error) {
	return "NOT IMPLEMENTED", nil
}

// Upgrade Argocd in a cluster
func (m *Argocd) Upgrade(ctx context.Context, kubeconfig *dagger.File) (string, error) {
	return dag.Helm(kubeconfig).UpgradeInstall(ctx, NAME, CHART)
}

// Uninstall Argocd from a cluster
func (m *Argocd) Uninstall(ctx context.Context, kubeconfig *dagger.File) (string, error) {
	return dag.Helm(kubeconfig).Uninstall(ctx, NAME)
}
