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

// Route represents mocked route table
type Route struct {
	ID         int
	Connection string
	Status     string
	Enabled    bool
}

var routes []Route

func init() {
	routes = []Route{
		Route{1, "first route", "connected", true},
		Route{2, "second route", "disconnected", false},
		Route{3, "third route", "error", false},
	}
}

func existingRouter(routerID int) bool {
	// TODO: routes as regular interface
	for _, route := range routes {
		if route.ID == routerID {
			return true
		}
	}
	return false
}

func updateRouteStatus(routerID int, status string, enable bool) {
	// TODO: routes as regular interface
	for i := 0; i < len(routes); i++ {
		if routes[i].ID == routerID {
			routes[i].Status = status
			routes[i].Enabled = enable
		}
	}
}
