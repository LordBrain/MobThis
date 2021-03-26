package utils

import (
	"time"
)

func Other(message chan string) {
	time.Sleep(3 * time.Second)
	message <- "User 1 joined"
	time.Sleep(3 * time.Second)
	message <- "User 2 joined"
	time.Sleep(3 * time.Second)
	message <- "User 3 joined"
	time.Sleep(3 * time.Second)
	message <- "User 4 joined"
	time.Sleep(3 * time.Second)
	message <- "User 5 joined"
	time.Sleep(3 * time.Second)
	message <- "User 6 joined"
	time.Sleep(3 * time.Second)
	message <- "User 7 joined"
	time.Sleep(3 * time.Second)
	message <- "User 8 joined"

}
