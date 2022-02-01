/*
Copyright © 2022 Sam Kim <netpple@gmail.com>
This file is part of Sam's CLI myapp.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/spf13/cobra"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		osInfo, err := ioutil.ReadFile("/etc/os-release")
		check(err)
		fmt.Print(string(osInfo))

		kernelInfo, err := exec.Command("uname", "-r").Output()
		check(err)
		fmt.Println("Kernel Version: ", string(kernelInfo))
	},
}

func init() {
	getCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
