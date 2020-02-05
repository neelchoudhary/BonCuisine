package utils

import "io/ioutil"

const (
	envFilePath = "config/env"
)

// WriteEnvToFile writes environment to file
func WriteEnvToFile(env string) error {
	return ioutil.WriteFile(envFilePath, []byte(env), 0600)
}

// ReadEnvFromFile reads environment from file
func ReadEnvFromFile() (string, error) {
	data, err := ioutil.ReadFile(envFilePath)
	return string(data), err
}

// WriteFile writes data to file
func WriteFile(path string, data string) error {
	return ioutil.WriteFile(path, []byte(data), 0600)
}

// ReadFile reads data from file
func ReadFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	return string(data), err
}
