/*
Copyright © 2023 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func retrieveRouterID(id string) (int, error) {
	if id == "" {
		return -1, errors.New("ID not provided, or multiple IDs provided")
	}

	routerID, err := strconv.Atoi(id)
	if err != nil {
		return -1, err
	}

	return routerID, nil
}

func enableRouteWithID(id string) error {
	routerID, err := retrieveRouterID(id)
	if err != nil {
		return err
	}

	backend := getRouteBackend(routerID)
	command := "enable server " + backend + "/" + backend + "\n"
	log.Println(command)

	output, err := sendCommandThroughSocket(command)
	log.Println(output)
	if err != nil {
		return err
	}

	err = updateRouteState(routerID, true)
	return err
}

func disableRouteWithID(id string) error {
	routerID, err := retrieveRouterID(id)
	if err != nil {
		return err
	}

	backend := getRouteBackend(routerID)
	command := "disable server " + backend + "/" + backend + "\n"
	log.Println(command)

	output, err := sendCommandThroughSocket(command)
	log.Println(output)
	if err != nil {
		return err
	}

	err = updateRouteState(routerID, false)
	return err
}

func updateRouteState(routerID int, enable bool) error {
	if !existingRouter(routerID) {
		return fmt.Errorf("Router with ID %d does not exist", routerID)
	}

	if enable {
		updateRouteStatus(routerID, "connected", true)
	} else {
		updateRouteStatus(routerID, "disconnected", false)
	}
	return nil
}

func readHAProxyConfiguration() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	path := filepath.Join(homedir, HAProxyConfigFile)
	config, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(config), nil
}
