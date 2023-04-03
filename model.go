package main

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
