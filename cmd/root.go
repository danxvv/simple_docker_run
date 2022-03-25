/*
Copyright © 2022 Daniel Mojica  <danxvv@gmail.com>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "simple_docker_run",
	Short: "use gocker <container> -p <port> to run a container",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
