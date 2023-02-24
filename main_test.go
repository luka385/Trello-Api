package main

import "testing"

func TestCreateTask(t *testing.T) {
	//Create a Test Task
	task := &Task{
		ID:       "",
		Title:    "Test title",
		Desc:     "Test Desc",
		Category: "Test Category",
		ListID:   "63f7d57dee675a5c3e3012f6",
		BoardID:  "63f7d57dee675a5c3e3012ef",
	}

	//Create the test authenticator
	auth := &TrelloAuth{
		AppKey:   "4e3b18f231da41059b410c8a9496a004",
		Token:    "ATTAbe8db41adb77a04b31d11c5e1d81150229b956af2d7ca6ab077eed80fad27ac8B11EF61C",
		EndPoint: "https://api.trello.com/1",
	}

	// Calling the test funtion
	err := CreateNewTask(task, auth)

	//verify error
	if err != nil {
		t.Errorf("Error when creating the task : %s", err)
	}

	// Cheking if the ID was updated correctly
	if task.ID == "" {
		t.Errorf("The ID of the task was not updated correctly : %s", err)
	}

}
