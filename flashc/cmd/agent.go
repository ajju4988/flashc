// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// agentCmd represents the agent command
var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "A brief description of your command",
	Long: `
	Gets the status from Agents. Status of the agents are maintained in database.
	On running the command, the information is fetched from DB`,
	Run: Agent,
}

// Agent Gets the status from agent.
func Agent(cmd *cobra.Command, args []string) {
	err := checkController(ControllerNode)
	if err != nil {
		fmt.Printf("\n\n\nCheck if Controller is running on the node with IP - " + ControllerNode + "\n")
		fmt.Println("Commands cannot be executed on that node")
		fmt.Printf("use -C option to preovide a different controller\n\n\n")
	}
	fmt.Println("Use -h for help")

}
func init() {
	RootCmd.AddCommand(agentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// agentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// agentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
