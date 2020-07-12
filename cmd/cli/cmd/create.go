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

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create and publish a new article to the server",
	Long:  `Create and publish a new article to the server with a wizard or by a file.`,
	Run: func(cmd *cobra.Command, args []string) {
		createArticle()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

}

func createArticle() {

	cookie := viper.GetString("cookie")
	fmt.Print("Please enter the article title: ")
	var title string
	fmt.Scanf("%s", &title)

	fmt.Print("Please enter content of the article: ")
	var content string
	fmt.Scanf("%s", &content)

	u := fmt.Sprintf("%s/article/create", viper.GetString("url"))
	formData := url.Values{"title": {title}, "content": {content}}
	c := &http.Client{}
	req, _ := http.NewRequest("POST", u, strings.NewReader(formData.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", cookie)
	res, err := c.Do(req)
	if err != nil {
		log.Fatalf("can not connect to URL: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		log.Fatalf("error calling to URL: %v", res.StatusCode)
	}
	fmt.Println("Article created successfully")
}
