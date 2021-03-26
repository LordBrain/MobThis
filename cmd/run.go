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
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/LordBrain/MobThis/utils"
	"github.com/Pallinder/sillyname-go"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// reader := bufio.NewReader(os.Stdin)
		fmt.Println("run called")
		messages := make(chan string)

		// go yolo(messages)
		fmt.Println("Start checking API")
		go checkAPI(messages)
		fmt.Println("Start reading input")
		// go readInput(reader)
		// var text string

		// for {
		// 	fmt.Print("-> ")
		// 	// text, _ := reader.ReadString('\n')
		// 	// // convert CRLF to LF
		// 	// text = strings.Replace(text, "\n", "", -1)
		// 	fmt.Scanln(&text)
		// 	if strings.Compare("hi", text) == 0 {
		// 		fmt.Println("hello, Yourself")
		// 	}
		// 	messages <- text
		// 	time.Sleep(1 * time.Millisecond)
		// }

		// _, err := git.PlainClone("/tmp/bar", false, &git.CloneOptions{
		// 	URL:      "git@git.target.com:ResiliencyEngineering/Freya-cli.git",
		// 	Progress: os.Stdout,
		// })

		// fmt.Println("Error: ", err)
		// cloneRepo("git@git.target.com:ResiliencyEngineering/Freya-cli.git")
		newMobName := strings.ReplaceAll(sillyname.GenerateStupidName(), " ", "-")
		// err := utils.GitCloneSSH("git@git.target.com:ResiliencyEngineering/Freya-cli.git", viper.GetString("codePath")+"/"+newMobName, viper.GetString("git.key"))
		// err := utils.GitCloneAuth("https://github.com/go-git/go-git.git", viper.GetString("codePath")+"/"+newMobName, viper.GetString("git.username"), "")
		err := utils.GitCloneAuth("https://github.com/go-git/go-git.git", viper.GetString("codePath")+"/"+newMobName, viper.GetString("git.username"), viper.GetString("git.token"))

		if err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func yolo(rando chan string) {
	count := 0
	for i := 1; i <= 1000; i++ {
		count += 1
		fmt.Println(count)
		time.Sleep(1 * time.Second)
	}
}

func checkAPI(state chan string) {

	// for {
	// 	select {
	// 	case msg := <-message:
	// 		fmt.Println("received message: ", msg)
	// 	default:

	// 	}
	// 	time.Sleep(1 * time.Millisecond)

	// }

}

func cloneRepo(url string) error {

	var publicKey *ssh.PublicKeys
	sshPath := os.Getenv("HOME") + "/.ssh/id_rsa"
	sshKey, _ := ioutil.ReadFile(sshPath)
	publicKey, keyError := ssh.NewPublicKeys("git", []byte(sshKey), "")
	if keyError != nil {
		fmt.Println("key error:", keyError)
		return keyError
	}
	_, err := git.PlainClone("/tmp/test", false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
		Auth:     publicKey,
	})
	if err != nil {
		fmt.Println("Clone error: ", err)
		return err
	}
	return nil
}
