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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all the articles from the server",
	Long:  `list all the articles from the server`,
	Run: func(cmd *cobra.Command, args []string) {
		url := viper.GetString("url")
		fmt.Printf("Server: %s\n", url)
		a := listArticles(url)
		fmt.Printf("First article content: %v\n", a[0].Content)
		fmt.Printf("Second article content: %v\n", a[1].Content)
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

type article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Body define the struct taht goes in Response
type Body struct {
	Bytes  []byte
	String string
}

func listArticles(url string) []article {

	c := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	res, err := c.Do(req)
	if err != nil {
		log.Fatalf("can not connect to URL: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("error calling to URL: %v", res.StatusCode)
	}

	var a []article
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &a)
	if err != nil {
		log.Fatalf("can not decode the JSON output: %v", err)
	}

	return a
}
