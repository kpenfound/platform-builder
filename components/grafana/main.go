// A module for managing Grafana installations

package main

import (
	"context"
	"dagger/grafana/internal/dagger"
)

type Grafana struct{}

const (
	NAME  = "grafana"
	CHART = "https://grafana.github.io/helm-charts/grafana"
)

// Installs Grafana into a cluster
func (m *Grafana) Install(ctx context.Context, kubeconfig *dagger.File) (string, error) {
	return dag.Helm(kubeconfig).UpgradeInstall(ctx, NAME, CHART)
}

// Checks the status of Grafana in a cluster
func (m *Grafana) Status(ctx context.Context, kubeconfig *dagger.File) (string, error) {
	return dag.Helm(kubeconfig).Status(ctx, NAME)
}

// Configures Grafana with a git repository
func (m *Grafana) ConfigureRepository(ctx context.Context, kubeconfig *dagger.File, gitRepo string) (string, error) {
	return "NOT IMPLEMENTED", nil
}

// Upgrade Grafana in a cluster
func (m *Grafana) Upgrade(ctx context.Context, kubeconfig *dagger.File) (string, error) {
	return dag.Helm(kubeconfig).UpgradeInstall(ctx, NAME, CHART)
}

// Uninstall Grafana from a cluster
func (m *Grafana) Uninstall(ctx context.Context, kubeconfig *dagger.File) (string, error) {
	return dag.Helm(kubeconfig).Uninstall(ctx, NAME)
}
