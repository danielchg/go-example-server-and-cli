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
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the server",
	Long:  `Login to the server in order to create new articles from the CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		u := fmt.Sprintf("%s/u/login", viper.GetString("url"))
		id := viper.GetString("id")
		pass := viper.GetString("pass")

		loginServer(u, id, pass)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func loginServer(u, id, pass string) {
	formData := url.Values{"username": {id}, "password": {pass}}
	c := &http.Client{}
	req, _ := http.NewRequest("POST", u, strings.NewReader(formData.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := c.Do(req)
	if err != nil {
		log.Fatalf("can not connect to URL: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		log.Fatalf("error calling to URL: %v", res.StatusCode)
	}

	viper.Set("cookie", res.Cookies())
	viper.WriteConfig()
	fmt.Printf("User %s loged successfully", id)

}
