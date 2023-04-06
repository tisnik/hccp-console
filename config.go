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

const (
	serverAddress = "localhost:8081"

	// HAProxyConfigFile contains name of configuration file to be read by controller
	HAProxyConfigFile = "haproxy.cfg"

	// HAProxyExecutableName contains name of HAProxy executable name stored anywhere on path
	HAProxyExecutableName = "haproxy"

	// HAProxyProcessName contains name of HAProxy process name
	HAProxyProcessName = "haproxy"

	// HAProxyUnixSocket contians Unix socket used to communicate with HAProxy
	HAProxyUnixSocket = "/tmp/haproxy.sock"
)
