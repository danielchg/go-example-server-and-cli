/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the CLI with the server access",
	Long:  `Start a wizard that allow to configure the access to the server`,
	Run: func(cmd *cobra.Command, args []string) {
		cfgWizard()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

}

func cfgWizard() {
	fmt.Print("Enter the server URL: ")
	var u string
	fmt.Scanf("%s", &u)

	fmt.Print("Username: ")
	var id string
	fmt.Scanf("%s", &id)

	fmt.Print("Password: ")
	var pass string
	fmt.Scanf("%s", &pass)

	fmt.Printf("Add to the config the following URL: %s", u)
	viper.Set("url", u)
	viper.Set("id", id)
	viper.Set("pass", pass)
	viper.WriteConfig()
}
