/*
Copyright 2019 The Volcano Authors.

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
	"os"

	"github.com/spf13/cobra"
	"k8s.io/component-base/cli"

	"volcano.sh/volcano/cmd/cli/util"
	"volcano.sh/volcano/pkg/cli/vqueues"
)

func main() {
	rootCmd := cobra.Command{
		Use:   "vqueues",
		Short: "view queue information",
		Long:  `view information of a queue with specified name or queues from the same namespace`,
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckError(cmd, vqueues.GetQueue())
		},
	}

	jobGetCmd := &rootCmd
	vqueues.InitGetFlags(jobGetCmd)

	code := cli.Run(&rootCmd)
	os.Exit(code)
}
