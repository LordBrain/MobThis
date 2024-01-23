/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/LordBrain/MobThis/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "set config values for MobThis",
	Long: `A configureation file is required for MobThis. 
It used for holding your prefered Mobbing name as well as how to connect to Github.

Values:
mobName - The name to display for each mob role.
codePath - location where to pull the Git repo's.
mobthisAddress - MobThis API address.

git:
  type - How you connect to Github, ssh key or authentication with a token.
  username - Your Github Username.
  token - Your Github Token.
  key - path to your id_rsa file.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Setting up your configuration file!")

		scanner := bufio.NewScanner(os.Stdin)
		var moberName string
		var codePath string
		var mobthisAddress string
		var gitType string
		var gitUsername string
		var gitToken string
		var gitKey string

		//Mober username
		if viper.IsSet("moberName") {

			fmt.Print("Enter Name [" + viper.GetString("moberName") + "]: ")

			if scanner.Scan() {
				line := scanner.Text()
				moberName = line
			}

			if moberName == "" {
				moberName = viper.GetString("moberName")
			}
		} else {
			githubUsername, _ := utils.GitUsername()
			fmt.Print("Enter Name [" + githubUsername + "]: ")
			if scanner.Scan() {
				line := scanner.Text()
				moberName = line
			}
			if moberName == "" {
				moberName = githubUsername
			}
		}
		//Set it in viper
		viper.Set("moberName", moberName)

		//Path to store the code being worked on
		if viper.IsSet("codePath") {
			fmt.Print("Working directory for code [" + viper.GetString("codePath") + "]: ")
			fmt.Scanln(&codePath)
			if codePath == "" {
				codePath = viper.GetString("codePath")
			}
		} else {
			fmt.Print("Working directory for code: ")
			fmt.Scanln(&codePath)
		}
		//Set it in viper
		viper.Set("codePath", codePath)

		//Mobthis API Address
		if viper.IsSet("mobthisAddress") {
			fmt.Print("MobThis API Host [" + viper.GetString("mobthisAddress") + "]: ")
			fmt.Scanln(&mobthisAddress)
			if codePath == "" {
				mobthisAddress = viper.GetString("mobthisAddress")
			}
		} else {
			fmt.Print("MobThis API Host: ")
			fmt.Scanln(&mobthisAddress)
		}
		//Set it in viper
		viper.Set("mobthisAddress", mobthisAddress)

		//Git Option
		//Git connection type, ssh or auth
		if viper.IsSet("git.type") {
			fmt.Print("How do you connect to git? (Options: ssh, auth) [" + viper.GetString("git.type") + "]: ")
			fmt.Scanln(&gitType)
			if gitType == "" {
				gitType = viper.GetString("git.type")
			}
			if strings.ToLower(gitType) != "ssh" && strings.ToLower(gitType) != "auth" {
				fmt.Println("type must be SSH or Auth only")
				os.Exit(1)
			}

		} else {
			fmt.Print("How do you connect to git? (Options: ssh, auth): ")
			fmt.Scanln(&gitType)
			if strings.ToLower(gitType) != "ssh" && strings.ToLower(gitType) != "auth" {
				fmt.Println("type must be ssh or auth only")
				os.Exit(1)
			}
		}
		//Set it in viper
		viper.Set("git.type", strings.ToLower(gitType))

		//if connection type is ssh, get rsa file
		if viper.GetString("git.type") == "ssh" {
			if viper.IsSet("git.key") {
				fmt.Print("Git Public Key file [" + viper.GetString("git.key") + "]: ")
				fmt.Scanln(&gitKey)
				if gitKey == "" {
					gitKey = viper.GetString("git.key")
				}

			} else {
				fmt.Print("Git Public Key file [" + os.Getenv("HOME") + "/.ssh/id_rsa" + "]: ")
				fmt.Scanln(&gitKey)
				if gitKey == "" {
					gitKey = os.Getenv("HOME") + "/.ssh/id_rsa"
				}

			}
			//Set it in viper
			viper.Set("git.key", gitKey)

		}

		//if connection type is auth, get git username
		if viper.GetString("git.type") == "auth" {
			if viper.IsSet("git.username") {
				fmt.Print("Git username [" + viper.GetString("git.username") + "]: ")
				fmt.Scanln(&gitUsername)
				if gitUsername == "" {
					gitUsername = viper.GetString("git.username")
				}

			} else {
				githubUsername, _ := utils.GitUsername()
				fmt.Print("Git username [" + githubUsername + "]: ")
				fmt.Scanln(&gitUsername)
				if gitUsername == "" {
					gitUsername = githubUsername
				}

			}
			//Set it in viper
			viper.Set("git.username", gitUsername)

			if viper.IsSet("git.token") {
				fmt.Print("Git Token [" + viper.GetString("git.token") + "]: ")
				fmt.Scanln(&gitToken)
				if gitToken == "" {
					gitToken = viper.GetString("git.token")
				}

			} else {
				fmt.Print("Git Token: ")
				fmt.Scanln(&gitToken)

			}
			//Set it in viper
			viper.Set("git.token", gitToken)

		}

		err := viper.WriteConfig()
		if err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
