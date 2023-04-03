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
	return nil
}

func disableRouteWithID(ids []string) error {
	routerID, err := retrieveRouterID(ids)
	if err != nil {
		return err
	}

	err = updateRouteState(routerID, false)
	return nil
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
