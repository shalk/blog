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
	"errors"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "create new post",
	Long:  `create new post `,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires a color argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fileName := "posts/" + args[0]
		_, err := os.Stat("posts")
		if err != nil {
			log.Fatal(err)
		}
		if os.IsNotExist(err) {
			err := os.Mkdir("posts", 644)
			if err != nil {
				log.Fatal(err)
			}
		}
		_, err = os.Stat(fileName)
		if os.IsNotExist(err) {
			file, err := os.Create(fileName)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("create file %s", fileName)
			defer file.Close()
		} else {
			log.Printf("file exists %s", fileName)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
