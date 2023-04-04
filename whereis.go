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
	"os/exec"
	"strings"
)

func executableExist(executableName string) (bool, error) {
	command := exec.Command("whereis", executableName)

	stdout, err := command.StdoutPipe()
	if err != nil {
		return false, err
	}

	err = command.Start()
	if err != nil {
		return false, err
	}

	slurp, err := io.ReadAll(stdout)
	if err != nil {
		return false, err
	}

	parts := strings.Split(string(slurp), ":")
	if len(parts) < 2 {
		return false, nil
	}

	if len(parts[1]) <= 1 {
		return false, nil
	}

	return true, nil
}
