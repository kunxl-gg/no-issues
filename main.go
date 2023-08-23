/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/kunxl-gg/no-issues/cmd"
	"github.com/kunxl-gg/no-issues/constants"
)

func main() {
	constants.InitialiseEnvironmentVariables()
	cmd.Execute()
}
