/*
Copyright © 2024 Timothy Chatman tjchatman1986@gmail.com
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

type OSInfo struct {
	OS       string
	Hostname string
	User     string
}

// osInfoCmd represents the osInfo command
var osInfoCmd = &cobra.Command{
	Use:   "osInfo",
	Short: "Query basic sys Info",
	Long:  `Get information for software that is OS related such as OS, Version or distribution, etc... `,
	Run: func(cmd *cobra.Command, args []string) {
		info := OSInfo{}
		fmt.Println("osInfo called")
		host, hostErr := os.Hostname()
		if hostErr != nil {
			hostErr = errors.New("Error retrieving hostname.")
			log.Fatal(hostErr)
		}
		kernel, errUname := exec.Command("uname", "-or").Output()
		if errUname != nil {
			errUname = errors.New("Error getting kernel info.")
			log.Fatal(errUname)
		}
		user, errWho := exec.Command("whoami").Output()
		if errWho != nil {
			errWho = errors.New("Error getting user info.")
			log.Fatal(errWho)
		}
		info.Hostname = host
		info.OS = string(kernel)
		info.User = string(user)
		fmt.Println(info)
	},
}

func init() {
	rootCmd.AddCommand(osInfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// osInfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// osInfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
