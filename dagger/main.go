// A generated module for DaggerDebug functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/dagger-debug/internal/dagger"
	"fmt"
)

type DaggerDebug struct{}

func (m *DaggerDebug) Debug(
	ctx context.Context,
	src *dagger.Directory,
) error {
	fmt.Println("🔨 Building gameserver...")

	// List files in src
	entries, err := src.Entries(ctx)
	if err != nil {
		return fmt.Errorf("failed to list directory contents: %w", err)
	}

	fmt.Println("⭐️ Contents of src directory:")
	for _, entry := range entries {
		fmt.Printf("- %s\n", entry)
	}

	return nil
}

// // Returns a container that echoes whatever string argument is provided
// func (m *DaggerDebug) ContainerEcho(stringArg string) *dagger.Container {
// 	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
// }

// // Returns lines that match a pattern in the files of the provided Directory
// func (m *DaggerDebug) GrepDir(ctx context.Context, directoryArg *dagger.Directory, pattern string) (string, error) {
// 	return dag.Container().
// 		From("alpine:latest").
// 		WithMountedDirectory("/mnt", directoryArg).
// 		WithWorkdir("/mnt").
// 		WithExec([]string{"grep", "-R", pattern, "."}).
// 		Stdout(ctx)
// }
