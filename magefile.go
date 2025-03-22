//go:build mage
// +build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Test runs all tests
func Test() error {
	return sh.RunV("go", "test", "./...")
}

type Lint mg.Namespace

func (Lint) Run() error {
	return sh.RunV("golangci-lint", "run", "--fix")
}

func (Lint) Check() error {
	return sh.RunV("golangci-lint", "run", "--fix")
}

type Format mg.Namespace

// ensure all files are formatted
func (Format) Check() error {
	violations, err := sh.Output("gofmt", "-l", ".")
	if err != nil {
		return err
	}

	if violations != "" {
		return fmt.Errorf("format errors in files: %s", violations)
	}

	return nil
}

// format the entire project
func (Format) Run() error {
	return sh.Run("gofmt", "-w", ".")
}
