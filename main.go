package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// struct for create tasks
type Task struct {
	ID       string `json:"id"`
	Title    string `json:"name"`
	Desc     string `json:"desc"`
	Category string `json:"category,omitempty"`
	ListID   string `json:"idList"`
	BoardID  string `json:"-"`
}

// Authenticator for trello API
type TrelloAuth struct {
	AppKey   string
	Token    string
	EndPoint string
}

func main() {

	// Creatring a problem
	FirstTask := Task{
		Title:   "a new satellite",
		Desc:    "this satellite will help the connections to remain more stable",
		ListID:  "63f7d57dee675a5c3e3012f6",
		BoardID: "63f7d57dee675a5c3e3012ef",
	}

	// configuring authentication in trello
	auth := TrelloAuth{
		AppKey:   "4e3b18f231da41059b410c8a9496a004",
		Token:    "ATTAbe8db41adb77a04b31d11c5e1d81150229b956af2d7ca6ab077eed80fad27ac8B11EF61C",
		EndPoint: "https://api.trello.com/1",
	}

	// We using our function and verify the error
	err := CreateNewTask(&FirstTask, &auth)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("SUCCESSFULLY CREATED TASK")
}

func CreateNewTask(task *Task, auth *TrelloAuth) error {
	// Encoding Json
	JsonTask, err := json.Marshal(task)
	if err != nil {
		log.Fatal(err)
	}

	// Creating POST request
	url := fmt.Sprintf("%s/card?key=%s&token=%s", auth.EndPoint, auth.AppKey, auth.Token)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(JsonTask))
	if err != nil {
		log.Fatal(err)
	}

	// Setting the Header
	req.Header.Set("content-type", "application/json")

	// Making the POST request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// Closing Connection
	defer resp.Body.Close()

	// Verify if request was successful
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("The task could not be created in Trello : %s", resp.Status)
	}

	//Decoding the JSON response
	var createTask Task
	err = json.NewDecoder(resp.Body).Decode(&createTask)
	if err != nil {
		log.Fatal(err)
	}

	//Updating the ID of the task created
	task.ID = createTask.ID

	return nil
}
