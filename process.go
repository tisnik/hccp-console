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

import "strings"
import "strconv"
import "io"
import "os/exec"

func processRunning(processName string) bool {
	command := exec.Command("pidof", processName)

	err := command.Run()
	return err == nil
}

func getProcessPID(processName string) (int, error) {
	command := exec.Command("pidof", processName)

	stdout, err := command.StdoutPipe()
	if err != nil {
		return -1, err
	}

	err = command.Start()
	if err != nil {
		return -1, err
	}

	slurp, _ := io.ReadAll(stdout)

	pid, err := strconv.Atoi(strings.TrimSpace(string(slurp)))
	if err != nil {
		return -1, err
	}

	err = command.Wait()
	if err != nil {
		return -1, err
	}

	return pid, nil
}
