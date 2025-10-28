// A module to test platform-builder

package main

import (
	"dagger/test/internal/dagger"
)

type Test struct{}

// Tests the platform-builder in k3s
func (m *Test) Test() *dagger.Container {
	// 1. setup k3s as a service
	// 2. run the platform-builder against k3s
	return dag.Container().From("alpine:latest")
}
