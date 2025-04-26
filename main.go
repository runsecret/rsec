/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import "github.com/runsecret/rsec/cmd"

var Version = "local"

func main() {
	cmd.SetVersion(Version)
	cmd.Execute()
}
