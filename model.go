package main

type Route struct {
	Connection string
	Status     string
	Enable     bool
}

var routes []Route

func init() {
	routes = []Route{
		Route{"first", "xyzzy", false},
		Route{"second", "abc", true},
	}
}
