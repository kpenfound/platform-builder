// An interface for platform-builder components

package main

import (
	"context"
	"dagger/platform/internal/dagger"
)

type Platform struct {
	kubeconfig *dagger.File
}

type Component interface {
	dagger.DaggerObject
	Install(ctx context.Context, kubeconfig *dagger.File) (string, error)
	Status(ctx context.Context, kubeconfig *dagger.File) (string, error)
	ConfigureRepository(ctx context.Context, kubeconfig *dagger.File, gitRepo string) (string, error)
	Upgrade(ctx context.Context, kubeconfig *dagger.File) (string, error)
	Uninstall(ctx context.Context, kubeconfig *dagger.File) (string, error)
}

// Install platform component into a cluster
func (pc *Platform) Install(ctx context.Context, component Component) (string, error) {
	return component.Install(ctx, pc.kubeconfig)
}

// Check the status of platform component in the cluster
func (pc *Platform) Status(ctx context.Context, component Component) (string, error) {
	return component.Status(ctx, pc.kubeconfig)
}

// Upgrade platform component in the cluster
func (pc *Platform) Upgrade(ctx context.Context, component Component) (string, error) {
	return component.Upgrade(ctx, pc.kubeconfig)
}

// Uninstall platform component from the cluster
func (pc *Platform) Uninstall(ctx context.Context, component Component) (string, error) {
	return component.Uninstall(ctx, pc.kubeconfig)
}

// Configure platform component in the cluster
func (pc *Platform) ConfigureRepository(ctx context.Context, component Component, config string) (string, error) {
	return component.ConfigureRepository(ctx, pc.kubeconfig, config)
}
