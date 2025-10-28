// A generated module for Helm functions

package main

import (
	"context"
	"dagger/helm/internal/dagger"
)

type Helm struct {
	Kubeconfig *dagger.File
}

func New(kubeconfig *dagger.File) *Helm {
	return &Helm{
		Kubeconfig: kubeconfig,
	}
}

// Returns lines that match a pattern in the files of the provided Directory
func (m *Helm) UpgradeInstall(ctx context.Context, name, chart string) (string, error) {
	return dag.Container().From("alpine/helm").
		WithExec([]string{"apk", "add", "kubectl"}).
		WithEnvVariable("KUBECONFIG", "/.kube/config").
		WithFile("/.kube/config", m.Kubeconfig).
		WithExec([]string{"helm", "upgrade", "--install", "--force", "--wait", "--debug", name, chart}).
		Stdout(ctx)
}
