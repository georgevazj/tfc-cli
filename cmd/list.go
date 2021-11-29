/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/georgevazj/tfc-cli/models"
	"github.com/spf13/cobra"
)

const TIMEOUT = time.Duration(5 * time.Second)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		organization, _ := cmd.Flags().GetString("organization")

		log.Printf("Organization: %s", organization)
		log.Printf("TFAPI: %s", TFAPI)

		url := TFAPI + "/organizations/" + organization + "/workspaces"
		client := http.Client{
			Timeout: TIMEOUT,
		}
		request, err := http.NewRequest("GET", url, nil)
		request.Header.Add("Authorization", "Bearer "+token)
		request.Header.Add("Content-Type", "application/vnd.api+json")
		if err != nil {
			log.Fatal(err)
		}

		response, err := client.Do(request)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		var d models.Data
		err = json.Unmarshal(body, &d)
		if err != nil {
			log.Fatal(err)
		}
		output, err := json.MarshalIndent(d, "", "  ")
		log.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
