// A generated module for Helm functions

package main

import (
	"context"
	"dagger/helm/internal/dagger"
	"strings"
)

type Helm struct {
	Kubeconfig *dagger.File
}

func New(kubeconfig *dagger.File) *Helm {
	return &Helm{
		Kubeconfig: kubeconfig,
	}
}

// Install or upgrade a helm chart
func (m *Helm) UpgradeInstall(ctx context.Context, name, chart string) (string, error) {
	return m.base().
		WithExec([]string{"helm", "upgrade", "--install", "--force", "--wait", "--debug", name, chart}).
		Stdout(ctx)
}

// Uninstall a helm chart
func (m *Helm) Uninstall(ctx context.Context, name string) (string, error) {
	return m.base().
		WithExec([]string{"helm", "uninstall", name, "-o", "json"}).
		Stdout(ctx)
}

// Status of a helm chart
func (m *Helm) Status(ctx context.Context, name string) (string, error) {
	status, err := m.base().
		WithExec([]string{"sh", "-c", "helm status " + name + " -o json | jq .info.status"}).
		Stdout(ctx)
	if err != nil {
		return "", err
	}
	if status == "" {
		return "not installed", nil
	}
	return strings.ReplaceAll(strings.TrimSpace(status), "\"", ""), nil
}

// Base container for Helm operations
func (m *Helm) base() *dagger.Container {
	return dag.Container().From("alpine/helm").
		WithExec([]string{"apk", "add", "kubectl"}).
		WithExec([]string{"apk", "add", "jq"}).
		WithEnvVariable("KUBECONFIG", "/.kube/config").
		WithFile("/.kube/config", m.Kubeconfig)
}
