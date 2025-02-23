package internal

import (
	"encoding/json"
	"errors"
	"maps"
	"os"
	"sync"
)

var filePath string = "passwords.json"

type PasswordStore struct {
	sync.Mutex                   // prevents concurent access of the passwordstore struct
	Passwords  map[string]string `json:"passwords"` //struct tags -> tells how to handle this field to a certain package -> in this case it is telling json package to put this variable under the field name of passwords
}

// loads existing passwords from a json file
func LoadPasswords() (*PasswordStore, error) {
	store := &PasswordStore{Passwords: make(map[string]string)}

	file, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return store, nil
		}
		return nil, err
	}

	err = json.Unmarshal(file, store)
	if err != nil {
		return nil, err
	}

	return store, nil
}

// saves the passwords inside the password store's Passwords variable into the json file
func (ps *PasswordStore) SavePassword() error {
	ps.Lock() // locks the sync.Mutex object
	defer ps.Unlock()

	data, err := json.MarshalIndent(ps, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}

// adds the new password into the PasswordStore's Passwords variable and stores it in the json file
func (ps *PasswordStore) AddPassword(service string, password string) error {
	if _, exists := ps.Passwords[service]; exists {
		return errors.New("service already exists")
	}
	ps.Passwords[service] = password
	return ps.SavePassword()
}

// retrieves the password for a specified service
func (ps *PasswordStore) GetPassword(service string) (string, error) {
	password, exists := ps.Passwords[service]
	if !exists {
		return "", errors.New("service not found")
	}
	return password, nil
}

// deletes a password
func (ps *PasswordStore) RemovePassword(service string) error {
	if _, exists := ps.Passwords[service]; !exists {
		return errors.New("service not found")
	}

	delete(ps.Passwords, service)
	return ps.SavePassword()
}

// lists all the stored services
func (ps *PasswordStore) ListServices() []string {
	keysIter := maps.Keys(ps.Passwords)
	var keys []string
	for key := range keysIter {
		keys = append(keys, key)
	}
	return keys
}
