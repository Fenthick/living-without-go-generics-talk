//go:build tools

// Package tools allows pinning versions of development dependencies in go.mod file.
package tools

import (
	_ "github.com/cmars/represent"
	_ "golang.org/x/tools/cmd/present"
)
