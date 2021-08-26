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
	"it-mis/internal/pkg/debugx"
	"it-mis/internal/pkg/pathx"
	"it-mis/internal/pkg/rsax"
	"os"
	"path"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var genrsaCmd = &cobra.Command{
	Use:   "install",
	Short: "install and init model",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		outputPath := path.Join(pathx.RootPath, "/assets/secure-rsa")
		// 判断secre-rsa是否存在
		if _, err := os.Stat(path.Join(outputPath, "/private.pem")); !os.IsNotExist(err) {
			debugx.PrintInfo("rsa file already exists")
		}
		if _, err := os.Stat(path.Join(outputPath, "/public.pem")); !os.IsNotExist(err) {
			debugx.PrintInfo("rsa file already exists")
		}
		rsax.GenRSA(1024, outputPath)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
