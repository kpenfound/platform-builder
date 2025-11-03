# platform-builder

A utility for platform builders to provide easy installation and management of platform components.

## Dagger Functions

```
Name                   Description
check-config           Checks the configuration of the platform
configure-repository   Configure repository for a component
get-pods               List all pods in the platform
install                Install a component
status                 Status of a component
uninstall              Uninstall a component
upgrade                Upgrade a component
```

## Components

Platform components available to platform builders.

### Platform Interface

The platform module defines the interface that all platform components must implement.

A component has two functions:

- `Install`: install the component into a cluster.
- `Status`: get the status of the component in the cluster.
- `ConfigureRepository`: configure the repository for the component.
- `Upgrade`: upgrade the component in the cluster.
- `Uninstall`: uninstall the component from a cluster.

### GitOps

The GitOps component installs ArgoCD.

### Creating a new component

1. First create a Dagger module in a subdirectory:

To create a new component called Foo, from the project root run:

```
dagger init --sdk go foo
```

2. Implement the interface functions in the module's main.go. Use the Argo CD implementation in `./argocd/main.go` as a reference.

3. Implement the component in the main module in `./main.go`. The PlatformBuilder constructor should create an instance of the new component and converted into the interface with `.AsPlatformComponent()`.

4. Install the new component in the Test function in `test/main.go`.

## Utility Modules

These utility modules can be used by components to interact with the platform.

### Helm

The Helm module is a utility for installing Helm charts in a Kubernetes cluster.

### Kubectl

The Kubectl module is a utility for interacting with Kubernetes clusters via kubectl.

## Testing

To test, from the project root run:

```
dagger -m test call test
```

This test sets up a k3s cluster and installs all available components and verifies their status.
