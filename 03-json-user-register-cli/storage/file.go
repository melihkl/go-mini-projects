package storage

import (
	"encoding/json"
	"json-user-register-cli/models"
	"os"
)

const filePath = "data/users.json"

func LoadUsers() ([]models.User, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var users []models.User
	if len(file) > 0 {
		err = json.Unmarshal(file, &users)
		if err != nil {
			return nil, err
		}
	}

	return users, nil
}

func SaveUser(newUser models.User) error {
	users, err := LoadUsers()
	if err != nil {
		users = []models.User{}
	}

	users = append(users, newUser)

	data, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
