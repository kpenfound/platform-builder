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

Platform components available to platform builders

### Platform Interface

The platform module defines the interface that all platform components must implement.

A component has two functions:

- `Install`: install the component into a cluster
- `Status`: get the status of the component in the cluster

### GitOps

The GitOps component installs ArgoCD.

## Utility Modules

These utility modules can be used by components to interact with the platform
### Helm

The Helm module is a utility for installing Helm charts in a Kubernetes cluster.

### Kubectl

The Kubectl module is a utility for interacting with Kubernetes clusters via kubectl.

## Testing

```
dagger -m test call test
```

This test sets up a k3s cluster and installs all available components and verifies their status.
