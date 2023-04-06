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
	"io"
	"log"
	"net"
	"time"
)

func reader(r io.Reader) string {
	output := ""
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil {
			log.Println("End of stream")
			return output
		}
		str := string(buf[0:n])
		output += str
		log.Println("Data read via socket", str)
	}
}

func sendCommandThroughSocket(command string) (string, error) {
	connection, err := net.Dial("unix", HAProxyUnixSocket)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	defer func() {
		err := connection.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	_, err = connection.Write([]byte(command))
	if err != nil {
		log.Println(err)
		return "", err
	}
	time.Sleep(1e9)
	return reader(connection), nil
}
