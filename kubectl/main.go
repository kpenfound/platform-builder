// A module for Kubectl functions

package main

import (
	"dagger/kubectl/internal/dagger"
	"fmt"
	"runtime"
)

type Kubectl struct {
	Container *dagger.Container
}

func New(kubeconfig *dagger.Secret) *Kubectl {
	installer := fmt.Sprintf("curl -LO \"https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/%s/kubectl\"", runtime.GOARCH)
	bin := dag.Container().
		From("alpine:3").
		WithExec([]string{"sh", "-c", installer}).
		File("kubectl")

	container := dag.Container().
		From("alpine:3").
		WithFile("/bin/kubectl", bin).
		WithMountedSecret("/root/.kube/config", kubeconfig)

	return &Kubectl{
		Container: container,
	}
}
