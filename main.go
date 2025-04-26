/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import "github.com/runsecret/rsec/cmd"

var version = "local"

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}
