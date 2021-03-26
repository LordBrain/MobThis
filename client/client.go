package client

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// V3 is the Gungnir-API struct type
type V1 struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
}

// NewClient provides a client to MobThis API
func NewClient(mobthisURL string) *V1 {

	httpClient := http.DefaultClient
	httpClient.Timeout = time.Second * 5

	c := &V1{HTTPClient: httpClient}

	c.BaseURL, _ = url.Parse(mobthisURL)

	return c
}

func CreateMob() {

	//Create a new mob session
}

func MobState() {
	//Check on the state of the session

}

func StartMob() {

	//start a new session

}

func JoinMob() {

	//Join a running mob session

}

func EndMob() {

	//End a mobbing session

}

func (c *V1) CheckAPI(state chan string) {

	for {
		//Poll API

		//Update state
		select {
		case stat := <-state:
			fmt.Println("Status recieved: ", stat)

		}
		// time.Sleep(50 * time.Millisecond)

	}
}
