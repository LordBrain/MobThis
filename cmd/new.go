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
	"fmt"

	"github.com/LordBrain/MobThis/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new mobbing session",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		messages := make(chan string)
		state := make(chan string)
		repoURL, _ := cmd.Flags().GetString("repo")
		var rotationTime int
		var retro bool

		// newMobName := strings.ReplaceAll(sillyname.GenerateStupidName(), " ", "-")
		// fmt.Println(newMobName)
		// fmt.Println("new called")

		// Things to set for a new session

		// Your name (if not set in config)
		// Time per station
		if !viper.IsSet("moberName") && !viper.IsSet("git.type") {
			fmt.Println("Mober name or Git configurations not set. Run `MobThis config` first!")
			goto Done
		}
		if len(args) == 0 && repoURL == "" {
			cmd.Help()
			goto Done
		}

		fmt.Print("Rotation time per station (minutes): ")
		fmt.Scanln(&rotationTime)
		if rotationTime <= 5 {
			fmt.Println("Rotation time must be greater then 5 or not a string.")
			goto Done
		}

		retro = utils.AskForConfirmation("Retro after each round?")

		fmt.Println("Mob Details:")
		fmt.Println("Rotation time: ", rotationTime)
		fmt.Println("Retro: ", retro)

		//go utils.CheckAPI(state)
		go utils.Other(messages)
		go utils.ReadStateChannel(messages, state)
		go utils.ReadMessageChannel(messages)

		//Listen to commands entered from the user
		fmt.Println("Mob session started. Wait for everyone to join then type 'Start' when ready.")
		utils.ReadInput(state)

	Done:
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	newCmd.Flags().StringP("repo", "r", "", "The name of the repository")
}
