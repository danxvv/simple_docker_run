/*
Copyright © 2022 Daniel Mojica  <danxvv@gmail.com>

*/
package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var gockerCmd = &cobra.Command{
	Use:   "gocker",
	Short: "Simple CLI for run docker container only with port",
	Long: `A simple CLI for run easy CDD containers with docker, also include a simple command for
	kill all running containers`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\nGocker\n")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if args[0] != "kill" {
			for i, arg := range args {

				ports := cmd.Flag("port").Value.String()
				ports = strings.Replace(ports, "[", "", -1)
				ports = strings.Replace(ports, "]", "", -1)
				portsParsers := strings.Split(ports, ",")
				runOneContainer(arg, portsParsers[i])
			}

			return nil
		}
		killAll()
		return nil

	},
}

func init() {
	rootCmd.AddCommand(gockerCmd)
	gockerCmd.Flags().StringArrayP("port", "p", []string{"8080"}, "port")

}

func runOneContainer(container string, port string) {
	containerFolder := fmt.Sprintf("/home/danxvv/curadeuda/%s", container)
	containerName := fmt.Sprintf("%s_local", container)
	containerImage := fmt.Sprintf("%s:latest", container)
	cmd := exec.Command("docker", "run", "--rm", "-d", "--env-file=.env", "-v", containerFolder+":/usr/src/app", "-p",
		port+":5000", "--name", containerName, containerImage)
	cmd.Dir = containerFolder
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		fmt.Println(stderr.String())
	}
}

func killAll() {
	cmd := exec.Command("docker", "ps", "-q")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		fmt.Println(stderr.String())
	}
	for _, container := range strings.Split(stdout.String(), "\n") {
		if container != "" {
			cmd = exec.Command("docker", "kill", container)
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr
			err = cmd.Run()
			if err != nil {
				fmt.Println(err)
				fmt.Println(stderr.String())
			}
		}
	}
}
