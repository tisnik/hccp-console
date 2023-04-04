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
	"strconv"
)

func retrieveRouterID(ids []string) (int, error) {
	if len(ids) != 1 {
		return -1, errors.New("ID not provided, or multiple IDs provided")
	}

	routerID, err := strconv.Atoi(ids[0])
	if err != nil {
		return -1, err
	}

	return routerID, nil
}

func enableRouteWithID(ids []string) error {
	routerID, err := retrieveRouterID(ids)
	if err != nil {
		return err
	}

	err = updateRouteState(routerID, true)
	return err
}

func disableRouteWithID(ids []string) error {
	routerID, err := retrieveRouterID(ids)
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
