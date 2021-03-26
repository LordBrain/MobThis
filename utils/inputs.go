package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func ReadInput(state chan string) {
	//Catch when someone does ctr+c to quite.
	CaptureExit()

	var text string
	for {

		fmt.Scanln(&text)

		if text != "" {
			state <- text
			text = ""
		}

	}
}

//ReadMessageChannel outputs to the terminal when a message is received.
func ReadMessageChannel(messages chan string) {
	var msg string

	for {
		select {
		case msg = <-messages:
			fmt.Println(msg)

		}
	}
}

//ReadStateChannel takes action based on a state message.
func ReadStateChannel(messages, state chan string) {
	var stateUpdate string

	for {
		select {

		case stateUpdate = <-state:
			if strings.ToLower(strings.TrimSpace(stateUpdate)) == "start" {
				messages <- "Mobbing session Starting..."
				MobNotify("MobThis", "Mob is starting!")
				stateUpdate = ""
			}

		}
	}
}

func CaptureExit() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Quitting Mob Session")
		os.Exit(0)
	}()
}

func AskForConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}
