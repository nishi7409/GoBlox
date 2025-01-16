package users

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type APIResponse struct {
	Data []struct {
		ID int `json:"id"`
	} `json:"data"`
}

func GetIDFromUsername(username string) (int, error) {
	client := &http.Client{}
	payload := map[string]interface{}{
		"usernames":          []string{username},
		"excludeBannedUsers": true,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal payload: %v", err)
		return -1, err
	}

	resp, err := client.Post("https://users.roblox.com/v1/usernames/users", "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Printf("POST request failed: %v", err)
		return -1, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("POST request failed: %v", resp.Status)
		return -1, nil
	}

	var apiResponse APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		log.Printf("Failed to decode response: %v", err)
		return -1, nil
	}
	resp.Body.Close()

	if len(apiResponse.Data) == 0 {
		return -1, errors.New("username is invalid")
	}

	return apiResponse.Data[0].ID, nil
}
