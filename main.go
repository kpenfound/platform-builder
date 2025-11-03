// A module for PlatformBuilder functions

package main

import (
	"context"
	"dagger/platform-builder/internal/dagger"
	"fmt"
)

type PlatformBuilder struct {
	// +private
	Kubeconfig *dagger.File
	// +private
	Components []NamedComponent
}

type NamedComponent struct {
	Name      string
	Component *dagger.PlatformComponent
}

func New(
	kubeconfig *dagger.File,
) *PlatformBuilder {
	components := []NamedComponent{}

	// Add components here
	components = append(components, NamedComponent{
		Name:      "gitops",
		Component: dag.Argocd().AsPlatformComponent(),
	})

	return &PlatformBuilder{
		Kubeconfig: kubeconfig,
		Components: components,
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

// Install a component
func (m *PlatformBuilder) Install(ctx context.Context, name string) (string, error) {
	component, err := m.findComponent(name)
	if err != nil {
		return "", err
	}
	return component.Install(ctx, m.Kubeconfig)
}

// Status of a component
func (m *PlatformBuilder) Status(ctx context.Context, name string) (string, error) {
	component, err := m.findComponent(name)
	if err != nil {
		return "", err
	}
	return component.Status(ctx, m.Kubeconfig)
}

// Configure repository for a component
func (m *PlatformBuilder) ConfigureRepository(ctx context.Context, name, repository string) (string, error) {
	component, err := m.findComponent(name)
	if err != nil {
		return "", err
	}
	return component.ConfigureRepository(ctx, m.Kubeconfig, repository)
}

// Upgrade a component
func (m *PlatformBuilder) Upgrade(ctx context.Context, name string) (string, error) {
	component, err := m.findComponent(name)
	if err != nil {
		return "", err
	}
	return component.Upgrade(ctx, m.Kubeconfig)
}

// Uninstall a component
func (m *PlatformBuilder) Uninstall(ctx context.Context, name string) (string, error) {
	component, err := m.findComponent(name)
	if err != nil {
		return "", err
	}
	return component.Uninstall(ctx, m.Kubeconfig)
}

// find a component in Components
func (m *PlatformBuilder) findComponent(name string) (*dagger.PlatformComponent, error) {
	for _, component := range m.Components {
		if component.Name == name {
			return component.Component, nil
		}
	}
	return nil, fmt.Errorf("component %s not found", name)
}
