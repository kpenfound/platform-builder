// A module for Kubectl functions

package main

import (
	"dagger/kubectl/internal/dagger"
)

type Kubectl struct {
	Container *dagger.Container
}

func New(kubeconfig *dagger.File) *Kubectl {
	container := dag.Container().
		From("alpine:3").
		WithExec([]string{"apk", "add", "kubectl"}).
		WithFile("/.kube/config", kubeconfig).
		WithEnvVariable("KUBECONFIG", "/.kube/config")

	return &Kubectl{
		Container: container,
	}
}
