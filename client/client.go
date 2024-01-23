package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/LordBrain/MobThis/utils"
)

// V1 is the MobThis-API struct type
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

func (client *V1) CreateMob(mobDetails utils.NewMobSession) (utils.MobSession, error) {
	//Create a new mob session
	var session utils.MobSession
	rel := &url.URL{Scheme: "http", Path: "/v1/mob"}
	u := client.BaseURL.ResolveReference(rel)

	fmt.Println(client.BaseURL.String())
	mob, err := json.Marshal(mobDetails)
	if err != nil {
		return session, errors.New("error marshaling mob")
	}

	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(mob))
	if err != nil {
		return session, err
	}

	q := req.URL.Query()

	req.Header.Set("Content-Type", "application/json")
	req.URL.RawQuery = q.Encode()
	session, err = commonHTTP(req)

	if err != nil {
		fmt.Println(err)
		return session, errors.New("failed creating session")

	}

	return session, nil

}

func (client *V1) MobState(session string) (utils.MobSession, error) {
	//Check on the state of the session
	var mobSession utils.MobSession
	rel := &url.URL{Path: "/v1/mob/" + session}
	u := client.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return mobSession, err
	}

	q := req.URL.Query()

	req.Header.Set("Content-Type", "application/json")
	req.URL.RawQuery = q.Encode()
	mobSession, err = commonHTTP(req)
	if err != nil {
		return mobSession, errors.New("failed creating session")

	}

	return mobSession, nil
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

func LeaveMob() {

	//Leave the mob session
}

func (c *V1) CheckAPI(mobDetails utils.MobSession, messages, state chan string) {

	myMob := mobDetails

	for {
		// fmt.Println("Poll")
		//Poll API
		currentCheck, _ := c.MobState(myMob.SessionName)
		addMobber, removeMobber := utils.CheckMobbers(myMob.Mobbers, currentCheck.Mobbers)
		if len(addMobber) > 0 {
			for _, mobber := range addMobber {
				myMob.Mobbers = append(myMob.Mobbers, mobber)
				messages <- mobber + " joined"
			}
		}
		if len(removeMobber) > 0 {
			for i, v := range myMob.Mobbers {
				if v == myMob.Mobbers[i] {
					copy(myMob.Mobbers[i:], myMob.Mobbers[i+1:])         // Shift a[i+1:] left one index.
					myMob.Mobbers[len(myMob.Mobbers)-1] = ""             // Erase last element (write zero value).
					myMob.Mobbers = myMob.Mobbers[:len(myMob.Mobbers)-1] // Truncate slice.
					messages <- v + " left"
				}

			}
		}
		//Update state
		// select {
		// case stat := <-state:
		// 	fmt.Println("Status recieved: ", stat)

		// }
		time.Sleep(1 * time.Second)

	}
}

func commonHTTP(req *http.Request) (utils.MobSession, error) {
	var statusCode int
	var bodyMessage utils.MobSession
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return bodyMessage, err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	statusCode = resp.StatusCode
	err = json.Unmarshal(body, &bodyMessage)
	if err != nil {
		statusCode = http.StatusNoContent
	}

	switch statusCode {
	case http.StatusInternalServerError:
		return bodyMessage, errors.New("error connecting to server")
	case http.StatusBadRequest:
		return bodyMessage, errors.New("invalid request")

	default:
		return bodyMessage, nil
	}
}
