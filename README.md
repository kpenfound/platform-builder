# platform-builder

A utility for platform builders to provide easy installation and management of platform components.

## Dagger Functions

```
Name              Description
check-config      Checks the configuration of the platform
get-pods          List all pods in the platform
install-git-ops   Install GitOps
status-git-ops    GitOps Status
```

## Components

Platform components available to platform builders.

### Platform Interface

The platform module defines the interface that all platform components must implement.

A component has two functions:

- `Install`: install the component into a cluster.
- `Status`: get the status of the component in the cluster.

### GitOps

The GitOps component installs ArgoCD.

### Creating a new component

1. First create a Dagger module in a subdirectory:

To create a new component called Foo, from the project root run:

```
dagger init --sdk go foo
```

2. Implement the `Install` and `Status` functions in the module's main.go. Use the Argo CD module as a reference.

3. Implement the component in the main module in `./main.go`:

- The PlatformBuilder struct will need a field for an instance of the new component as a `*dagger.PlatformComponent` that must be decorated with `// +private`.
- The PlatformBuilder constructor should create an instance of the new component and converted into the interface with `.AsPlatformComponent()`.
- Create the corresponding `InstallFoo` and `StatusFoo` functions in the platform module for the new component.

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
